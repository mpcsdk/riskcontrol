package rulesdb

import (
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type sRulesDb struct {
	ctx g.Ctx
}

func (s *sRulesDb) Set(ruleId, rules string) error {
	// g.Redis().Set(s.ctx, name, rules)

	i, err := dao.Rule.Ctx(s.ctx).Data(do.Rule{RuleId: ruleId, Rules: rules}).Where(do.Rule{
		RuleId: ruleId,
	}).Count()
	if i == 0 {
		_, err = dao.Rule.Ctx(s.ctx).Data(do.Rule{RuleId: ruleId, Rules: rules}).Insert()
	} else {
		_, err = dao.Rule.Ctx(s.ctx).Data(do.Rule{RuleId: ruleId, Rules: rules}).Where(do.Rule{
			RuleId: ruleId,
		}).Update()
	}
	return err
}
func (s *sRulesDb) Get(ruleId string) (string, error) {
	// v, _ := g.Redis().Get(s.ctx, name)
	rule := &entity.Rule{}
	err := dao.Rule.Ctx(s.ctx).Where(dao.Rule.Columns().RuleId, ruleId).Scan(rule)
	return rule.Rules, err
}

func (s *sRulesDb) AllRules() map[string]string {
	rule := []entity.Rule{}
	dao.Rule.Ctx(s.ctx).Scan(&rule)
	rst := map[string]string{}
	for _, i := range rule {
		rst[i.RuleId] = i.Rules
	}
	return rst
}
func (s *sRulesDb) GetAbi(to string) (string, error) {

	contracts := &entity.ContractAbi{}
	err := dao.ContractAbi.Ctx(s.ctx).Where(dao.ContractAbi.Columns().Addr, to).Scan(contracts)
	return contracts.Abi, err
}

func new() *sRulesDb {
	g.Redis().Exists(gctx.GetInitCtx())
	return &sRulesDb{
		ctx: gctx.GetInitCtx(),
	}
}

func init() {
	service.RegisterRulesDb(new())
}
