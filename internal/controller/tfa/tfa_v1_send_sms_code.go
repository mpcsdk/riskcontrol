package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

func (c *ControllerV1) SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendSmsCode")
	defer span.End()
	//limit
	// if err := c.counter(ctx, req.Token, "SendSmsCode"); err != nil {
	// 	return nil, err
	// }
	if err := c.limitSendVerification(ctx, req.Token, "SendSmsCode"); err != nil {
		// return nil, gerror.NewCode(consts.ErrLimitSendPhoneCode.
		// 	AddDetail("req.token", req.Token).
		// 	AddDetail("req.riskSerial", req.RiskSerial), err.Error())
		err = gerror.Wrap(err, gjson.MustEncodeString(map[string]interface{}{
			"req.token":      req.Token,
			"req.riskSerial": req.RiskSerial,
		}))
		g.Log().Warning(ctx, "SendSmsCode limit", "token:", req.Token)
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(consts.ErrLimitSendPhoneCode)
	}
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	// err = service.Risk().RiskPhoneCode(ctx, req.RiskSerial)
	_, err = service.TFA().SendPhoneCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err

	}
	return nil, err
}
