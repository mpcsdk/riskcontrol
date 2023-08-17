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

func (s *sRulesDb) Set(name, rules string) error {
	// g.Redis().Set(s.ctx, name, rules)

	i, err := dao.Rules.Ctx(s.ctx).Data(do.Rules{Name: name, Rule: rules}).Where(do.Rules{
		Name: name,
	}).Count()
	if i == 0 {
		_, err = dao.Rules.Ctx(s.ctx).Data(do.Rules{Name: name, Rule: rules}).Insert()
	} else {
		_, err = dao.Rules.Ctx(s.ctx).Data(do.Rules{Name: name, Rule: rules}).Where(do.Rules{
			Name: name,
		}).Update()
	}
	return err
}
func (s *sRulesDb) Get(name string) (string, error) {
	// v, _ := g.Redis().Get(s.ctx, name)
	rule := &entity.Rules{}
	err := dao.Rules.Ctx(s.ctx).Where(dao.Rules.Columns().Name, name).Scan(rule)
	return rule.Rule, err
}

func (s *sRulesDb) AllRules() map[string]string {
	rule := []entity.Rules{}
	dao.Rules.Ctx(s.ctx).Scan(&rule)
	rst := map[string]string{}
	for _, i := range rule {
		rst[i.Name] = i.Rule
	}
	return rst
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
