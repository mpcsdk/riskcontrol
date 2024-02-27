package nrpcserver

import (
	"context"
	"riskcontral/api/riskctrl"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

func (s *NrpcServer) RpcSendMailCode(ctx context.Context, req *riskctrl.SendMailCodeReq) (res *riskctrl.SendMailCodeRes, err error) {
	g.Log().Notice(ctx, "RpcSendMailCode:", "req:", req)
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	///
	err = service.TFA().SendMailCode(ctx, req.UserId, req.RiskSerial, req.Mail)
	return nil, err
}

func (s *NrpcServer) RpcSendPhoneCode(ctx context.Context, req *riskctrl.SendPhoneCodeReq) (res *riskctrl.SendPhoneCodeRes, err error) {
	g.Log().Notice(ctx, "RpcSendPhoneCode:", "req:", req)

	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendSmsCode")
	defer span.End()
	///
	///
	err = service.TFA().SendPhoneCode(ctx, req.UserId, req.RiskSerial, req.Phone)
	return nil, err
}
