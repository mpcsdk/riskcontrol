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

	key := s.verifyPenddingKey(userId, riskSerial, code)
	if task, ok := s.verifyPendding[key]; ok {
		if task != nil {
			task()
		}
		delete(s.verifyPendding, key)
		return nil
	}
	g.Log().Warning(ctx, "VerifyCode:", userId, riskSerial, code)
	return gerror.NewCode(consts.CodeRiskVerifyInvalid)
}
