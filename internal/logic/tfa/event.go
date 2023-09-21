package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) verifyPenddingKey(userId string, riskSerial string, code string) string {
	return userId + riskSerial + "verifyPenddingKey" + code
}
func (s *sTFA) sendPhoneCode(ctx context.Context, userId, phone, riskSerial string) (string, error) {
	////
	code, err := service.SmsCode().SendCode(ctx, phone)
	g.Log().Debug(ctx, "SendPhoneCode:", userId, phone, riskSerial, code, err)
	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", userId, riskSerial, err, code)
		return "", gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	key := s.verifyPenddingKey(userId, riskSerial, code)
	s.verifyPendding[key] = func() {
		s.recordPhone(ctx, userId, phone)
	}
	return "", nil
}

func (s *sTFA) sendMailOTP(ctx context.Context, userId, mail, riskSerial string) (string, error) {
	////
	code, err := service.MailCode().SendMailCode(ctx, mail)

	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", userId, riskSerial, err, code)
		return "", gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	key := s.verifyPenddingKey(userId, riskSerial, code)
	s.verifyPendding[key] = func() {
		s.recordMail(ctx, userId, mail)
	}
	return "", nil
}
func (s *sTFA) riskEventPhone(ctx context.Context, phone string, after func()) *riskEvent {
	return &riskEvent{
		Kind:           Key_RiskEventPhone,
		Phone:          phone,
		afterPhoneFunc: after,
	}
}
func (s *sTFA) riskEventMail(ctx context.Context, mail string, after func()) *riskEvent {
	return &riskEvent{
		Kind:          Key_RiskEventMail,
		Mail:          mail,
		afterMailFunc: after,
	}
}

func (s *sTFA) addRiskEvent(ctx context.Context, userId, riskSerial string, event *riskEvent) {
	if v, ok := s.riskPendding[riskSerial]; ok {
		v.riskEvent[event.Kind] = event
	} else {
		risk := &riskPendding{
			UserId:     userId,
			RiskSerial: riskSerial,
			riskEvent:  map[string]*riskEvent{event.Kind: event},
		}
		s.riskPendding[riskSerial] = risk
	}
}
func (s *sTFA) fetchRiskEvent(ctx context.Context, userId string, riskSerial string, kind string) *riskEvent {
	if r, ok := s.riskPendding[riskSerial]; ok {
		if e, ok := r.riskEvent[kind]; ok {
			return e
		}
	}
	return nil
}
