// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tfa is the golang structure for table tfa.
type Tfa struct {
	UserId         string      `json:"userId"         ` //
	CreatedAt      *gtime.Time `json:"createdAt"      ` //
	DeletedAt      *gtime.Time `json:"deletedAt"      ` //
	Phone          string      `json:"phone"          ` //
	Mail           string      `json:"mail"           ` //
	PhoneUpdatedAt *gtime.Time `json:"phoneUpdatedAt" ` //
	MailUpdatedAt  *gtime.Time `json:"mailUpdatedAt"  ` //
	TokenData string `json:"tokenData"`
}
