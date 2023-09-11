// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"riskcontral/internal/model/entity"
)

type (
	ITFA interface {
		TFAInfo(ctx context.Context, token string) (*entity.Tfa, error)
		UpPhone(ctx context.Context, token string, phone string) error
		UpMail(ctx context.Context, token string, mail string) error
		SendPhoneCode(ctx context.Context, token string, opt string) (string, error)
		SendMailOTP(ctx context.Context, token string, opt string) (string, error)
		VerifyCode(ctx context.Context, token string, kind, code string) error
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
