package risk

import (
	"context"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sRisk) checkTFAUpPhone(ctx context.Context, userId string) (int32, error) {
	/////
	info, err := service.TFA().TFAInfo(ctx, userId)
	if err != nil {
		return -1, err
	}

	befor24h := gtime.Now().Add(BeforH24)
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

	befor24h := gtime.Now().Add(BeforH24)
	if info.MailUpdatedAt.Before(befor24h) {
		return 0, nil
	}
	return 1, nil
}

func (s *sRisk) checkTfaCreate(ctx context.Context, userId string) (int32, error) {
	return 0, nil
}
