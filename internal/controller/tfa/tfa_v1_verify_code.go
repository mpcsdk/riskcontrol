package tfa

import (
	"context"
	v1 "riskcontrol/api/tfa/v1"
	"riskcontrol/internal/model"
	"riskcontrol/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) VerifyCode(ctx context.Context, req *v1.VerifyCodeReq) (res *v1.VerifyCodeRes, err error) {
	g.Log().Notice(ctx, "VerifyCode:", "req:", req)
	if err := c.limiter.ApiLimit(ctx, req.Token, "VerifyCode"); err != nil {
		return nil, err
	}
	//trace
	ctx, span := gtrace.NewSpan(ctx, "VerifyCode")
	defer span.End()
	//
	// ///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, mpccode.CodeTFANotExist()
	}
	///
	code := &model.VerifyCode{
		PhoneCode: req.PhoneCode,
		MailCode:  req.MailCode,
	}
	///
	err = service.TFA().VerifyCode(ctx, userInfo.UserId, req.RiskSerial, code)

	if err != nil {
		return nil, err
	}
	return nil, err
}
