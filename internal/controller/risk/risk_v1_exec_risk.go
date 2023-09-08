package risk

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "riskcontral/api/risk/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) ExecRisk(ctx context.Context, req *v1.ExecRiskReq) (res *v1.ExecRiskRes, err error) {
	param := map[string]interface{}{}
	rst, err := service.LEngine().Exec(req.Name, param)
	res = &v1.ExecRiskRes{
		Result: rst,
	}
	return res, err
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
