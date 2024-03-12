// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Contractabi is the golang structure for table contractabi.
type Contractabi struct {
	Id              int         `json:"id"              ` // ID
	CreateTime      *gtime.Time `json:"createTime"      ` // 创建时间
	UpdateTime      *gtime.Time `json:"updateTime"      ` // 更新时间
	ContractName    string      `json:"contractName"    ` // 合约名
	ContractAddress string      `json:"contractAddress" ` // 合约地址
	SceneNo         string      `json:"sceneNo"         ` // 场景号
	AbiContent      string      `json:"abiContent"      ` // 合约abi
	ContractKind    string      `json:"contractKind"    ` // 合约类型
	ChainId         string      `json:"chainId"         ` // 链id
}
