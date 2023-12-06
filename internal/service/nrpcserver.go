// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "riskcontral/api/risk/nrpc/v1"

	"github.com/golang/protobuf/ptypes/empty"
)

type (
	INrpcServer interface {
		RpcContractAbiBriefs(ctx context.Context, req *v1.ContractAbiBriefsReq) (res *v1.ContractAbiBriefsRes, err error)
		RpcContractAbi(ctx context.Context, req *v1.ContractAbiReq) (res *v1.ContractAbiRes, err error)
		NatsPub()
		// /
		RpcContractRuleBriefs(ctx context.Context, req *v1.ContractRuleBriefsReq) (res *v1.ContractRuleBriefsRes, err error)
		RpcContractRule(ctx context.Context, req *v1.ContractRuleReq) (res *v1.ContractRuleRes, err error)
		RpcAlive(ctx context.Context, in *empty.Empty) (*empty.Empty, error)
		RpcRiskTFA(ctx context.Context, req *v1.TFARiskReq) (res *v1.TFARiskRes, err error)
		RpcRiskTxs(ctx context.Context, req *v1.TxRiskReq) (res *v1.TxRiskRes, err error)
	}
)

var (
	localNrpcServer INrpcServer
)

func NrpcServer() INrpcServer {
	if localNrpcServer == nil {
		panic("implement not found for interface INrpcServer, forgot register?")
	}
	return localNrpcServer
}

func RegisterNrpcServer(i INrpcServer) {
	localNrpcServer = i
}
