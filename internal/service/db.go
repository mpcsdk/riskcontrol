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
		InsertTfaInfo(ctx context.Context, data *do.Tfa) error
		// //
		UpdateTfaInfo(ctx context.Context, data *do.Tfa) error
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
