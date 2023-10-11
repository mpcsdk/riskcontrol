// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	analzyer "riskcontral/common/ethtx/analzyer"
	"riskcontral/internal/model"
)

type (
	IEthTx interface {
		AnalzyTxs(ctx context.Context, signtxs *analzyer.SignTx) (*model.AnalzyTx, error)
		Data2Args(ctx context.Context, target string, data string) (map[string]interface{}, error)
	}
)

var (
	localEthTx IEthTx
)

func EthTx() IEthTx {
	if localEthTx == nil {
		panic("implement not found for interface IEthTx, forgot register?")
	}
	return localEthTx
}

func RegisterEthTx(i IEthTx) {
	localEthTx = i
}
