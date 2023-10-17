package risk

import (
	"context"

	v1 "riskcontral/api/risk/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// @Summary 验证token，执行交易风控,
// @Tags 风控
// @Accept
// @Produce
// @Param
// @Success 200 {object} riskcontral/internal/service.RiskTx
// @Failure 200 {object} riskcontral/internal/service.RiskTx
// @Router /v1/risk/execrisk [post]
func (c *ControllerV1) ExecRisk(ctx context.Context, req *v1.ExecRiskReq) (res *v1.ExecRiskRes, err error) {
	//trace
	// ctx, span := gtrace.NewSpan(ctx, "ExecRisk")
	// defer span.End()
	// //
	return nil, gerror.NewCode(gcode.CodeNotFound)
}
