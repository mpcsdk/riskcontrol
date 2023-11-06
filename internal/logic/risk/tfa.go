package risk

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sRisk) isBefor(upAt *gtime.Time, befor *gtime.Time) bool {
	if upAt.Before(befor) {
		return true
	}
	return false
}
func (s *sRisk) checkTFAUpPhone(ctx context.Context, userId string) (int32, error) {
	/////
	info, err := service.TFA().TFAInfo(ctx, userId)
	if err != nil || info == nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userid", userId),
		))
		return consts.RiskCodeError, err
	}
	// if info == nil {
	// 	g.Log().Warning(ctx, "checkTFAUpPhone userinfo not exists:", userId)
	// 	return consts.RiskCodeNeedVerification, nil
	// }
	if info.PhoneUpdatedAt == nil {
		g.Log().Notice(ctx, "checkTFAUpPhone notuptime :",
			"userId:", userId,
			"info:", info,
			"info.PhoneUpdatedAt:", info.PhoneUpdatedAt)
		return consts.RiskCodeNeedVerification, nil
	}

	befor24h := gtime.Now().Add(BeforH24)
	if info.PhoneUpdatedAt.Before(befor24h) {
		g.Log().Notice(ctx, "checkTFAUpPhone check phoneUpTime:",
			"userId:", userId,
			"info:", info,
			"info.PhoneUpdatedAt:", info.PhoneUpdatedAt,
			"beforAt:", befor24h,
		)
		return consts.RiskCodeNeedVerification, nil
	}
	return consts.RiskCodeForbidden, nil
}

func (s *sRisk) checkTfaUpMail(ctx context.Context, userId string) (int32, error) {
	/////
	info, err := service.TFA().TFAInfo(ctx, userId)
	if err != nil || info == nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userid", userId),
		))
		return consts.RiskCodeError, err
	}
	if info == nil {
		g.Log().Warning(ctx, "checkTfaUpMail userinfo not exists:", userId)
		return consts.RiskCodeNeedVerification, nil
	}
	if info.MailUpdatedAt == nil {
		g.Log().Notice(ctx, "checkTfaUpMail notuptime :",
			"userId:", userId,
			"info:", info,
			"info.PhoneUpdatedAt:", info.PhoneUpdatedAt)
		return consts.RiskCodeNeedVerification, nil
	}
	befor24h := gtime.Now().Add(BeforH24)

	if info.MailUpdatedAt.Before(befor24h) {
		g.Log().Notice(ctx, "checkTfaUpMail notuptime :",
			"userId:", userId,
			"info:", info,
			"info.PhoneUpdatedAt:", info.PhoneUpdatedAt)
		return consts.RiskCodeNeedVerification, nil
	}
	g.Log().Notice(ctx, "checkTfaUpMail check phoneUpTime:",
		"userId:", userId,
		"info:", info,
		"info.PhoneUpdatedAt:", info.PhoneUpdatedAt,
		"beforAt:", befor24h,
	)
	return consts.RiskCodeForbidden, nil
}

func (s *sRisk) checkTfaCreate(ctx context.Context, userId string) (int32, error) {
	return consts.RiskCodePass, nil
}
