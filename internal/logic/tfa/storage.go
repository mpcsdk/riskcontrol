package tfa

import (
	"context"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sTFA) createTFA(ctx context.Context, userId, mail, phone string) error {

	// _, err := dao.Tfa.Ctx(s.ctx).
	// 	Data(do.Tfa{
	// 		Phone:          phone,
	// 		PhoneUpdatedAt: gtime.Now(),
	// 	}).
	// 	Where(dao.Tfa.Columns().UserId, userId).
	// 	Update()

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
	err := service.DB().InsertTfaInfo(ctx, &e)

	return err
}

func (s *sTFA) recordPhone(ctx context.Context, userId, phone string) error {
	err := service.DB().UpdateTfaInfo(ctx, &entity.Tfa{
		UserId:         userId,
		Phone:          phone,
		PhoneUpdatedAt: gtime.Now(),
	})

	return err
}
func (s *sTFA) recordMail(ctx context.Context, userId, mail string) error {

	err := service.DB().UpdateTfaInfo(ctx, &entity.Tfa{
		UserId:        userId,
		Mail:          mail,
		MailUpdatedAt: gtime.Now(),
	})

	return err
}

// //
func (s *sTFA) insertPhone(ctx context.Context, userId, phone string) error {
	err := service.DB().InsertTfaInfo(ctx, &entity.Tfa{
		UserId:         userId,
		Phone:          phone,
		PhoneUpdatedAt: gtime.Now(),
	})

	return err
}
func (s *sTFA) insertMail(ctx context.Context, userId, mail string) error {

	err := service.DB().InsertTfaInfo(ctx, &entity.Tfa{

		UserId:        userId,
		Mail:          mail,
		MailUpdatedAt: gtime.Now(),
	})
	return err
}
