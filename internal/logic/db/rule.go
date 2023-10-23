package db

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
	"github.com/mpcsdk/mpcCommon/mpccode"

	// _ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var RuleChName = "rule_ch"
var AbiChName = "abi_ch"

func (s *sDB) GetRules(ctx context.Context, ruleId string) (string, error) {
	rule := &entity.Rule{}

	err := dao.Rule.Ctx(ctx).Where(dao.Rule.Columns().RuleId, ruleId).Scan(rule)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("ruleid", ruleId),
		))
		return "", err
	}
	return rule.Rules, err
}

func (s *sDB) AllRules(ctx context.Context) (map[string]string, error) {
	rule := []entity.Rule{}
	err := dao.Rule.Ctx(ctx).Scan(&rule)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("getallrule", ""),
		))
		return nil, err
	}
	rst := map[string]string{}
	for _, i := range rule {
		rst[i.RuleId] = i.Rules
	}
	return rst, nil
}

func (s *sDB) subscription(conn *pgxpool.Conn, name string, notificationChannel chan *pgconn.Notification) {

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
							rules, _ := s.GetRules(s.ctx, ruleId)
							service.LEngine().UpRules(ruleId, rules)
						}
						if op == "rm" {
							service.LEngine().UpRules(ruleId, "")
						}
					}
				case AbiChName:
				}

				g.Log().Notice(context.TODO(), "Received notification:", notification)
			}
		}
	}()
}

func (s *sDB) listenNotify(subNames []string) {
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
