package nrpcserver

import (
	"context"
	"riskcontrol/api/riskctrl"
	"riskcontrol/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

func (s *NrpcServer) RpcSendMailCode(ctx context.Context, req *riskctrl.SendMailCodeReq) (res *riskctrl.SendMailCodeRes, err error) {
	g.Log().Notice(ctx, "RpcSendMailCode:", "req:", req)
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	///
	err = service.TFA().SendMailCode(ctx, req.UserId, req.RiskSerial)
	return nil, err
}

func (s *NrpcServer) RpcSendPhoneCode(ctx context.Context, req *riskctrl.SendPhoneCodeReq) (res *riskctrl.SendPhoneCodeRes, err error) {
	g.Log().Notice(ctx, "RpcSendPhoneCode:", "req:", req)

	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendSmsCode")
	defer span.End()
	///
	///
	err = service.TFA().SendPhoneCode(ctx, req.UserId, req.RiskSerial)
	return nil, err
}

// 资产查询方法描述：查询多链上资产转移统计，
// queryAssetCnt(chainIds list, from string, contract string, start timeStemp, end timeStamp) big.Int
// 参数：
// chainIds：链id列表,
// from：from地址
// contract：合约地址
// start：开始时间
// end：结束时间
// 返回：ft转移总数，nft转移次数
