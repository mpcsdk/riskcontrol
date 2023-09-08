package v1

import "github.com/gogf/gf/v2/frame/g"

type ExecRiskReq struct {
	g.Meta `path:"/exec" tags:"exec" method:"post" summary:"You first hello api"`
	Name   string                 `json:"name"`
	Param  map[string]interface{} `json:"param"`
}
type ExecRiskRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Result interface{} `json:"result"`
}

/////
