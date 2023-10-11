// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EthTx is the golang structure for table eth_tx.
type EthTx struct {
	Id          int64       `json:"id"          ` //
	CreatedAt   *gtime.Time `json:"createdAt"   ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` //
	DeletedAt   *gtime.Time `json:"deletedAt"   ` //
	Address     string      `json:"address"     ` //
	Contract    string      `json:"contract"    ` //
	MethodName  string      `json:"methodName"  ` //
	MethodSig   string      `json:"methodSig"   ` //
	EventName   string      `json:"eventName"   ` //
	EventSig    string      `json:"eventSig"    ` //
	Topics      string      `json:"topics"      ` //
	From        string      `json:"from"        ` //
	To          string      `json:"to"          ` //
	Value       string      `json:"value"       ` //
	Kind        string      `json:"kind"        ` //
	BlockNumber int64       `json:"blockNumber" ` //
	BlockHash   string      `json:"blockHash"   ` //
	TxHash      string      `json:"txHash"      ` //
	TxIndex     int64       `json:"txIndex"     ` //
	LogIndex    int64       `json:"logIndex"    ` //
	Data        string      `json:"data"        ` //
	Name        string      `json:"name"        ` //
}
