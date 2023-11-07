// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tfa is the golang structure of table tfa for DAO operations like Where/Data.
type Tfa struct {
	g.Meta         `orm:"table:tfa, do:true"`
	UserId         interface{} //
	CreatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
	Phone          interface{} //
	Mail           interface{} //
	PhoneUpdatedAt *gtime.Time //
	MailUpdatedAt  *gtime.Time //
	TokenData  interface{}
}
