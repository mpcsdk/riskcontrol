package nrpcclient

import (
	"context"
	"riskcontral/api/riskctrl"
	"riskcontral/api/riskengine"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *sNrpcClient) RiskTxs(ctx context.Context, req *riskengine.TxRiskReq) (res *riskengine.TxRiskRes, err error) {
	res, err = c.riskengine.RpcRiskTxs(req)
	if err != nil {
		if err.Error() == mpccode.ErrNrpcTimeOut.Error() {
			g.Log().Warning(ctx, "RiskTxs TimeOut:", req)
			c.Flush()
			return nil, mpccode.CodeInternalError()
		}
		return nil, mpccode.CodeInternalError(err.Error())
	}
	return res, nil
}

// /
func (c *sNrpcClient) RiskTfaRequest(ctx context.Context, req *riskctrl.TfaRequestReq) (res *riskctrl.TfaRequestRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
