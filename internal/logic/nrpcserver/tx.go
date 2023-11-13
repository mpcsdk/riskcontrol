package nats

import (
	"context"
	v1 "riskcontral/api/risk/nrpc/v1"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (*sNrpcServer) RpcRiskTxs(ctx context.Context, req *v1.TxRiskReq) (res *v1.TxRiskRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "performRiskTxs")
	defer span.End()
	///
	serial, code := service.Risk().RiskTxs(ctx, req.UserId, req.SignTxData)
	if code == mpccode.RiskCodeError {
		return &v1.TxRiskRes{
			Ok: code,
		}, nil
		// gerror.NewCode(mpccode.CodeRpcRiskError)
	}
	///: pass or forbidden
	//
	if code == mpccode.RiskCodePass {
		return &v1.TxRiskRes{
			Ok: code,
		}, nil
	}
	if code == mpccode.RiskCodeForbidden {
		return &v1.TxRiskRes{
			Ok: code,
		}, nil
	}
	///
	//
	//notice:  tfatx  need verification
	// tfaInfo, err := service.NrpcClient().RpcTfaInfo(ctx, req.UserId)
	// kinds, err := service.TFA().TFATx(ctx, req.UserId, serial)
	kinds, err := service.NrpcClient().RpcTfaTx(ctx, req.UserId, serial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodePerformRiskError)
	}
	///
	g.Log().Notice(ctx, "RpcRiskTFA:", req.UserId, serial)
	return &v1.TxRiskRes{
		Ok:         code,
		RiskSerial: serial,
		//todo:
		RiskKind: kinds,
	}, err
}
