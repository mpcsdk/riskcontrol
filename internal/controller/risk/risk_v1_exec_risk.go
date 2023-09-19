package risk

import (
	"context"

	v1 "riskcontral/api/risk/v1"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/net/gtrace"
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
	ctx, span := gtrace.NewSpan(ctx, "ExecRisk")
	defer span.End()
	//

	riskData := &conrisk.RiskTx{
		Token:    req.Token,
		Address:  req.Address,
		Contract: req.Target,
		From:     req.From,
		To:       req.To,
	}

	_, rst, err := service.Risk().PerformRiskTxs(ctx, "userId", req.Address, []*conrisk.RiskTx{riskData})
	if err != nil {
		return nil, err
	}
	res = &v1.ExecRiskRes{
		Result: rst,
	}
	return res, err
}
