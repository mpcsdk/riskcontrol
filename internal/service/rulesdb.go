// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IRulesDb interface {
		Set(ruleId, rules string) error
		Get(ruleId string) (string, error)
		AllRules() map[string]string
		GetAbi(to string) (string, error)
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
