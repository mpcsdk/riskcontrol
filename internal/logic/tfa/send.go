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

	_, err := s.TFAInfo(ctx, userId)
	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", userId, riskSerial, err)
		return "", gerror.NewCode(consts.CodeTFANotExist)
	}
	///
	event := s.fetchRiskEvent(ctx, userId, riskSerial, Key_RiskEventPhone)
	if event == nil {
		//todo
		return "", gerror.NewCode(consts.CodeRiskPerformFailed)
	}

	code, err := service.SmsCode().SendCode(ctx, event.Phone)
	g.Log().Debug(ctx, "SendPhoneCode:", userId, riskSerial, code, event, err)
	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", userId, riskSerial, event, err, code)
		return "", gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	///wait verification
	key := s.verifyPenddingKey(userId, riskSerial, code)
	s.verifyPendding[key] = event.afterPhoneFunc
	return "", nil

}

func (s *sTFA) SendMailCode(ctx context.Context, userId string, riskSerial string) (string, error) {
	_, err := s.TFAInfo(ctx, userId)
	if err != nil {
		g.Log().Warning(ctx, "SendMailCode:", userId, riskSerial, err)
		return "", gerror.NewCode(consts.CodeTFANotExist)
	}
	event := s.fetchRiskEvent(ctx, userId, riskSerial, Key_RiskEventMail)
	if event == nil {
		//todo
		return "", gerror.NewCode(consts.CodeRiskPerformFailed)
	}

	code, err := service.MailCode().SendMailCode(ctx, event.Mail)
	g.Log().Debug(ctx, "SendMailCode:", userId, riskSerial, code, event, err)
	if err != nil {
		g.Log().Warning(ctx, "SendMailCode:", userId, riskSerial, event, err, code)
		return "", gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	///wait verification
	key := s.verifyPenddingKey(userId, riskSerial, code)
	s.verifyPendding[key] = event.afterMailFunc
	return "", nil
}