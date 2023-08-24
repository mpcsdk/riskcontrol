// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractAbi is the golang structure of table contract_abi for DAO operations like Where/Data.
type ContractAbi struct {
	g.Meta    `orm:"table:contract_abi, do:true"`
	Id        interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Addr      interface{} //
	Abi       interface{} //
}
