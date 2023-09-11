package tfa

import (
	"context"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) upPhone(token, phone string) error {
	_, err := dao.Tfa.Ctx(s.ctx).Where(dao.Tfa.Columns().Token, token).Update(entity.Tfa{Phone: phone})
	return err
}
func (s *sTFA) UpPhone(ctx context.Context, token string, phone string) error {
	//todo:
	rst, err := service.Risk().PerformRisk(ctx, "phone", nil)
	if err != nil {
		code, err := s.SendMailOTP(ctx, token, "upMail")
		if err != nil {
			return err
		}
		s.pendding[token+"upPhone"+code] = func() {
			s.UpPhone(ctx, token, phone)
		}
		return err
	}
	g.Log().Info(ctx, "UpPhone risk:", rst)
	return err
}
func (s *sTFA) upMail(token, mail string) error {
	_, err := dao.Tfa.Ctx(s.ctx).Where(dao.Tfa.Columns().Token, token).Update(entity.Tfa{Mail: mail})
	return err
}
func (s *sTFA) UpMail(ctx context.Context, token string, mail string) error {
	//todo:
	riskData := &conrisk.TfaUpMail{
		Token: token,
		Mail:  mail,
	}
	rst, err := service.Risk().PerformRisk(ctx, "upMail", riskData)
	if err != nil {
		_, err := s.SendMailOTP(ctx, token, "upMail")
		//todo:
		// if err != nil {
		// 	return err
		// }
		//todo:
		// if err == nil
		s.pendding[token+"upMail"+"123456"] = func() {
			s.upMail(token, mail)
		}
		return err
	}
	g.Log().Info(ctx, "UpPhone risk:", rst)
	return err
}
