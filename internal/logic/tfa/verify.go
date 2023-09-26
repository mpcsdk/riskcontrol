package tfa

import (
	"context"
	"riskcontral/internal/consts"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) VerifyCode(ctx context.Context, userId string, riskSerial string, code string) error {
	// riskSerial []string, []codes string) error {
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
		g.Log().Debug(ctx, "VerifyCode:", userId, riskSerial, code)
		_, err := s.verifyRiskPendding(ctx, userId, riskSerial, code, risk)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *sTFA) DoneVerifyCode(ctx context.Context, userId string, riskSerial string) (string, error) {
	key := keyUserRiskId(userId, riskSerial)
	if risk, ok := s.riskPendding[key]; ok {
		for _, event := range risk.riskEvent {
			if !event.isDone() {
				return string(event.Kind), gerror.NewCode(consts.CodeRiskVerifyCodeInvalid)
			}
		}
		//done

		for _, event := range risk.riskEvent {
			if f := event.afterFunc(); f != nil {
				err := f(ctx)
				if err != nil {
					return string(event.Kind), err
				}
			}
			// if kind == Key_RiskEventMail {
			// 	if event.afterMailFunc != nil {
			// 		err := event.afterMailFunc()
			// 		if err != nil {
			// 			return err
			// 		}
			// 		g.Log().Debug(ctx, "doneRiskPendding:", event)

			// 	}
			// }
			// if kind == Key_RiskEventPhone {
			// 	if event.afterPhoneFunc != nil {
			// 		err := event.afterPhoneFunc()
			// 		if err != nil {
			// 			return err
			// 		}
			// 		g.Log().Debug(ctx, "doneRiskPendding:", event)
			// 	}
			// }
		}
	}

	return "", nil
}
