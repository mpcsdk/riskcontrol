package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	//
	err = service.Risk().RiskMailCode(ctx, req.RiskSerial)
	return nil, err
}

func (c *ControllerV1) VerifyMailCode(ctx context.Context, req *v1.VerifyMailCodeReq) (res *v1.VerifyMailCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "VerifyMailCode")
	defer span.End()
	//
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		g.Log().Warning(ctx, "VerifyMailCode", req, err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	///
	err = service.TFA().VerifyCode(ctx, userInfo.UserId, req.RiskSerial, req.Code)
	if err != nil {
		g.Log().Warning(ctx, "VerifyMailCode", req, err)
		return nil, err
	}
	return nil, err
}
