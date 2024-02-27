package nrpcserver

import (
	"context"
	"riskcontral/api/riskctrl"
	"riskcontral/internal/model"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

func (s *NrpcServer) RpcVerifyCode(ctx context.Context, req *riskctrl.VerifyCodeReq) (res *riskctrl.VerifyCodeRes, err error) {

	g.Log().Notice(ctx, "RpcVerifyCode:", "req:", req)
	//trace
	ctx, span := gtrace.NewSpan(ctx, "VerifyCode")
	defer span.End()

	code := &model.VerifyCode{
		PhoneCode: req.PhoneCode,
		MailCode:  req.MailCode,
	}

	err = service.TFA().VerifyCode(ctx, req.UserId, req.RiskSerial, code)
	if err != nil {
		g.Log().Warning(ctx, "VerifyCode", req, err)
		return nil, err
	}

	return nil, err
}
