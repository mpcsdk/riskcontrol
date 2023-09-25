package tfa

import (
	"context"
	"riskcontral/internal/consts"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) VerifyCode(ctx context.Context, userId string, riskSerial string, code string) error {
	// _, err := s.TFAInfo(ctx, userId)
	// if err != nil {
	// 	g.Log().Warning(ctx, "VerifyCode:", userId, riskSerial, err)
	// 	return gerror.NewCode(consts.CodeTFANotExist)
	// }
	// // 验证验证码

	// key := s.verifyPenddingKey(userId, riskSerial, code)
	// tasks := []func(){}
	key := keyUserRiskId(userId, riskSerial)
	if risk, ok := s.riskPendding[key]; ok {
		err := s.verifyRiskPendding(ctx, userId, riskSerial, code, risk)
		if err != nil {
			return err
		}
		err = s.doneRiskPendding(ctx, userId, riskSerial, code, risk)
		g.Log().Debug(ctx, "VerifyCode:", userId, riskSerial, code, risk)
		return err
	}
	return gerror.NewCode(consts.CodeRiskVerifyInvalid)
}
