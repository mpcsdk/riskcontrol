// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package rules

import (
	"context"
	
	"riskcontral/api/rules/v1"
)

type IRulesV1 interface {
	GetRules(ctx context.Context, req *v1.GetRulesReq) (res *v1.GetRulesRes, err error)
	UpRule(ctx context.Context, req *v1.UpRuleReq) (res *v1.UpRuleRes, err error)
	ExecRule(ctx context.Context, req *v1.ExecRuleReq) (res *v1.ExecRuleRes, err error)
}


