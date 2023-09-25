package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) VerifyCode(ctx context.Context, userId string, vreq []*v1.VerifyReq) error {
	// riskSerial []string, []codes string) error {
	// _, err := s.TFAInfo(ctx, userId)
	// if err != nil {
	// 	g.Log().Warning(ctx, "VerifyCode:", userId, riskSerial, err)
	// 	return gerror.NewCode(consts.CodeTFANotExist)
	// }
	// // 验证验证码

	// key := s.verifyPenddingKey(userId, riskSerial, code)
	// tasks := []func(){}

	for _, v := range vreq {
		key := keyUserRiskId(userId, v.RiskSerial)
		if risk, ok := s.riskPendding[key]; ok {
			g.Log().Debug(ctx, "VerifyCode:", userId, vreq, risk)
			err := s.verifyRiskPendding(ctx, userId, v.RiskSerial, v.Code, risk)
			if err != nil {
				return err
			}
		}
	}

	for _, v := range vreq {
		key := keyUserRiskId(userId, v.RiskSerial)
		if risk, ok := s.riskPendding[key]; ok {
			g.Log().Debug(ctx, "VerifyCode doneRiskPendding:", userId, vreq, risk)
			err := s.doneRiskPendding(ctx, userId, v.RiskSerial, v.Code, risk)
			if err != nil {
				g.Log().Warning(ctx, "VerifyCode:", userId, v.RiskSerial, v.Code, risk, err)
				return err
			}
		}
	}

	return nil
}
