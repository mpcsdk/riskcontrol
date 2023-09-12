package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

// @Summary 验证token，注册用户tfa
func (c *ControllerV1) UpMail(ctx context.Context, req *v1.UpMailReq) (res *v1.UpMailRes, err error) {
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeAuthFailed)
	}
	///

	return nil, service.TFA().UpMail(ctx, userInfo.UserId, req.Mail)
}
