// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EthTx is the golang structure of table eth_tx for DAO operations like Where/Data.
type EthTx struct {
	g.Meta     `orm:"table:eth_tx, do:true"`
	Id         interface{} //
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	DeletedAt  *gtime.Time //
	Address    interface{} //
	Target     interface{} //
	MethodId   interface{} //
	MethodName interface{} //
	Sig        interface{} //
	Data       interface{} //
	Args       interface{} //
	From       interface{} //
	To         interface{} //
	Value      interface{} //
}
