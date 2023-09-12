// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"riskcontral/common"
)

type (
	IUserInfo interface {
		GetUserInfo(ctx context.Context, userToken string) (userInfo *common.UserInfo, err error)
	}
)

var (
	localUserInfo IUserInfo
)

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement not found for interface IUserInfo, forgot register?")
	}
	return localUserInfo
}

func RegisterUserInfo(i IUserInfo) {
	localUserInfo = i
}
