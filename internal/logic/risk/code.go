package risk

import (
	"context"
	"riskcontral/internal/service"

	"github.com/gogf/gf/frame/g"
)

func (s *sRisk) RiskPhoneCode(ctx context.Context, riskserial string) error {
	return nil
}
func (s *sRisk) RiskMailCode(ctx context.Context, riskserial string) error {

	userId, err := s.cache.Get(ctx, riskserial+"riskUserId")
	if err != nil {
		return err
	}
	info, err := service.TFA().TFAInfo(ctx, userId.String())
	if err != nil {
		return err
	}
	//todo: senderr
	code, err := service.MailCode().SendMailCode(ctx, info.Mail)
	///recode code
	s.cache.Set(ctx, riskserial+"riskCode", code, 0)
	g.Log().Debug(ctx, "RiskMailCode:", riskserial, code)
	return err
}

func (s *sRisk) VerifyCode(ctx context.Context, serial string, code string) error {
	//verify code
	vcode, err := s.cache.Get(ctx, serial+"riskCode")
	if err != nil {
		return err
	}
	if vcode == nil {
		//todo: checkcode
		// return gerror.NewCode(gcode.CodeInternalError)
	}
	if vcode.String() != code {
		//todo: checkcode
		// return errors.New("VerificationCode failed")
	}
	s.cache.Remove(ctx, serial+"riskCode")
	return nil
}
