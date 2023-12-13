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
	ITFA interface {
		GetRiskVerify(ctx context.Context, userId, riskSerial string) (risk *model.RiskVerifyPendding)
		// /
		SendPhoneCode(ctx context.Context, userId string, riskSerial string) (string, error)
		SendMailCode(ctx context.Context, userId string, riskSerial string) (string, error)
		TfaSetMail(ctx context.Context, tfaInfo *entity.Tfa, mail string, riskSerial string, riskKind model.RiskKind) (string, error)
		// //
		TfaBindMail(ctx context.Context, tfaInfo *entity.Tfa, mail string, riskSerial string) (string, error)
		TfaUpMail(ctx context.Context, tfaInfo *entity.Tfa, mail string, riskSerial string) (string, error)
		TfaSetPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string, riskKind model.RiskKind) (string, error)
		// //
		TfaBindPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string) (string, error)
		TfaUpPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string) (string, error)
		TfaRiskKind(ctx context.Context, tfaInfo *entity.Tfa, riskSerial string) (model.RiskKind, error)
		TfaRiskTidy(ctx context.Context, tfaInfo *entity.Tfa, riskSerial string, riskKind model.RiskKind) ([]string, error)
		TFATx(ctx context.Context, tfaInfo *entity.Tfa, riskSerial string) ([]string, error)
		VerifyCode(ctx context.Context, userId string, riskSerial string, code *model.VerifyCode) error
	}
)

var (
	localTFA ITFA
)

func TFA() ITFA {
	if localTFA == nil {
		panic("implement not found for interface ITFA, forgot register?")
	}
	return localTFA
}

func RegisterTFA(i ITFA) {
	localTFA = i
}
