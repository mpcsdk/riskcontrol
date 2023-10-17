// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// AggFt24H is the golang structure for table agg_ft_24h.
type AggFt24H struct {
	From       string          `json:"from"       ` //
	To         string          `json:"to"         ` //
	Value      decimal.Decimal `json:"value"      ` //
	Contract   string          `json:"contract"   ` //
	UpdatedAt  *gtime.Time     `json:"updatedAt"  ` //
	MethodName string          `json:"methodName" ` //
	FromBlock  int64           `json:"fromBlock"  ` //
	ToBlock    int64           `json:"toBlock"    ` //
	MethodSig  string          `json:"methodSig"  ` //
	FtName     string          `json:"ftName"     ` //
}
