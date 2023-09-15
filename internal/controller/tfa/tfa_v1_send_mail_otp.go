package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) SendMailOTP(ctx context.Context, req *v1.SendMailOTPReq) (res *v1.SendMailOTPRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

// @Summary 验证token，注册用户tfa
func (c *ControllerV1) VerifyMailOTP(ctx context.Context, req *v1.VerifyMailOTPReq) (res *v1.VerifyMailOTPRes, err error) {
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeAuthFailed)
	}
	///
	return nil, service.TFA().VerifyCode(ctx, userInfo.UserId, "upMail", req.Code)
}
