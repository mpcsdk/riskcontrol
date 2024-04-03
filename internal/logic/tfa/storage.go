package tfa

import (
	"context"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/do"
)

func (s *sTFA) recordPhone(ctx context.Context, userId, phone string, phoneExists bool) error {
	if !phoneExists {
		err := service.DB().TfaDB().UpdateTfaInfo(ctx, userId, &do.Tfa{
			UserId: userId,
			Phone:  phone,
		})
		if err != nil {
			g.Log().Warning(ctx, "recordPhone:", "userId:", userId, "phone:", phone, "err:", err)
			return mpccode.CodeInternalError()
		}
		return nil
	} else {
		err := service.DB().TfaDB().UpdateTfaInfo(ctx, userId, &do.Tfa{
			UserId:         userId,
			Phone:          phone,
			PhoneUpdatedAt: gtime.Now(),
		})
		if err != nil {
			g.Log().Warning(ctx, "recordPhone:", "userId:", userId, "phone:", phone, "err:", err)
			return mpccode.CodeInternalError()
		}
		return err
	}

}
func (s *sTFA) recordMail(ctx context.Context, userId, mail string, upMail bool) error {

	if !upMail {
		err := service.DB().TfaDB().UpdateTfaInfo(ctx, userId, &do.Tfa{
			UserId: userId,
			Mail:   mail,
		})
		if err != nil {
			g.Log().Warning(ctx, "recordPhone:", "userId:", userId, "mail:", mail, "err:", err)
			return mpccode.CodeInternalError()
		}
		return err
	} else {
		err := service.DB().TfaDB().UpdateTfaInfo(ctx, userId, &do.Tfa{
			UserId:        userId,
			Mail:          mail,
			MailUpdatedAt: gtime.Now(),
		})
		if err != nil {
			g.Log().Warning(ctx, "recordPhone:", "userId:", userId, "mail:", mail, "err:", err)
			return mpccode.CodeInternalError()
		}
		return err
	}
}

// //
// func (s *sTFA) insertPhone(ctx context.Context, userId string, phone *string) error {
// 	err := service.DB().TfaDB().InsertTfaInfo(ctx, userId, &do.Tfa{
// 		UserId:         userId,
// 		Phone:          phone,
// 		PhoneUpdatedAt: gtime.Now(),
// 	})

// 	return err
// }
// func (s *sTFA) insertMail(ctx context.Context, userId string, mail *string) error {

// 	err := service.DB().TfaDB().InsertTfaInfo(ctx, userId, &do.Tfa{

// 		UserId:        userId,
// 		Mail:          mail,
// 		MailUpdatedAt: gtime.Now(),
// 	})
// 	return err
// }
