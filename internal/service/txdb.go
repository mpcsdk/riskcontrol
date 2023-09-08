// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ITxDb interface {
		Set(ctx context.Context, ruleId, rules string) error
		Get(ctx context.Context, ruleId string) (string, error)
	}
)

var (
	localTxDb ITxDb
)

func TxDb() ITxDb {
	if localTxDb == nil {
		panic("implement not found for interface ITxDb, forgot register?")
	}
	return localTxDb
}

func RegisterTxDb(i ITxDb) {
	localTxDb = i
}
