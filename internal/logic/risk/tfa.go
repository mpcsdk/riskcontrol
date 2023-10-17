package risk

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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
	if err != nil {
		g.Log().Warning(ctx, "checkTFAUpPhone:", userId, "err:", err)
		return consts.RiskCodeError, err
	}
	if info == nil {
		g.Log().Debug(ctx, "checkTFAUpPhone:", userId, "info:", info)
		return consts.RiskCodeNeedVerification, nil
	}
	if info.PhoneUpdatedAt == nil {
		g.Log().Debug(ctx, "checkTFAUpPhone:", userId, "info.PhoneUpdatedAt:", info.PhoneUpdatedAt)
		return consts.RiskCodeNeedVerification, nil
	}

	befor24h := gtime.Now().Add(BeforH24)
	g.Log().Debug(ctx, "checkTFAUpPhone:", "befor24h:", befor24h.String(), "info.PhoneUpdatedAt:", info.PhoneUpdatedAt.String())
	if info.PhoneUpdatedAt.Before(befor24h) {
		return consts.RiskCodeNeedVerification, nil
	}
	return consts.RiskCodeForbidden, nil
}

func (s *sRisk) checkTfaUpMail(ctx context.Context, userId string) (int32, error) {
	/////
	info, err := service.TFA().TFAInfo(ctx, userId)
	if err != nil {
		return consts.RiskCodeError, err
	}
	if info == nil {
		return consts.RiskCodeNeedVerification, nil
	}
	if info.MailUpdatedAt == nil {
		return consts.RiskCodeNeedVerification, nil
	}
	befor24h := gtime.Now().Add(BeforH24)
	g.Log().Debug(ctx, "checkTfaUpMail:", "befor24h:", befor24h.String(), "info.PhoneUpdatedAt:", info.PhoneUpdatedAt.String())
	if info.MailUpdatedAt.Before(befor24h) {
		return consts.RiskCodeNeedVerification, nil
	}
	return consts.RiskCodeForbidden, nil
}

func (s *sRisk) checkTfaCreate(ctx context.Context, userId string) (int32, error) {
	return consts.RiskCodePass, nil
}
