package rulesdb

import (
	"context"
	"fmt"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"
	"strings"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"

	// _ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type sRulesDb struct {
	sctx g.Ctx
}

var RuleChName = "rule_ch"
var AbiChName = "abi_ch"

// func (s *sRulesDb) Set(ctx context.Context, ruleId, rules string) error {
// 	// g.Redis().Set(s.ctx, name, rules)

//		i, err := dao.Rule.Ctx(s.ctx).Data(do.Rule{RuleId: ruleId, Rules: rules}).Where(do.Rule{
//			RuleId: ruleId,
//		}).Count()
//		if i == 0 {
//			_, err = dao.Rule.Ctx(s.ctx).Data(do.Rule{RuleId: ruleId, Rules: rules}).Insert()
//		} else {
//			_, err = dao.Rule.Ctx(s.ctx).Data(do.Rule{RuleId: ruleId, Rules: rules}).Where(do.Rule{
//				RuleId: ruleId,
//			}).Update()
//		}
//		return err
//	}
func (s *sRulesDb) Get(ctx context.Context, ruleId string) (string, error) {
	// v, _ := g.Redis().Get(s.ctx, name)
	rule := &entity.Rule{}
	err := dao.Rule.Ctx(ctx).Where(dao.Rule.Columns().RuleId, ruleId).Scan(rule)
	return rule.Rules, err
}

func (s *sRulesDb) AllRules(ctx context.Context) map[string]string {
	rule := []entity.Rule{}
	err := dao.Rule.Ctx(ctx).Scan(&rule)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	rst := map[string]string{}
	for _, i := range rule {
		rst[i.RuleId] = i.Rules
	}
	return rst
}
func (s *sRulesDb) GetAbi(ctx context.Context, to string) (string, error) {
	contracts := &entity.ContractAbi{}
	err := dao.ContractAbi.Ctx(ctx).Where(dao.ContractAbi.Columns().Addr, to).Scan(contracts)
	return contracts.Abi, err
}

func (s *sRulesDb) subscription(conn *pgxpool.Conn, name string, notificationChannel chan *pgconn.Notification) {

	channelName := name
	_, err := conn.Exec(context.Background(), fmt.Sprintf("LISTEN %s", channelName))
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			case notification := <-notificationChannel:
				//todo: up rules, ctx.done
				switch notification.Channel {
				case RuleChName:
					load := notification.Payload
					ops := strings.Split(load, ",")
					if len(ops) == 2 {
						ruleId := ops[0]
						op := ops[1]
						if op == "up" {
							rules, _ := s.Get(s.sctx, ruleId)
							service.LEngine().UpRules(ruleId, rules)
						}
						if op == "rm" {
							service.LEngine().UpRules(ruleId, "")
						}
					}
				case AbiChName:
				}

				fmt.Println("Received notification:", notification)
			}
		}
	}()
}

func (s *sRulesDb) listenNotify(subNames []string) {
	l, _ := g.Cfg().Get(context.Background(), "database.default.0.link")
	fmt.Println(l.String())
	link := l.String()
	link = strings.Replace(link, "pgsql:", "postgres://", -1)
	link = strings.Replace(link, "tcp(", "", -1)
	link = strings.Replace(link, ")", "", -1)

	ctx := gctx.GetInitCtx()
	cfg, err := pgxpool.ParseConfig(link)
	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Release()

	///subscription
	notificationChannel := make(chan *pgconn.Notification)
	for _, n := range subNames {
		s.subscription(conn, n, notificationChannel)
	}

	for {
		_, err := conn.Exec(context.Background(), "SELECT 1")
		if err != nil {
			panic(err)
		}

		notifications, err := conn.Conn().WaitForNotification(context.Background())
		if err != nil {
			panic(err)
		}
		notificationChannel <- notifications
	}
}

func new() *sRulesDb {
	g.Redis().Exists(gctx.GetInitCtx())
	s := &sRulesDb{
		sctx: gctx.GetInitCtx(),
	}
	//todo: notify
	go s.listenNotify([]string{RuleChName, AbiChName})
	return s
}

func init() {
	service.RegisterRulesDb(new())
}
