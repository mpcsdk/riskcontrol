// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AggFt24H is the golang structure of table agg_ft_24h for DAO operations like Where/Data.
type AggFt24H struct {
	g.Meta     `orm:"table:agg_ft_24h, do:true"`
	From       interface{} //
	To         interface{} //
	Value      interface{} //
	Contract   interface{} //
	UpdatedAt  *gtime.Time //
	MethodName interface{} //
	FromBlock  interface{} //
	ToBlock    interface{} //
	MethodSig  interface{} //
	FtName     interface{} //
}
