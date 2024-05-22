// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/logic/tfa/tfaconst"
	"riskcontral/internal/model"
)

type (
	ITFA interface {
		SendMailCode(ctx context.Context, userId string, riskSerial string) error
		// /
		SendPhoneCode(ctx context.Context, userId string, riskSerial string) error
		TfaRequest(ctx context.Context, userId string, riskKind tfaconst.RISKKIND, data *v1.RequestData) (*v1.TfaRequestRes, error)
		TfaTidyMail(ctx context.Context, userId string, riskSerial string, mail string) error
		TfaTidyPhone(ctx context.Context, userId string, phone string, riskSerial string) error
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
