// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package risk

import (
	"riskcontral/api/risk"
	v1 "riskcontral/api/risk/v1"
	"context"
)
type ControllerV1 struct{}

func NewV1() risk.IRiskV1 {
	return &ControllerV1{}
}

func (*ControllerV1) ExecRisk(ctx context.Context, req *v1.ExecRiskReq) (res *v1.ExecRiskRes, err error) {
	return nil, nil
}
