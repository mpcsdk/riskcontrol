package tfa

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) VerifyCode(ctx context.Context, userId string, riskSerial string, code string) error {

	key := keyUserRiskId(userId, riskSerial)
	if risk, ok := s.riskPendding[key]; ok {
		g.Log().Debug(ctx, "VerifyCode:", userId, riskSerial, risk, code)
		_, err := s.verifyRiskPendding(ctx, userId, riskSerial, code, risk)
		if err != nil {
			return err
		}

		if risk.AllDone() {
			risk.DoAfter(ctx, risk)
		}
	}

	return nil
}
