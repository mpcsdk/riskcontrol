// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

type (
	IDB interface {
		TfaDB() *mpcdao.RiskTfa
		RiskCtrl() *mpcdao.RiskCtrlRule
		ChainCfg() *mpcdao.ChainCfg
	}
)

var (
	localDB IDB
)

func DB() IDB {
	if localDB == nil {
		panic("implement not found for interface IDB, forgot register?")
	}
	return localDB
}

func RegisterDB(i IDB) {
	localDB = i
}
