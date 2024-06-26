package riskctrl

import (
	"context"

	v1 "riskcontrol/api/riskctrl/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) TxsRequest(ctx context.Context, req *v1.TxsRequestReq) (*v1.TxsRequestRes, error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
	// g.Log().Notice(ctx, "TxsRequest:", "req:", req)
	// //limit
	// // if err := c.apiLimit(ctx, req.Token, "TxsRequest"); err != nil {
	// // 	return nil, err
	// // }
	// //trace
	// ctx, span := gtrace.NewSpan(ctx, "performRiskTxs")
	// defer span.End()
	// ///
	// res, err := service.NrpcClient().RiskTxs(ctx, &riskengine.TxRiskReq{
	// 	UserId: req.UserId,
	// 	SignTx: req.SignTx,
	// })
	// if err != nil {
	// 	g.Log().Warning(ctx, "RpcRiskTxs:", "req:", req, "err:", err)
	// 	return nil, mpccode.CodeInternalError()
	// }
	// ///
	// if res.Ok == mpccode.RiskCodeError {
	// 	return &v1.TxsRequestRes{
	// 		Code: res.Ok,
	// 	}, mpccode.CodeInternalError()
	// }
	// //
	// if res.Ok == mpccode.RiskCodePass {
	// 	return &v1.TxsRequestRes{
	// 		Code: res.Ok,
	// 	}, nil
	// }
	// if res.Ok == mpccode.RiskCodeForbidden {
	// 	return &v1.TxsRequestRes{
	// 		Code: res.Ok,
	// 	}, nil
	// }
	// ///
	// //
	// //notice:  tfatx  need verification
	// tfaInfo, err := service.DB().TfaDB().FetchTfaInfo(ctx, req.UserId)
	// if err != nil || tfaInfo == nil {
	// 	g.Log().Warning(ctx, "FetchTfaInfo:", "req:", req, "err:", err)
	// 	return nil, mpccode.CodeTFANotExist()
	// }
	// rst, err := service.TFA().TfaRequest(ctx, tfaInfo.UserId, tfaconst.RiskKind_Tx, nil)
	// if err != nil {
	// 	g.Log().Warning(ctx, "TfaRequest:", "req:", req, "err:", err)
	// 	return nil, mpccode.CodeInternalError()
	// }
	// // riskserial, kinds, err := service.TFA().RiskTxTidy(ctx, tfaInfo)
	// // if err != nil {
	// // 	return nil, mpccode.CodePerformRiskError()
	// // }
	// ///
	// g.Log().Notice(ctx, "RpcRiskTFA:", req.UserId, rst.RiskSerial)
	// return &v1.TxsRequestRes{
	// 	Code:       res.Ok,
	// 	RiskSerial: rst.RiskSerial,
	// 	VList:      rst.VList,
	// 	Msg:        res.Msg,
	// }, err
}
