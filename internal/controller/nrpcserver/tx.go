package nats

import (
	"context"
	"riskcontral/api/risk/nrpc"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (*NrpcServer) RpcRiskTxs(ctx context.Context, req *nrpc.TxRiskReq) (res *nrpc.TxRiskRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "performRiskTxs")
	defer span.End()
	///
	serial, code := service.Risk().RiskTxs(ctx, req.UserId, req.SignTxData)
	if code == mpccode.RiskCodeError {
		return &nrpc.TxRiskRes{
			Ok: code,
		}, nil
		// gerror.NewCode(mpccode.CodeRpcRiskError)
	}
	///: pass or forbidden
	//
	if code == mpccode.RiskCodePass {
		return &nrpc.TxRiskRes{
			Ok: code,
		}, nil
	}
	if code == mpccode.RiskCodeForbidden {
		return &nrpc.TxRiskRes{
			Ok: code,
		}, nil
	}
	///
	//
	//notice:  tfatx  need verification
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
	if err != nil || tfaInfo == nil {
		g.Log().Warning(ctx, "SendSmsCode:", req, err)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	// kinds, err := service.TFA().TFATx(ctx, req.UserId, serial)
	kinds, err := service.TFA().TFATx(ctx, tfaInfo, serial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodePerformRiskError)
	}
	///
	g.Log().Notice(ctx, "RpcRiskTFA:", req.UserId, serial)
	return &nrpc.TxRiskRes{
		Ok:         code,
		RiskSerial: serial,
		//todo:
		RiskKind: kinds,
	}, err
}
