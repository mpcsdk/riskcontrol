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
		SendCode(ctx context.Context, receiver, code string) error
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
