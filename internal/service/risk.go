// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"
)

type (
	IRisk interface {
		RiskTxs(ctx context.Context, userId string, signTx string) (string, int32)
		RiskTFA(ctx context.Context, tfaInfo *entity.Tfa, riskData *model.RiskTfa) (string, int32)
	}
)

var (
	localRisk IRisk
)

func Risk() IRisk {
	if localRisk == nil {
		panic("implement not found for interface IRisk, forgot register?")
	}
	return localRisk
}

func RegisterRisk(i IRisk) {
	localRisk = i
}
