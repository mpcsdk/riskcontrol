// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EthTx is the golang structure for table eth_tx.
type EthTx struct {
	Id         int64       `json:"id"         ` //
	CreatedAt  *gtime.Time `json:"createdAt"  ` //
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` //
	DeletedAt  *gtime.Time `json:"deletedAt"  ` //
	Address    string      `json:"address"    ` //
	Target     string      `json:"target"     ` //
	MethodId   string      `json:"methodId"   ` //
	MethodName string      `json:"methodName" ` //
	Sig        string      `json:"sig"        ` //
	Data       string      `json:"data"       ` //
	Args       string      `json:"args"       ` //
	From       string      `json:"from"       ` //
	To         string      `json:"to"         ` //
	Value      string      `json:"value"      ` //
}
