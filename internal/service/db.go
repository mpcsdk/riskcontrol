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
	IDB interface {
		// /
		UpAggFT(ctx context.Context, ft *entity.AggFt24H) error
		UpAggNFT(ctx context.Context, nft *entity.AggNft24H) error
		GetAggFT(ctx context.Context, from, contract, methodName string) (*entity.AggFt24H, error)
		GetAggNFT(ctx context.Context, from, contract, methodName string) (int, error)
		GetContractAbiBriefs(ctx context.Context, SceneNo string, kind string) ([]*entity.Contractabi, error)
		// /
		GetContractAbi(ctx context.Context, SceneNo string, address string) (*entity.Contractabi, error)
		GetContractRuleBriefs(ctx context.Context, SceneNo string, kind string) ([]*entity.Contractrule, error)
		// /
		GetContractRule(ctx context.Context, SceneNo string, address string) (*entity.Contractrule, error)
	}
)

var (
	localDB IDB
)

func DB() IDB {
	if localDB == nil {
		panic("implement not found for interface IDB, forgot register?")
	}
	return localDB
}

func RegisterDB(i IDB) {
	localDB = i
}
