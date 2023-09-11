package risk

import (
	"context"
	v1 "riskcontral/api/risk/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) PerformRisk(ctx context.Context, req *v1.RiskReq) (res *v1.RiskRes, err error) {
	if req.RuleName == "phone" {
		return nil, nil
	} else if req.RuleName == "mail" {
		return nil, gerror.NewCode(gcode.CodeNotImplemented)
	}

	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformSmsCode(ctx context.Context, req *v1.SmsCodeReq) (res *v1.SmsCodeRes, err error) {
	// service.SmsCode().SendCode(ctx)
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformMailCode(ctx context.Context, req *v1.MailCodekReq) (res *v1.MailCodekRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
