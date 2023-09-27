package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

// @Summary 验证token，注册用户tfa
func (c *ControllerV1) UpPhone(ctx context.Context, req *v1.UpPhoneReq) (res *v1.UpPhoneRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "UpPhone")
	defer span.End()
	//
	///check token
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil || userInfo == nil {
		g.Log().Warning(ctx, "UpPhone:", req, err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	///upphoe riskcontrol
	//
	riskData := &conrisk.RiskTfa{
		UserId: userInfo.UserId,
		Kind:   consts.KEY_TFAKindUpPhone,
		Phone:  req.Phone,
	}
	riskSerial, code, err := service.Risk().PerformRiskTFA(ctx, userInfo.UserId, riskData)
	if err != nil {
		return nil, err
	}
	if code != 0 {
		return nil, gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	if code == 0 {

	}
	///
	serial, err := service.TFA().TFAUpPhone(ctx, userInfo.UserId, req.Phone, riskSerial)
	if serial == "" {
		g.Log().Warning(ctx, "UpPhone:", req, err)
		return nil, err
	}
	res = &v1.UpPhoneRes{
		RiskSerial: serial,
	}
	return res, err
}
