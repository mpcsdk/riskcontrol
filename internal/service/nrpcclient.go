// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"riskcontrol/api/riskctrl"
)

type (
	INrpcClient interface {
		Flush()
		// RiskTxs(ctx context.Context, req *riskengine.TxRiskReq) (res *riskengine.TxRiskRes, err error)
		RiskTfaRequest(ctx context.Context, req *riskctrl.TfaRequestReq) (res *riskctrl.TfaRequestRes, err error)
	}
)

var (
	localNrpcClient INrpcClient
)

func NrpcClient() INrpcClient {
	if localNrpcClient == nil {
		panic("implement not found for interface INrpcClient, forgot register?")
	}
	return localNrpcClient
}

func RegisterNrpcClient(i INrpcClient) {
	localNrpcClient = i
}
