// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"riskcontral/internal/consts/conrisk"
)

type (
	IRisk interface {
		PerformRiskTxs(ctx context.Context, userId string, signTx string) (string, int32)
		PerformRiskTFA(ctx context.Context, userId string, riskData *conrisk.RiskTfa) (string, int32)
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
