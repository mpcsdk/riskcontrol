package riskserver

import (
	"context"

	"riskcontral/api/riskctrl"
	v1 "riskcontral/api/riskserver/v1"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) RiskTxs(ctx context.Context, req *v1.RiskTxsReq) (*v1.RiskTxsRes, error) {
	res, err := service.NrpcClient().RiskTxs(ctx, &riskctrl.TxRiskReq{
		UserId: req.UserId,
		SignTx: req.SignTx,
	})
	if err != nil {
		g.Log().Warning(ctx, err)
		return nil, err
	}
	///
	return &v1.RiskTxsRes{
		Code: int(res.Ok),
		Msg:  res.Msg,
	}, nil
}
