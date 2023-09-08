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
		SendMailCode(ctx context.Context, to string, code string) (string, error)
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
