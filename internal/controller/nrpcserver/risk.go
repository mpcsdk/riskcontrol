package nats

import (
	"context"
	"riskcontral/api/riskctrl"
	"riskcontral/api/riskserver"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/rand"
)

func (*NrpcServer) RpcRiskTxs(ctx context.Context, req *riskserver.TxRiskReq) (*riskserver.TxRiskRes, error) {
	g.Log().Notice(ctx, "RpcRiskTxs:", "req:", req)
	//trace
	ctx, span := gtrace.NewSpan(ctx, "performRiskTxs")
	defer span.End()
	///
	res, err := service.NrpcClient().RiskTxs(ctx, &riskctrl.TxRiskReq{
		UserId: req.UserId,
		SignTx: req.SignTxData,
	})
	if err != nil {
		g.Log().Warning(ctx, "RpcRiskTxs:", "req:", req, "err:", err)
		return nil, mpccode.CodeInternalError()
	}
	riskserial := rand.GenNewSid()
	///
	if res.Ok == mpccode.RiskCodeError {
		return &riskserver.TxRiskRes{
			Ok: res.Ok,
		}, mpccode.CodeInternalError()
		// gerror.NewCode(mpccode.CodeRpcRiskError)
	}
	///: pass or forbidden
	//
	if res.Ok == mpccode.RiskCodePass {
		return &riskserver.TxRiskRes{
			Ok: res.Ok,
		}, nil
	}
	if res.Ok == mpccode.RiskCodeForbidden {
		return &riskserver.TxRiskRes{
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
	// kinds, err := service.TFA().TFATx(ctx, req.UserId, serial)
	kinds, err := service.TFA().TFATx(ctx, tfaInfo, riskserial)
	if err != nil {
		return nil, mpccode.CodePerformRiskError()
	}
	///
	g.Log().Notice(ctx, "RpcRiskTFA:", req.UserId, riskserial)
	return &riskserver.TxRiskRes{
		Ok:         res.Ok,
		RiskSerial: riskserial,
		//todo:
		RiskKind: kinds,
	}, err
}
