package riskctrl

import (
	"context"

	v1 "riskcontral/api/riskctrl/v1"
	"riskcontral/api/riskengine"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) TxsRequest(ctx context.Context, req *v1.TxsRequestReq) (*v1.TxsRequestRes, error) {
	g.Log().Notice(ctx, "TxsRequest:", "req:", req)
	//limit
	// if err := c.apiLimit(ctx, req.Token, "TxsRequest"); err != nil {
	// 	return nil, err
	// }
	//trace
	ctx, span := gtrace.NewSpan(ctx, "performRiskTxs")
	defer span.End()
	///
	res, err := service.NrpcClient().RiskTxs(ctx, &riskengine.TxRiskReq{
		UserId: req.UserId,
		SignTx: req.SignTx,
	})
	if err != nil {
		g.Log().Warning(ctx, "RpcRiskTxs:", "req:", req, "err:", err)
		return nil, mpccode.CodeInternalError()
	}
	///
	if res.Ok == mpccode.RiskCodeError {
		return &v1.TxsRequestRes{
			Code: res.Ok,
		}, mpccode.CodeInternalError()
	}
	//
	if res.Ok == mpccode.RiskCodePass {
		return &v1.TxsRequestRes{
			Code: res.Ok,
		}, nil
	}
	if res.Ok == mpccode.RiskCodeForbidden {
		return &v1.TxsRequestRes{
			Code: res.Ok,
		}, nil
	}
	///
	//
	//notice:  tfatx  need verification
	tfaInfo, err := service.DB().TfaDB().FetchTfaInfo(ctx, req.UserId)
	if err != nil || tfaInfo == nil {
		return nil, mpccode.CodeTFANotExist()
	}
	riskserial, kinds, err := service.TFA().RiskTxTidy(ctx, tfaInfo)
	if err != nil {
		return nil, mpccode.CodePerformRiskError()
	}
	///
	g.Log().Notice(ctx, "RpcRiskTFA:", req.UserId, riskserial)
	return &v1.TxsRequestRes{
		Code:       res.Ok,
		RiskSerial: riskserial,
		VList:      kinds,
		Msg:        res.Msg,
	}, err
}
