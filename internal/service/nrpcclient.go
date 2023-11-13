// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"riskcontral/internal/model/entity"
)

type (
	INrpcClient interface {
		Flush()
		RpcTfaTx(ctx context.Context, userId string, riskSerial string) ([]string, error)
		RpcTfaInfo(ctx context.Context, userId string) (*entity.Tfa, error)
		RpcAlive(ctx context.Context) error
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
