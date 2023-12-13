// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"riskcontral/api/risk/nrpc"

	"github.com/golang/protobuf/ptypes/empty"
)

type (
	INrpcServer interface {
		RpcContractAbiBriefs(ctx context.Context, req *nrpc.ContractAbiBriefsReq) (res *nrpc.ContractAbiBriefsRes, err error)
		RpcContractAbi(ctx context.Context, req *nrpc.ContractAbiReq) (res *nrpc.ContractAbiRes, err error)
		NatsPub()
		// /
		RpcContractRuleBriefs(ctx context.Context, req *nrpc.ContractRuleBriefsReq) (res *nrpc.ContractRuleBriefsRes, err error)
		RpcContractRule(ctx context.Context, req *nrpc.ContractRuleReq) (res *nrpc.ContractRuleRes, err error)
		RpcAlive(ctx context.Context, in *empty.Empty) (*empty.Empty, error)
		RpcRiskTFA(ctx context.Context, req *nrpc.TFARiskReq) (res *nrpc.TFARiskRes, err error)
		RpcRiskTxs(ctx context.Context, req *nrpc.TxRiskReq) (res *nrpc.TxRiskRes, err error)
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
