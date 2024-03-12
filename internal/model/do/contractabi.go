// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Contractabi is the golang structure of table contractabi for DAO operations like Where/Data.
type Contractabi struct {
	g.Meta          `orm:"table:contractabi, do:true"`
	Id              interface{} // ID
	CreateTime      *gtime.Time // 创建时间
	UpdateTime      *gtime.Time // 更新时间
	ContractName    interface{} // 合约名
	ContractAddress interface{} // 合约地址
	SceneNo         interface{} // 场景号
	AbiContent      interface{} // 合约abi
	ContractKind    interface{} // 合约类型
	ChainId         interface{} // 链id
}
