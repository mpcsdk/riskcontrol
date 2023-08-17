package rules

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	v1 "riskcontral/api/rules/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) ExecRule(ctx context.Context, req *v1.ExecRuleReq) (res *v1.ExecRuleRes, err error) {
	param := map[string]interface{}{}
	for k, v := range req.Param {
		fmt.Println(reflect.TypeOf(v))
		fmt.Println(reflect.ValueOf(v).Type())
		n := v.(json.Number)
		param[k], _ = n.Int64()
	}
	rst, err := service.LEngine().Exec(req.Name, param)
	res = &v1.ExecRuleRes{
		Result: rst,
	}
	return res, err
}
