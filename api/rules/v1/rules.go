package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetRulesReq struct {
	g.Meta `path:"/getrules" tags:"getrules" method:"post" summary:"You first hello api"`
	Name   string `json:"name"`
}
type GetRulesRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Rules  map[string]string `json:"rules"`
}

// /
type UpRuleReq struct {
	g.Meta `path:"/uprule" tags:"uprules" method:"post" summary:"You first hello api"`
	Name   string `json:"name"`
	Rule   string `json:"rule"`
}
type UpRuleRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// /
type ExecRuleReq struct {
	g.Meta `path:"/exec" tags:"exec" method:"post" summary:"You first hello api"`
	Name   string                 `json:"name"`
	Param  map[string]interface{} `json:"param"`
}
type ExecRuleRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Result interface{} `json:"result"`
}

/////
