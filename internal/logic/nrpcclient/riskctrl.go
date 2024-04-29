package nrpcclient

import (
	"context"
	"riskcontral/api/riskctrl"
	"riskcontral/api/riskengine"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *sNrpcClient) RiskTxs(ctx context.Context, req *riskengine.TxRiskReq) (res *riskengine.TxRiskRes, err error) {
	res, err = c.riskengine.RpcRiskTxs(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// /
func (c *sNrpcClient) RiskTfaRequest(ctx context.Context, req *riskctrl.TfaRequestReq) (res *riskctrl.TfaRequestRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
