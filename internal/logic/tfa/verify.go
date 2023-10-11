package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) VerifyCode(ctx context.Context, userId string, riskSerial string, code *model.VerifyCode) error {

	// key := keyUserRiskId(userId, riskSerial)

	// if risk, ok := s.riskPendding[key]; ok {
	g.Log().Debug(ctx, "VerifyCode:", userId, riskSerial, code)
	k, err := s.riskPenddingContainer.VerifierCode(userId, riskSerial, code)
	// _, err := s.verifyRiskPendding(ctx, userId, riskSerial, code, risk)
	if err != nil {
		g.Log().Warning(ctx, "VerifyCode:", k, err)
		return err
	}
	err = s.riskPenddingContainer.DoAfter(ctx, userId, riskSerial)
	if err != nil {
		g.Log().Warning(ctx, "VerifyCode DoAfter:", err)
		return gerror.NewCode(consts.CodeRiskVerifyCodeInvalid)
	}

	// if risk.AllDone() {
	// 	risk.DoAfter(ctx, risk)
	// 	return nil
	// }
	// }

	return nil
}
