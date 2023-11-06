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
	if err := c.counter(ctx, req.Token, "UpMail"); err != nil {
		return nil, err
	}
	//
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		g.Log().Warning(ctx, "UpMail:", req, err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	///check phone exists
	err = service.DB().TfaMailNotExists(ctx, req.Mail)
	if err != nil {
		g.Log().Warning(ctx, "UpMail:", req, err)
		return nil, gerror.NewCode(consts.CodeTFAMailExists)
	}
	///
	tfaInfo, err := service.TFA().TFAInfo(ctx, userInfo.UserId)
	if err != nil {
		g.Log().Warning(ctx, "UpMail:", req, err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}

	//upmail riskcontrol
	riskData := &conrisk.RiskTfa{
		UserId: userInfo.UserId,
		Kind:   consts.KEY_TFAKindUpMail,
		Mail:   req.Mail,
	}
	if tfaInfo == nil {
		riskData.Kind = consts.KEY_TFAKindCreate
	}
	riskSerial, code := service.Risk().RiskTFA(ctx, userInfo.UserId, riskData)

	if code == consts.RiskCodeForbidden {
		return nil, gerror.NewCode(consts.CodePerformRiskForbidden)
	}
	if code == consts.RiskCodeError {
		return nil, gerror.NewCode(consts.CodePerformRiskError)
	}
	serial, err := service.TFA().TFAUpMail(ctx, tfaInfo, req.Mail, riskSerial)
	res = &v1.UpMailRes{
		RiskSerial: serial,
	}
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return res, err
	}

	return res, err

	///
	///

}
