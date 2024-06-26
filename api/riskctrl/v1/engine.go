package v1

import "github.com/gogf/gf/v2/frame/g"

type RiskTxsReq struct {
	g.Meta `path:"/RiskTxs" tags:"RiskTxs" method:"post" summary:"RiskTxs"`
	UserId string `json:"userId"`
	SignTx string `json:"signTx"`
}
type RiskTxsRes struct {
	Code int    `json:"code"`
	Msg  string `json"msg"`
}

type StateReq struct {
	g.Meta   `path:"/state" tags:"state" method:"post" summary:"state"`
	CodeType int `json:"codeType"`
}
type StateRes struct {
	Msg interface{} `json"msg"`
}
