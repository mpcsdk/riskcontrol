package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

// @Summary 验证token，注册用户tfa
func (c *ControllerV1) VerifySmsCode(ctx context.Context, req *v1.VerifySmsCodeReq) (res *v1.VerifySmsCodeRes, err error) {
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeAuthFailed)
	}
	///
	return nil, service.TFA().VerifyCode(ctx, userInfo.UserId, "upPhone", req.Code)
}
