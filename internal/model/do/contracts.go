// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Contracts is the golang structure of table contracts for DAO operations like Where/Data.
type Contracts struct {
	g.Meta `orm:"table:contracts, do:true"`
	Addr   interface{} //
	Abi    interface{} //
}
