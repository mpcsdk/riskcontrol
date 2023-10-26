// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IMailCode interface {
		SendVerificationCode(ctx context.Context, to string) (string, error)
		SendBindingMailCode(ctx context.Context, to string) (string, error)
		SendBindingCompletionMail(ctx context.Context, to string) error
		// //
		SendUpMailCode(ctx context.Context, to string) (string, error)
		SendUpCompletionMail(ctx context.Context, to string) error
	}
)

var (
	localMailCode IMailCode
)

func MailCode() IMailCode {
	if localMailCode == nil {
		panic("implement not found for interface IMailCode, forgot register?")
	}
	return localMailCode
}

func RegisterMailCode(i IMailCode) {
	localMailCode = i
}
