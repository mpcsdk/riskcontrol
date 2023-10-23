// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"math/big"
)

type (
	IRPC interface {
		RpcNftCnt(ctx context.Context, addr string, contract string, method string) (int, error)
		RpcFtCnt(ctx context.Context, addr string, contract string, method string) (*big.Int, error)
		RpcAlive(ctx context.Context) error
	}
)

var (
	localRPC IRPC
)

func RPC() IRPC {
	if localRPC == nil {
		panic("implement not found for interface IRPC, forgot register?")
	}
	return localRPC
}

func RegisterRPC(i IRPC) {
	localRPC = i
}
