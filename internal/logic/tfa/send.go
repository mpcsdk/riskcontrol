package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// /
func (s *sTFA) SendPhoneCode(ctx context.Context, userId string, riskSerial string) (string, error) {

	event := s.riskPenddingContainer.GetEvent(userId, riskSerial, Key_RiskEventPhone)
	if event == nil {
		g.Log().Warning(ctx, "SendPhoneCode not event:", "userid:", userId, "riskSerail:", riskSerial)
		return "", gerror.NewCode(consts.CodeRiskSerialNotExist)
	}

	code, err := service.SmsCode().SendCode(ctx, event.Phone)
	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", "userid:", userId, "riskSerail:", "sendEvent:", event)
		g.Log().Errorf(ctx, "%+v", err)
		return "", gerror.NewCode(consts.CodeTFASendSmsFailed)
	}

	event.upCode(code)

	return "", nil

}

func (s *sTFA) SendMailCode(ctx context.Context, userId string, riskSerial string) (string, error) {
	// _, err := s.TFAInfo(ctx, userId)
	// if err != nil {
	// 	g.Log().Warning(ctx, "SendMailCode:", userId, riskSerial, err)
	// 	return "", gerror.NewCode(consts.CodeTFANotExist)
	// }

	event := s.riskPenddingContainer.GetEvent(userId, riskSerial, Key_RiskEventMail)
	if event == nil {
		g.Log().Warning(ctx, "SendMailCode:", userId, riskSerial)
		//todo
		return "", gerror.NewCode(consts.CodeRiskSerialNotExist)
	}

	code, err := service.MailCode().SendMailCode(ctx, event.Mail)
	if err != nil {
		g.Log().Warning(ctx, "SendMailCode:", userId, riskSerial, event, err, code)
		return "", gerror.NewCode(consts.CodeTFASendMailFailed)
	}

	event.upCode(code)
	// s.upRiskEventCode(ctx, event, code)
	g.Log().Debug(ctx, "SendMailCode:", userId, riskSerial, code, event, err)
	///wait verification
	// s.verifyRiskPendding(ctx, userId, riskSerial, code, event)
	// key := s.verifyPenddingKey(userId, riskSerial, code)
	// s.verifyPendding[key] = event.afterMailFunc
	return "", nil
}
