// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package riskctrl

import (
	"context"
	
	"riskcontrol/api/riskctrl/v1"
)

type IRiskctrlV1 interface {
	TxsRequest(ctx context.Context, req *v1.TxsRequestReq) (res *v1.TxsRequestRes, err error)
}


