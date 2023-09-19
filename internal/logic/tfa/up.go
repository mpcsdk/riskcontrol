package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sTFA) createTFA(ctx context.Context, userId, mail, phone string) error {
	_, err := dao.Tfa.Ctx(s.ctx).
		Data(do.Tfa{
			Phone:          phone,
			PhoneUpdatedAt: gtime.Now(),
		}).
		Where(dao.Tfa.Columns().UserId, userId).
		Update()

	e := entity.Tfa{
		UserId:    userId,
		CreatedAt: gtime.Now(),
	}
	if mail != "" {
		e.Mail = mail
		e.MailUpdatedAt = gtime.Now()
	}
	if phone != "" {
		e.Phone = phone
		e.PhoneUpdatedAt = gtime.Now()
	}
	_, err = dao.Tfa.Ctx(ctx).Insert(e)

	return err
}

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

func (s *sTFA) CreateTFA(ctx context.Context, userId string, phone string, mail string) (string, error) {
	// create nft
	riskData := &conrisk.RiskTfa{
		UserId: userId,
		Kind:   consts.KEY_TFAKindCreate,
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
	g.Log().Info(ctx, "CreateTFA PerformRiskTFA:", riskSerial, code)
	return riskSerial, err
}

func (s *sTFA) UpPhone(ctx context.Context, userId string, phone string) (string, error) {
	_, err := s.TFAInfo(ctx, userId)
	if err != nil {
		return s.CreateTFA(ctx, userId, phone, "")
	} else {
		//upphone
		//
		riskData := &conrisk.RiskTfa{
			UserId: userId,
			Kind:   consts.KEY_TFAKindUpPhone,
			Phone:  phone,
		}
		riskSerial, code, err := service.Risk().PerformRiskTFA(ctx, userId, riskData)
		if err != nil {
			return "", err
		}

		if code != 0 {
			s.sendpendding[userId+riskSerial] = func() {
				s.sendPhoneCode(ctx, userId, phone, riskSerial)
			}

			return riskSerial, gerror.NewCode(consts.CodeRiskVerification)
		} else {
			s.recordPhone(ctx, userId, phone)
			return "", nil
		}
	}
}

func (s *sTFA) UpMail(ctx context.Context, userId string, mail string) (string, error) {
	_, err := s.TFAInfo(ctx, userId)
	if err != nil {
		return s.CreateTFA(ctx, userId, "", mail)
	} else {
		//
		riskData := &conrisk.RiskTfa{
			UserId: userId,
			Kind:   consts.KEY_TFAKindUpMail,
			Mail:   mail,
		}
		riskSerial, code, err := service.Risk().PerformRiskTFA(ctx, userId, riskData)
		if err != nil {
			return "", err
		}
		if code != 0 {
			s.sendpendding[userId+riskSerial] = func() {
				s.sendMailOTP(ctx, userId, mail, riskSerial)
			}

			return riskSerial, gerror.NewCode(consts.CodeRiskVerifyCodeInvalid)
		} else {
			s.recordMail(ctx, userId, mail)
			return "", nil
		}
		//

		g.Log().Info(ctx, "UpPhone UpMail:", riskSerial, code)
		return riskSerial, err
	}
}
