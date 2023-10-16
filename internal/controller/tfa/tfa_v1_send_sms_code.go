package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/net/gtrace"
)

func (c *ControllerV1) SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendSmsCode")
	defer span.End()
	if err := c.counter(ctx, req.Token, "SendSmsCode"); err != nil {
		return nil, err
	}
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	// err = service.Risk().RiskPhoneCode(ctx, req.RiskSerial)
	_, err = service.TFA().SendPhoneCode(ctx, info.UserId, req.RiskSerial)
	return nil, err
}
