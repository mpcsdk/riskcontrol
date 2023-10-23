// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"

	"github.com/mpcsdk/mpcCommon/mpcmodel"
)

type (
	IDB interface {
		GetAbi(ctx context.Context, addr string) (string, error)
		GetAbiAll(ctx context.Context) ([]*entity.ContractAbi, error)
		// /
		UpAggFT(ctx context.Context, ft *entity.AggFt24H) error
		UpAggNFT(ctx context.Context, nft *entity.AggNft24H) error
		GetAggFT(ctx context.Context, from, contract, methodName string) (*entity.AggFt24H, error)
		GetAggNFT(ctx context.Context, from, contract, methodName string) (int, error)
		GetRules(ctx context.Context, ruleId string) (string, error)
		AllRules(ctx context.Context) (map[string]string, error)
		GetNftRules(ctx context.Context) (map[string]*mpcmodel.NftRule, error)
		GetFtRules(ctx context.Context) (map[string]*mpcmodel.FtRule, error)
		TfaMailNotExists(ctx context.Context, mail string) error
		TfaPhoneNotExists(ctx context.Context, phone string) error
		InsertTfaInfo(ctx context.Context, userId string, data *do.Tfa) error
		// //
		UpdateTfaInfo(ctx context.Context, userId string, data *do.Tfa) error
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
