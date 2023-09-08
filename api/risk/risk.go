// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package risk

import (
	"context"
	
	"riskcontral/api/risk/v1"
)

type IRiskV1 interface {
	ExecRisk(ctx context.Context, req *v1.ExecRiskReq) (res *v1.ExecRiskRes, err error)
}


