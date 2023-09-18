package tfa

import (
	"context"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sTFA) recordPhone(ctx context.Context, userId, phone string) error {
	_, err := dao.Tfa.Ctx(s.ctx).
		Data(do.Tfa{
			Phone:          phone,
			PhoneUpdatedAt: gtime.Now(),
		}).
		Where(dao.Tfa.Columns().UserId, userId).
		Update()
	return err
}
func (s *sTFA) recordMail(ctx context.Context, userId, mail string) error {
	_, err := dao.Tfa.Ctx(s.ctx).
		Data(do.Tfa{
			Mail:          mail,
			MailUpdatedAt: gtime.Now(),
		}).
		Where(dao.Tfa.Columns().UserId, userId).
		Update()
	return err
}

func (s *sTFA) UpPhone(ctx context.Context, userId string, phone string) (string, error) {
	/////todo:
	err := s.hasTFA(ctx, userId)
	if err != nil {

	}
	//todo:
	riskData := &conrisk.RiskTfa{
		UserId: userId,
		Kind:   "upPhone",
		Phone:  phone,
	}
	riskSerial, code, err := service.Risk().PerformRiskTFA(ctx, userId, riskData)
	if err != nil || code != 0 {
		s.pendding[userId+riskSerial] = func() {
			s.recordPhone(ctx, userId, phone)
		}
		return riskSerial, err
	} else {
		s.recordPhone(ctx, userId, phone)
	}
	//
	g.Log().Info(ctx, "UpPhone risk:", riskSerial, code)
	return riskSerial, err
}

func (s *sTFA) UpMail(ctx context.Context, userId string, mail string) (string, error) {
	/////todo:
	err := s.hasTFA(ctx, userId)
	if err != nil {

	}

	//todo:
	riskData := &conrisk.RiskTfa{
		UserId: userId,
		Kind:   "upMail",
		Mail:   mail,
	}
	riskSerial, code, err := service.Risk().PerformRiskTFA(ctx, userId, riskData)
	if err != nil || code != 0 {
		s.pendding[userId+riskSerial] = func() {
			s.recordMail(ctx, userId, mail)
		}
		return riskSerial, err
	} else {
		s.recordMail(ctx, userId, mail)
	}
	//

	g.Log().Info(ctx, "UpPhone risk:", riskSerial, code)
	return riskSerial, err
}
