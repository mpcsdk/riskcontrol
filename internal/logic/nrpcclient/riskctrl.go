package nrpcclient

import (
	"context"
	"riskcontral/api/riskctrl"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *sNrpcClient) RiskTxs(ctx context.Context, req *riskctrl.TxRiskReq) (res *riskctrl.TxRiskRes, err error) {
	res, err = c.riskcli.RpcRiskTxs(req)
	if err != nil {
		if err.Error() == mpccode.ErrNrpcTimeOut.Error() {
			g.Log().Warning(ctx, "RiskTxs TimeOut:", req)
			c.Flush()
			return nil, mpccode.CodeInternalError()
		}
		err = gerror.Wrap(mpccode.CodeInternalError(), mpccode.ErrDetails(
			mpccode.ErrDetail("req", req),
			mpccode.ErrDetail("err", err),
		))
		return nil, err
	}
	return res, err
}
func (c *sNrpcClient) RiskTfaRequest(ctx context.Context, req *riskctrl.TfaRiskReq) (res *riskctrl.TfaRiskRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
