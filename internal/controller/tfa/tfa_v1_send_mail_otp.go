package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {
	err = service.Risk().RiskMailCode(ctx, req.RiskSerial)
	return nil, err
}

// @Summary 验证token，注册用户tfa
func (c *ControllerV1) VerifyMailCode(ctx context.Context, req *v1.VerifyMailCodeReq) (res *v1.VerifyMailCodeRes, err error) {
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeAuthFailed)
	}
	///
	return nil, service.TFA().VerifyCode(ctx, userInfo.UserId, req.RiskSerial, req.Code)
}
