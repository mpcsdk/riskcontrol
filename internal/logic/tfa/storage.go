package tfa

import (
	"context"
	"riskcontral/internal/model/do"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sTFA) createTFA(ctx context.Context, userId string, mail, phone string) error {

	e := do.Tfa{
		UserId:    userId,
		CreatedAt: gtime.Now(),
	}
	// if mail != "" {
	// 	e.Mail = mail
	// 	// e.MailUpdatedAt = gtime.Now()
	// }
	// if phone != "" {
	// 	e.Phone = phone
	// 	// e.PhoneUpdatedAt = gtime.Now()
	// }
	err := service.DB().InsertTfaInfo(ctx, userId, &e)

	return err
}

func (s *sTFA) recordPhone(ctx context.Context, userId, phone string, phoneExists bool) error {
	if !phoneExists {
		err := service.DB().UpdateTfaInfo(ctx, userId, &do.Tfa{
			UserId: userId,
			Phone:  phone,
		})
		return err
	} else {
		err := service.DB().UpdateTfaInfo(ctx, userId, &do.Tfa{
			UserId:         userId,
			Phone:          phone,
			PhoneUpdatedAt: gtime.Now(),
		})
		return err
	}

}
func (s *sTFA) recordMail(ctx context.Context, userId, mail string, upMail bool) error {

	if !upMail {
		err := service.DB().UpdateTfaInfo(ctx, userId, &do.Tfa{
			UserId: userId,
			Mail:   mail,
		})

		return err
	} else {
		err := service.DB().UpdateTfaInfo(ctx, userId, &do.Tfa{
			UserId:        userId,
			Mail:          mail,
			MailUpdatedAt: gtime.Now(),
		})

		return err
	}
}

// //
// func (s *sTFA) insertPhone(ctx context.Context, userId string, phone *string) error {
// 	err := service.DB().InsertTfaInfo(ctx, userId, &do.Tfa{
// 		UserId:         userId,
// 		Phone:          phone,
// 		PhoneUpdatedAt: gtime.Now(),
// 	})

// 	return err
// }
// func (s *sTFA) insertMail(ctx context.Context, userId string, mail *string) error {

// 	err := service.DB().InsertTfaInfo(ctx, userId, &do.Tfa{

// 		UserId:        userId,
// 		Mail:          mail,
// 		MailUpdatedAt: gtime.Now(),
// 	})
// 	return err
// }
