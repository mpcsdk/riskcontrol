package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {
	err = service.Risk().RiskMailCode(ctx, req.RiskSerial)
	return nil, err
}

func (c *ControllerV1) VerifyMailCode(ctx context.Context, req *v1.VerifyMailCodeReq) (res *v1.VerifyMailCodeRes, err error) {
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
