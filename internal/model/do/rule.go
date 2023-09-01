// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Rule is the golang structure of table rule for DAO operations like Where/Data.
type Rule struct {
	g.Meta    `orm:"table:rule, do:true"`
	Id        interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	RuleId    interface{} //
	RuleDesc  interface{} //
	Rules     interface{} //
}
