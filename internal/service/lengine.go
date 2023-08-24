// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	ILEngine interface {
		UpRules(ruleId, rules string) error
		Exec(ruleId string, param map[string]interface{}) (bool, error)
		List(ruleId string) map[string]string
	}
)

var (
	localLEngine ILEngine
)

func LEngine() ILEngine {
	if localLEngine == nil {
		panic("implement not found for interface ILEngine, forgot register?")
	}
	return localLEngine
}

func RegisterLEngine(i ILEngine) {
	localLEngine = i
}
