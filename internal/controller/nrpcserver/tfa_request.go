package nrpcserver

import (
	"context"
	"riskcontrol/api/riskctrl"
	"riskcontrol/internal/logic/tfa/tfaconst"
	"riskcontrol/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *NrpcServer) RpcTfaRequest(ctx context.Context, req *riskctrl.TfaRequestReq) (*riskctrl.TfaRequestRes, error) {
	g.Log().Notice(ctx, "RpcTfaRequest:", "req:", req)
	//trace
	ctx, span := gtrace.NewSpan(ctx, "TfaRequest")
	defer span.End()
	/////
	info, _ := service.UserInfo().GetUserInfo(ctx, req.Token)
	if info == nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	///
	if req.CodeType == tfaconst.CodeType_TxNeedVerify {
		return nil, mpccode.CodeParamInvalid()
	}
	//
	riskKind := tfaconst.CodeType2RiskKind(req.CodeType)
	res, err := service.TFA().TfaRequest(ctx, info.UserId, riskKind, nil)
	if err != nil {
		return nil, err
	}
	return &riskctrl.TfaRequestRes{
		RiskSerial: res.RiskSerial,
		VList:      res.VList,
	}, nil
	///
}
