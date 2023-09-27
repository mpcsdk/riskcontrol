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
func (c *ControllerV1) UpMail(ctx context.Context, req *v1.UpMailReq) (res *v1.UpMailRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "UpMail")
	defer span.End()
	//
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		g.Log().Warning(ctx, "UpMail:", req, err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	///upphoe riskcontrol
	//
	riskData := &conrisk.RiskTfa{
		UserId: userInfo.UserId,
		Kind:   consts.KEY_TFAKindUpMail,
		Mail:   req.Mail,
	}
	riskSerial, code := service.Risk().PerformRiskTFA(ctx, userInfo.UserId, riskData)
	if code == consts.RiskCodeError {
		return nil, gerror.NewCode(consts.CodePerformRiskFailed)
	}
	if code == consts.RiskCodeNeedVerification {
		// return nil, gerror.NewCode(consts.CodePerformRiskNeedVerification)
	}
	if code == consts.RiskCodePass {

	}
	///
	///
	serial, err := service.TFA().TFAUpMail(ctx, userInfo.UserId, req.Mail, riskSerial)
	if serial == "" {
		g.Log().Warning(ctx, "UpMail:", req, err)
		return nil, err
	}
	res = &v1.UpMailRes{
		RiskSerial: serial,
	}
	return res, err
}
