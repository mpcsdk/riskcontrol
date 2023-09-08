package rulesdb

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	// _ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// rule "mud"
// begein
//
//	if $USER.MUD.$Transfer24cnt > 1000
//		return false
//	else
//		return true
//
// end
// rule "nft"
// begin
//
//	if $USER.Token.Index.$Transfer24cnt > 5
//		return false
//	else
//		return true
//
// end
type sTxDb struct {
}

func (s *sTxDb) Set(ctx context.Context, ruleId, rules string) error {

	return nil
}
func (s *sTxDb) Get(ctx context.Context, ruleId string) (string, error) {
	rule := &entity.Rule{}
	err := dao.Rule.Ctx(ctx).Where(dao.Rule.Columns().RuleId, ruleId).Scan(rule)
	return rule.Rules, err
}

func new() *sTxDb {
	g.Redis().Exists(gctx.GetInitCtx())
	s := &sTxDb{}
	return s
}

func init() {
	// service.RegisterRulesDb(new())
}
