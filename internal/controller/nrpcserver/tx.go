package nrpcserver

import (
	"context"
	"riskcontral/api/riskctrl"
	"riskcontral/api/riskengine"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (*NrpcServer) RpcTxsRequest(ctx context.Context, req *riskctrl.TxRequestReq) (*riskctrl.TxRequestRes, error) {
	g.Log().Notice(ctx, "RpcRiskTxs:", "req:", req)
	//trace
	ctx, span := gtrace.NewSpan(ctx, "performRiskTxs")
	defer span.End()
	///
	res, err := service.NrpcClient().RiskTxs(ctx, &riskengine.TxRiskReq{
		UserId:  req.UserId,
		SignTx:  req.SignTxData,
		ChainId: req.ChainId,
	})
	if err != nil {
		g.Log().Warning(ctx, "RpcRiskTxs:", "req:", req, "err:", err)
		return nil, gerror.Wrap(mpccode.CodeInternalError(), mpccode.ErrDetails(
			mpccode.ErrDetail("engineErr", err.Error()),
		))
	}
	///
	if res.Ok == mpccode.RiskCodeError {
		return &riskctrl.TxRequestRes{
			Ok: res.Ok,
		}, mpccode.CodeInternalError()
	}
	//
	if res.Ok == mpccode.RiskCodePass {
		return &riskctrl.TxRequestRes{
			Ok: res.Ok,
		}, nil
	}
	if res.Ok == mpccode.RiskCodeForbidden {
		return &riskctrl.TxRequestRes{
			Ok: res.Ok,
		}, nil
	}
	///
	//
	//notice:  tfatx  need verification
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
	if err != nil || tfaInfo == nil {
		return nil, mpccode.CodeTFANotExist()
	}
	riskserial, kinds, err := service.TFA().RiskTxTidy(ctx, tfaInfo)
	if err != nil {
		return nil, mpccode.CodePerformRiskError()
	}
	///
	g.Log().Notice(ctx, "RpcRiskTFA:", req.UserId, riskserial)
	return &riskctrl.TxRequestRes{
		Ok:         res.Ok,
		RiskSerial: riskserial,
		RiskKind:   kinds,
	}, err
}
