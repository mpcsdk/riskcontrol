package rules

import (
	"context"

	v1 "riskcontral/api/rules/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) UpRule(ctx context.Context, req *v1.UpRuleReq) (res *v1.UpRuleRes, err error) {

	err = service.LEngine().UpRules(req.Name, req.Rule)
	return nil, err
}
