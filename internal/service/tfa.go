// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/model/entity"
)

type (
	ITFA interface {
		TFAInfo(ctx context.Context, userId string) (*entity.Tfa, error)
		// /
		SendPhoneCode(ctx context.Context, userId string, riskSerial string) (string, error)
		SendMailCode(ctx context.Context, userId string, riskSerial string) (string, error)
		TFACreate(ctx context.Context, userId string, phone string, mail string) (string, []string, error)
		TFAUpPhone(ctx context.Context, userId string, phone string) (string, error)
		TFAUpMail(ctx context.Context, userId string, mail string) (string, error)
		TFATx(ctx context.Context, userId string, riskSerial string) ([]string, error)
		VerifyCode(ctx context.Context, userId string, vreq []*v1.VerifyReq) error
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
