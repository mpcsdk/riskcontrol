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
		GetRules(ctx context.Context, ruleId string) (string, error)
		AllRules(ctx context.Context) map[string]string
		GetNftRules(ctx context.Context) (map[string]*mpcmodel.NftRule, error)
		GetFtRules(ctx context.Context) (map[string]*mpcmodel.FtRule, error)
		TfaMailExists(ctx context.Context, mail string) error
		TfaPhoneExists(ctx context.Context, phone string) error
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
