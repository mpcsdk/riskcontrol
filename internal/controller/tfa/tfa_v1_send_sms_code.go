package tfa

import (
	"context"
	v1 "riskcontrol/api/tfa/v1"
	"riskcontrol/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error) {
	g.Log().Notice(ctx, "SendSmsCode:", "req:", req)
	//limit
	if err := c.limiter.ApiLimit(ctx, req.Token, "SendSmsCode"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendSmsCode")
	defer span.End()
	//
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil || info == nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	////
	////tidy mail cfg
	service.TFA().TfaTidyPhone(ctx, info.UserId, req.RiskSerial, req.Phone)
	////
	////
	err = service.TFA().SendPhoneCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
