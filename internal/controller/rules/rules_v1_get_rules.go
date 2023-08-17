package rules

import (
	"context"

	v1 "riskcontral/api/rules/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) GetRules(ctx context.Context, req *v1.GetRulesReq) (res *v1.GetRulesRes, err error) {

	rmap := service.LEngine().List(req.Name)
	res = &v1.GetRulesRes{
		Rules: rmap,
	}
	return res, nil
}
