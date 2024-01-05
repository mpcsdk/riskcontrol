package riskserver

import (
	"context"

	"riskcontral/api/riskserver"
	v1 "riskcontral/api/riskserver/v1"
)

func (c *ControllerV1) RiskTxs(ctx context.Context, req *v1.RiskTxsReq) (*v1.RiskTxsRes, error) {
	res, err := c.nrpc.RpcRiskTxs(ctx, &riskserver.TxRiskReq{
		UserId:     req.UserId,
		SignTxData: req.SignTx,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RiskTxsRes{
		Code:       int(res.Ok),
		RiskSerial: res.RiskSerial,
		VList:      res.RiskKind,
		Msg:        res.Msg,
	}, nil
}
