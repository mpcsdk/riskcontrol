// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"
)

type (
	IDB interface {
		GetAggFT(ctx context.Context, from, contract, methodName string) (*entity.AggFt24H, error)
		// /
		GetAggNFT(ctx context.Context, from, contract, methodName string) (int, error)
		GetContractAbiBriefs(ctx context.Context, chainId string, kind string) ([]*entity.Contractabi, error)
		// /
		GetContractAbi(ctx context.Context, chainId string, address string) (*entity.Contractabi, error)
		GetContractRuleBriefs(ctx context.Context, chainId string, kind string) ([]*entity.Contractrule, error)
		// /
		GetContractRule(ctx context.Context, chainId string, address string) (*entity.Contractrule, error)
		TfaMailNotExists(ctx context.Context, mail string) (bool, error)
		TfaPhoneNotExists(ctx context.Context, phone string) (bool, error)
		InsertTfaInfo(ctx context.Context, userId string, data *do.Tfa) error
		// //
		UpdateTfaInfo(ctx context.Context, userId string, data *do.Tfa) error
		ExistsTfaInfo(ctx context.Context, userId string) (bool, error)
		FetchTfaInfo(ctx context.Context, userId string) (*entity.Tfa, error)
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
