package risk

import (
	"context"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sRisk) checkTFAUpPhone(ctx context.Context, userId string) (int32, error) {
	/////
	info, err := service.TFA().TFAInfo(ctx, userId)
	if err != nil {
		return -1, err
	}
	if info == nil {
		return 0, nil
	}
	if info.PhoneUpdatedAt == nil {
		return 0, nil
	}

	befor24h := gtime.Now().Add(BeforH24)
	g.Log().Debug(ctx, "checkTFAUpPhone:", "befor24h:", befor24h.String(), "info.PhoneUpdatedAt:", info.PhoneUpdatedAt.String())
	if info.PhoneUpdatedAt.Before(befor24h) {
		return 0, nil
	}
	return 1, nil
}

func (s *sRisk) checkTfaUpMail(ctx context.Context, userId string) (int32, error) {

	/////
	info, err := service.TFA().TFAInfo(ctx, userId)
	if err != nil {
		return -1, err
	}
	if info == nil {
		return 0, nil
	}
	if info.MailUpdatedAt == nil {
		return 0, nil
	}
	befor24h := gtime.Now().Add(BeforH24)
	g.Log().Debug(ctx, "checkTFAUpPhone:", "befor24h:", befor24h.String(), "info.PhoneUpdatedAt:", info.PhoneUpdatedAt.String())
	if info.MailUpdatedAt.Before(befor24h) {
		return 0, nil
	}
	return 1, nil
}

func (s *sRisk) checkTfaCreate(ctx context.Context, userId string) (int32, error) {
	return 0, nil
}
