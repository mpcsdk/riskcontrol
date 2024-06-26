package nrpcserver

import (
	"context"
	"riskcontrol/api/riskctrl"
	v1 "riskcontrol/api/tfa/v1"
	"riskcontrol/internal/logic/tfa/tfaconst"
	"riskcontrol/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

func (*NrpcServer) RpcTxsRequest(ctx context.Context, req *riskctrl.TxRequestReq) (*riskctrl.TxRequestRes, error) {
	g.Log().Notice(ctx, "RpcRiskTxs:", "req:", req)
	//trace
	ctx, span := gtrace.NewSpan(ctx, "performRiskTxs")
	defer span.End()

	//
	//notice:  tfatx  need verification
	// tfaInfo, err := service.DB().TfaDB().FetchTfaInfo(ctx, req.UserId)
	// if err != nil || tfaInfo == nil {
	// 	return nil, mpccode.CodeTFANotExist()
	// }
	res, err := service.TFA().TfaRequest(ctx, req.UserId, tfaconst.RiskKind_Tx, &v1.RequestData{
		UserId:      req.UserId,
		ChainId:     req.ChainId,
		SignDataStr: req.SignTxData,
	})
	if err != nil {
		g.Log().Warning(ctx, "TfaRequest:", "req:", req, "err:", err)
		return nil, err
	}
	///
	g.Log().Notice(ctx, "RpcRiskTFA:", req.UserId, res.RiskSerial)
	return &riskctrl.TxRequestRes{
		Ok:         res.Ok,
		RiskSerial: res.RiskSerial,
		RiskKind:   res.VList,
	}, nil
}
