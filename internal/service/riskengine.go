// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "riskcontrol/api/riskctrl/v1"
	"riskcontrol/internal/model"
)

type (
	IRiskEngine interface {
		Exec(ctx context.Context, ruleName string, data *model.RiskExecData) (int32, error)
		ExecTx(ctx context.Context, data *model.RiskExecData) (int32, error)
		ExecRule(ctx context.Context, ruleName string, data *model.RiskExecData) (int32, error)
		VerifyRules(ruleName, ruleStr string) error
		UpRules(ruleName, ruleDesc, ruleStr string, salience int) (string, error)
		// /
		AddApi(key string, api interface{})
		Stat(ctx context.Context, req *v1.StateReq) interface{}
	}
)

var (
	localRiskEngine IRiskEngine
)

func RiskEngine() IRiskEngine {
	if localRiskEngine == nil {
		panic("implement not found for interface IRiskEngine, forgot register?")
	}
	return localRiskEngine
}

func RegisterRiskEngine(i IRiskEngine) {
	localRiskEngine = i
}
