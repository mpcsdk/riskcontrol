package tfa

import (
	"context"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) upPhone(userId, phone string) error {
	_, err := dao.Tfa.Ctx(s.ctx).Where(dao.Tfa.Columns().UserId, userId).Update(entity.Tfa{Phone: phone})
	return err
}
func (s *sTFA) UpPhone(ctx context.Context, userId string, phone string) error {
	//todo:
	riskData := &conrisk.RiskTfa{
		UserId: userId,
		Kind:   "upPhone",
		Phone:  "phone",
	}
	rst, err := service.Risk().PerformRiskTFA(ctx, userId, riskData)
	//todo:
	if err != nil {
		code, err := s.SendMailOTP(ctx, userId, "upMail")
		if err != nil {
			return err
		}
		s.pendding[userId+"upPhone"+code] = func() {
			s.UpPhone(ctx, userId, phone)
		}
		return err
	}
	g.Log().Info(ctx, "UpPhone risk:", rst)
	return err
}
func (s *sTFA) upMail(userId, mail string) error {
	_, err := dao.Tfa.Ctx(s.ctx).Where(dao.Tfa.Columns().UserId, userId).Update(entity.Tfa{Mail: mail})
	return err
}
func (s *sTFA) UpMail(ctx context.Context, userId string, mail string) error {
	//todo:
	riskData := &conrisk.RiskTfa{
		UserId: userId,
		Kind:   "upMail",
		Mail:   mail,
	}
	rst, err := service.Risk().PerformRiskTFA(ctx, userId, riskData)
	if err != nil {
		_, err := s.SendMailOTP(ctx, userId, "upMail")
		//todo:
		// if err != nil {
		// 	return err
		// }
		//todo:
		// if err == nil
		s.pendding[userId+"upMail"+"123456"] = func() {
			s.upMail(userId, mail)
		}
		return err
	}
	g.Log().Info(ctx, "UpPhone risk:", rst)
	return err
}
