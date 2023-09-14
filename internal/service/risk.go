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
		RiskPhoneCode(ctx context.Context, riskserial string) error
		RiskMailCode(ctx context.Context, riskserial string) error
		VerifyCode(ctx context.Context, serial string, code string) error
		PerformRiskTxs(ctx context.Context, userId string, address string, txs []*conrisk.RiskTx) (string, int32, error)
		PerformRiskTFA(ctx context.Context, userId string, riskData *conrisk.RiskTfa) (string, error)
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
