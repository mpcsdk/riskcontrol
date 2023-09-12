package risk

import (
	"context"

	v1 "riskcontral/api/risk/v1"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/service"
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
	// param := map[string]interface{}{}
	// rst, err := service.LEngine().Exec(req.Name, param)
	// res = &v1.ExecRiskRes{
	// 	Result: rst,
	// }

	riskData := &conrisk.RiskTx{
		Token:    req.Token,
		Address:  req.Address,
		Contract: req.Target,
		From:     req.From,
		To:       req.To,
	}

	rst, err := service.Risk().PerformRisk(ctx, "checkTx", riskData)
	if err != nil {
		return nil, err
	}
	res = &v1.ExecRiskRes{
		Result: rst,
	}
	return res, err
}
