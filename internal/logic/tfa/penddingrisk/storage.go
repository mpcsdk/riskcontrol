package pendding

import (
	"context"
	"riskcontrol/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/do"
)

func recordPerson(ctx context.Context, userId string, enable bool) error {
	err := service.DB().TfaDB().UpdateTfaInfo(ctx, userId, &do.Tfa{
		UserId:       userId,
		TxNeedVerify: enable,
	})
	if err != nil {
		g.Log().Warning(ctx, "recordPerson:", "userId:", userId, "enable:", enable, "err:", err)
		return mpccode.CodeInternalError()
	}
	return err
}

func recordPhone(ctx context.Context, userId, phone string, phoneExists bool) error {
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
func recordMail(ctx context.Context, userId, mail string, upMail bool) error {

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
