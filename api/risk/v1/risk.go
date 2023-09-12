package v1

import "github.com/gogf/gf/v2/frame/g"

// 执行address和target对应的风控规则
type ExecRiskReq struct {
	g.Meta  `path:"/exec" tags:"exec" method:"post" summary:"You first hello api"`
	Token   string `json:"token"`
	Address string `json:"address"`
	Target  string `json:"target"`
	//
	From string `json:"from"`
	To   string `json:"to"`
	///
	Name  string                 `json:"name"`
	Param map[string]interface{} `json:"param"`
}
type ExecRiskRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Result interface{} `json:"result"`
}

/////
