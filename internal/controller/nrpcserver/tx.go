package nrpcserver

import (
	"context"
	"riskcontral/api/riskctrl"
	"riskcontral/api/riskengine"
	"riskcontral/internal/logic/tfa/tfaconst"
	"riskcontral/internal/service"

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
		return nil, err
	}
	///
	if res.Ok == mpccode.RiskCodeError {
		return &riskctrl.TxRequestRes{
			Ok: mpccode.RiskCodeError,
		}, err
	}
	//
	if res.Ok == mpccode.RiskCodePass {
		return &riskctrl.TxRequestRes{
			Ok: mpccode.RiskCodePass,
		}, nil
	}
	if res.Ok == mpccode.RiskCodeForbidden {
		return &riskctrl.TxRequestRes{
			Ok: mpccode.RiskCodeForbidden,
		}, nil
	}
	///
	//
	//notice:  tfatx  need verification
	tfaInfo, err := service.DB().TfaDB().FetchTfaInfo(ctx, req.UserId)
	if err != nil || tfaInfo == nil {
		return nil, mpccode.CodeTFANotExist()
	}
	rst, err := service.TFA().TfaRequest(ctx, tfaInfo.UserId, tfaconst.RiskKind_Tx, nil)
	if err != nil {
		g.Log().Warning(ctx, "TfaRequest:", "req:", req, "err:", err)
		return nil, mpccode.CodeInternalError()
	}
	///
	g.Log().Notice(ctx, "RpcRiskTFA:", req.UserId, rst.RiskSerial)
	return &riskctrl.TxRequestRes{
		Ok:         res.Ok,
		RiskSerial: rst.RiskSerial,
		RiskKind:   rst.VList,
	}, nil
}
