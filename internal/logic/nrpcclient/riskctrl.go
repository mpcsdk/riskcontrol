package nrpcclient

import (
	"context"
	"riskcontral/api/riskctrl"
)

func (c *sNrpcClient) RiskTxs(ctx context.Context, req *riskctrl.TxRiskReq) (res *riskctrl.TxRiskRes, err error) {
	return c.riskcli.RpcRiskTxs(req)
}
