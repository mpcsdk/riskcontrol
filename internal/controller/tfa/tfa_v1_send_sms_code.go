package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"
)

func (c *ControllerV1) SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendSmsCode")
	defer span.End()
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	// err = service.Risk().RiskPhoneCode(ctx, req.RiskSerial)
	_, err = service.TFA().SendPhoneCode(ctx, info.UserId, req.RiskSerial)
	return nil, err
}

// @Summary 验证token，注册用户tfa
func (c *ControllerV1) VerifySmsCode(ctx context.Context, req *v1.VerifySmsCodeReq) (res *v1.VerifySmsCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "VerifySmsCode")
	defer span.End()
	//
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	///
	return nil, service.TFA().VerifyCode(ctx, userInfo.UserId, req.RiskSerial, req.Code)
}
