package tfa

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"

	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/model"
	"riskcontral/internal/service"
)

func (c *ControllerV1) VerifyCode(ctx context.Context, req *v1.VerifyCodeReq) (res *v1.VerifyCodeRes, err error) {

	//trace
	ctx, span := gtrace.NewSpan(ctx, "VerifyCode")
	defer span.End()
	if err := c.counter(ctx, req.Token, "VerifyCode"); err != nil {
		return nil, err
	}
	// ///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}

	code := &model.VerifyCode{
		PhoneCode: req.PhoneCode,
		MailCode:  req.MailCode,
	}

	err = service.TFA().VerifyCode(ctx, userInfo.UserId, req.RiskSerial, code)
	if err != nil {
		g.Log().Warning(ctx, "VerifyCode", req, err)
		return nil, gerror.NewCode(mpccode.CodeRiskVerifyCodeInvalid)
	}

	// }
	return nil, err

}
