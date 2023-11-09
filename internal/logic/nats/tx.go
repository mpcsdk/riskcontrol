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

func (*sNats) PerformRiskTxs(ctx context.Context, req *v1.TxRiskReq) (res *v1.TxRiskRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "performRiskTxs")
	defer span.End()
	///
	serial, code := service.Risk().RiskTxs(ctx, req.UserId, req.SignTxData)
	if code == mpccode.RiskCodeError {
		return &v1.TxRiskRes{
			Ok: code,
		}, nil
		// gerror.NewCode(mpccode.CodePerformRiskError)
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
	kinds, err := service.TFA().TFATx(ctx, req.UserId, serial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodePerformRiskError)
	}
	///
	g.Log().Notice(ctx, "PerformRiskTFA:", req.UserId, serial, kinds)
	return &v1.TxRiskRes{
		Ok:         code,
		RiskSerial: serial,
		//todo:
		RiskKind: kinds,
	}, err
}
