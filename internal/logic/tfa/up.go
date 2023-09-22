package tfa

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"

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

// //
func (s *sTFA) insertPhone(ctx context.Context, userId, phone string) error {
	_, err := dao.Tfa.Ctx(s.ctx).
		Data(do.Tfa{
			UserId:         userId,
			Phone:          phone,
			PhoneUpdatedAt: gtime.Now(),
		}).
		Where(dao.Tfa.Columns().UserId, userId).
		Insert()
	return err
}
func (s *sTFA) insertMail(ctx context.Context, userId, mail string) error {
	_, err := dao.Tfa.Ctx(s.ctx).
		Data(do.Tfa{
			UserId:        userId,
			Mail:          mail,
			MailUpdatedAt: gtime.Now(),
		}).
		Insert()
	return err
}
