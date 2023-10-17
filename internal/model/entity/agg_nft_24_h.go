// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AggNft24H is the golang structure for table agg_nft_24h.
type AggNft24H struct {
	From       string      `json:"from"       ` //
	To         string      `json:"to"         ` //
	Value      int64       `json:"value"      ` //
	Contract   string      `json:"contract"   ` //
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` //
	MethodName string      `json:"methodName" ` //
	FromBlock  int64       `json:"fromBlock"  ` //
	ToBlock    int64       `json:"toBlock"    ` //
	MethodSig  string      `json:"methodSig"  ` //
	NftName    string      `json:"nftName"    ` //
}
