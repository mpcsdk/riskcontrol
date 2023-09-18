package risk

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sRisk) RiskPhoneCode(ctx context.Context, riskserial string) error {

	userId, err := service.Cache().Get(ctx, riskserial+consts.KEY_RiskUId)
	if err != nil {
		g.Log().Warning(ctx, "RiskPhoneCode:", riskserial, err)
		return err
	}
	info, err := service.TFA().TFAInfo(ctx, userId.String())
	if err != nil {
		g.Log().Warning(ctx, "RiskPhoneCode:", riskserial, info, err)
		return err
	}

	code, err := service.SmsCode().SendCode(ctx, info.Phone)
	if err != nil {
		g.Log().Warning(ctx, "RiskPhoneCode:", riskserial, info, err)
		return gerror.NewCode(consts.CodeRiskPhoneInvalid)
	}
	///recode code
	service.Cache().Set(ctx, riskserial+consts.KEY_RiskCode, code, 0)
	g.Log().Debug(ctx, "RiskPhoneCode:", riskserial, code)
	return err
}

func (s *sRisk) RiskMailCode(ctx context.Context, riskserial string) error {

	userId, err := service.Cache().Get(ctx, riskserial+consts.KEY_RiskUId)
	if err != nil {
		g.Log().Warning(ctx, "RiskMailCode:", riskserial, err)
		return err
	}
	info, err := service.TFA().TFAInfo(ctx, userId.String())
	if err != nil {
		g.Log().Warning(ctx, "RiskMailCode:", riskserial, info, err)
		return err
	}
	code, err := service.MailCode().SendMailCode(ctx, info.Mail)
	if err != nil {
		g.Log().Warning(ctx, "RiskMailCode:", riskserial, err)
		return gerror.NewCode(consts.CodeRiskMailInvalid)
	}
	///recode code
	service.Cache().Set(ctx, riskserial+consts.KEY_RiskCode, code, 0)
	g.Log().Debug(ctx, "RiskMailCode:", riskserial, code)
	return err
}

func (s *sRisk) VerifyCode(ctx context.Context, serial string, code string) error {
	//verify code
	vcode, err := service.Cache().Get(ctx, serial+consts.KEY_RiskCode)
	if err != nil {
		g.Log().Warning(ctx, "VerifyCode:", serial, err)
		return err
	}
	if vcode == nil {
		return gerror.NewCode(consts.CodeRiskVerifyCodeNotExist)
	}
	if vcode.String() != code {
		g.Log().Warning(ctx, "VerifyCode:", serial, err)
		return gerror.NewCode(consts.CodeRiskVerifyCodeInvalid)
	}
	service.Cache().Remove(ctx, serial+consts.KEY_RiskCode)
	return nil
}
