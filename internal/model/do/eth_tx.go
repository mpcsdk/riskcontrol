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
	g.Meta      `orm:"table:eth_tx, do:true"`
	Id          interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
	Address     interface{} //
	Contract    interface{} //
	MethodName  interface{} //
	MethodSig   interface{} //
	EventName   interface{} //
	EventSig    interface{} //
	Topics      interface{} //
	From        interface{} //
	To          interface{} //
	Value       interface{} //
	Kind        interface{} //
	BlockNumber interface{} //
	BlockHash   interface{} //
	TxHash      interface{} //
	TxIndex     interface{} //
	LogIndex    interface{} //
	Data        interface{} //
	Name        interface{} //
}
