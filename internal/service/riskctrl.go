// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type (
	IRiskCtrl interface {
		RiskCtrlTx(ctx context.Context, userId string, tfaInfo *entity.Tfa, signDataStr string, chainId string) (int32, error)
	}
)

var (
	localRiskCtrl IRiskCtrl
)

func RiskCtrl() IRiskCtrl {
	if localRiskCtrl == nil {
		panic("implement not found for interface IRiskCtrl, forgot register?")
	}
	return localRiskCtrl
}

func RegisterRiskCtrl(i IRiskCtrl) {
	localRiskCtrl = i
}
