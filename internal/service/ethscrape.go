// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"math/big"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	IEthEventGeter interface {
		InitByService() error
		Stop()
		RunBySerivce()
		Run(lastBlock int64, abis []*entity.ContractAbi, nftRules map[string]*model.NftRule, ftRules map[string]*model.FtRule)
		GetBlockNumber(ctx context.Context) (int64, error)
		GetLogs(ctx context.Context, from, to *big.Int, addresses []common.Address, topic [][]common.Hash) ([]*types.Log, error)
		FilterTx(ctx context.Context, logs []*types.Log, nftrules map[string]*model.NftRule, ftrules map[string]*model.FtRule) []*entity.EthTx
	}
)

var (
	localEthEventGeter IEthEventGeter
)

func EthEventGeter() IEthEventGeter {
	if localEthEventGeter == nil {
		panic("implement not found for interface IEthEventGeter, forgot register?")
	}
	return localEthEventGeter
}

func RegisterEthEventGeter(i IEthEventGeter) {
	localEthEventGeter = i
}
