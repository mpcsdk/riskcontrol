// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ISmsCode interface {
		// /
		SendVerificationCode(ctx context.Context, to string) (string, error)
		SendBindingPhoneCode(ctx context.Context, to string) (string, error)
		SendBindingCompletionPhone(ctx context.Context, to string) error
		// //
		SendUpPhoneCode(ctx context.Context, to string) (string, error)
		SendUpCompletionPhone(ctx context.Context, receiver string) error
	}
)

var (
	localSmsCode ISmsCode
)

func SmsCode() ISmsCode {
	if localSmsCode == nil {
		panic("implement not found for interface ISmsCode, forgot register?")
	}
	return localSmsCode
}

func RegisterSmsCode(i ISmsCode) {
	localSmsCode = i
}
