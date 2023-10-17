package risk

// import (
// 	"context"
// 	"riskcontral/internal/consts"
// 	"riskcontral/internal/service"

// 	"github.com/gogf/gf/v2/errors/gerror"
// 	"github.com/gogf/gf/v2/frame/g"
// )

// func (s *sRisk) RiskPhoneCodeTmp(ctx context.Context, riskserial, phone string) error {

// 	g.Log().Debug(ctx, "RiskPhonecodeTmp:", riskserial, phone)
// 	// userId, err := service.Cache().Get(ctx, riskserial+consts.KEY_RiskUId)
// 	// if err != nil {
// 	// 	g.Log().Warning(ctx, "RiskPhoneCode:", riskserial, err)
// 	// 	return err
// 	// }
// 	// info, err := service.TFA().TFAInfo(ctx, userId.String())
// 	// if err != nil {
// 	// 	g.Log().Warning(ctx, "RiskPhoneCode:", riskserial, info, err)
// 	// 	return err
// 	// }

// 	code, err := service.SmsCode().SendCode(ctx, phone)
// 	if err != nil {
// 		g.Log().Warning(ctx, "RiskPhoneCode:", riskserial, phone, err)
// 		return gerror.NewCode(consts.CodeRiskPhoneInvalid)
// 	}
// 	///recode code
// 	service.Cache().Set(ctx, riskserial+consts.KEY_RiskCode, code, 0)
// 	g.Log().Debug(ctx, "RiskPhoneCode:", riskserial, code)
// 	return err
// }

// func (s *sRisk) RiskMailCodeTmp(ctx context.Context, riskserial, mail string) error {

// 	g.Log().Debug(ctx, "RiskPhonecodeTmp:", riskserial, mail)
// 	// userId, err := service.Cache().Get(ctx, riskserial+consts.KEY_RiskUId)
// 	// if err != nil {
// 	// 	g.Log().Warning(ctx, "RiskMailCode:", riskserial, err)
// 	// 	return err
// 	// }
// 	// info, err := service.TFA().TFAInfo(ctx, userId.String())
// 	// if err != nil {
// 	// 	g.Log().Warning(ctx, "RiskMailCode:", riskserial, info, err)
// 	// 	return err
// 	// }
// 	code, err := service.MailCode().SendMailCode(ctx, mail)
// 	if err != nil {
// 		g.Log().Warning(ctx, "RiskMailCode:", riskserial, err)
// 		return gerror.NewCode(consts.CodeRiskMailInvalid)
// 	}
// 	///recode code
// 	service.Cache().Set(ctx, riskserial+consts.KEY_RiskCode, code, 0)
// 	g.Log().Debug(ctx, "RiskMailCode:", riskserial, code)
// 	return err
// }
