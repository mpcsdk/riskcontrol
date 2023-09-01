// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractAbi is the golang structure for table contract_abi.
type ContractAbi struct {
	Id        int64       `json:"id"        ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
	Addr      string      `json:"addr"      ` //
	Abi       string      `json:"abi"       ` //
}
