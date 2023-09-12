// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IRulesDb interface {
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
		Get(ctx context.Context, ruleId string) (string, error)
		AllRules(ctx context.Context) map[string]string
		GetAbi(ctx context.Context, to string) (string, error)
	}
)

var (
	localRulesDb IRulesDb
)

func RulesDb() IRulesDb {
	if localRulesDb == nil {
		panic("implement not found for interface IRulesDb, forgot register?")
	}
	return localRulesDb
}

func RegisterRulesDb(i IRulesDb) {
	localRulesDb = i
}
