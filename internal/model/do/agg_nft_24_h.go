// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AggNft24H is the golang structure of table agg_nft_24h for DAO operations like Where/Data.
type AggNft24H struct {
	g.Meta     `orm:"table:agg_nft_24h, do:true"`
	From       interface{} //
	To         interface{} //
	Value      interface{} //
	Contract   interface{} //
	UpdatedAt  *gtime.Time //
	MethodName interface{} //
	FromBlock  interface{} //
	ToBlock    interface{} //
	MethodSig  interface{} //
	NftName    interface{} //
}
