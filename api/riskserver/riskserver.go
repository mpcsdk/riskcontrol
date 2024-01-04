// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package riskserver

import (
	"context"
	
	"riskcontral/api/riskserver/v1"
)

type IRiskserverV1 interface {
	RiskTxs(ctx context.Context, req *v1.RiskTxsReq) (res *v1.RiskTxsRes, err error)
}


