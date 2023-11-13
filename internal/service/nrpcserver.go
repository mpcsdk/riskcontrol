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
		RpcAlive(ctx context.Context, in *empty.Empty) (*empty.Empty, error)
		RpcAllAbi(ctx context.Context, req *v1.AllAbiReq) (res *v1.AllAbiRes, err error)
		RpcAllNftRules(ctx context.Context, req *v1.NftRulesReq) (res *v1.NftRulesRes, err error)
		RpcAllFtRules(ctx context.Context, req *v1.FtRulesReq) (res *v1.FtRulesRes, err error)
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
