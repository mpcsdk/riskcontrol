package tfa

import (
	"context"
	"riskcontral/internal/consts"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sTFA) riskEventPhone(ctx context.Context, phone string, after func()) *riskEvent {
	return &riskEvent{
		Kind:           Key_RiskEventPhone,
		Phone:          phone,
		afterPhoneFunc: after,

		donePhone: false,
		doneMail:  true,
	}
}
func (s *sTFA) riskEventMail(ctx context.Context, mail string, after func()) *riskEvent {
	return &riskEvent{
		Kind:          Key_RiskEventMail,
		Mail:          mail,
		afterMailFunc: after,

		////
		donePhone: true,
		doneMail:  false,
	}
}
func (s *sTFA) riskEventKind(ctx context.Context, event *riskEvent) RiskKind {
	return ""
}
func (s *sTFA) addRiskEvent(ctx context.Context, userId, riskSerial string, event *riskEvent) {

	key := keyUserRiskId(userId, riskSerial)
	if v, ok := s.riskPendding[key]; ok {
		v.riskEvent[event.Kind] = event
	} else {

		risk := &riskPendding{
			UserId:     userId,
			RiskSerial: riskSerial,
			riskEvent: map[RiskKind]*riskEvent{
				event.Kind: event,
			},
		}
		s.riskPendding[key] = risk
	}
}
func (s *sTFA) fetchRiskEvent(ctx context.Context, userId string, riskSerial string, kind RiskKind) *riskEvent {
	key := keyUserRiskId(userId, riskSerial)
	if r, ok := s.riskPendding[key]; ok {
		if e, ok := r.riskEvent[kind]; ok {
			return e
		}
	}
	return nil
}

func (s *sTFA) verifyRiskPendding(ctx context.Context, userId string, riskSerial string, code string, risk *riskPendding) error {

	for kind, event := range risk.riskEvent {
		if kind == Key_RiskEventMail {
			if event.VerifyMailCode == code {
				event.doneMail = true
				return nil
			}
		}
		if kind == Key_RiskEventPhone {
			if event.VerifyPhoneCode == code {
				event.donePhone = true
				return nil
			}
		}
	}

	return gerror.NewCode(consts.CodeRiskVerifyCodeInvalid)
}
func (s *sTFA) doneRiskPendding(ctx context.Context, userId string, riskSerial string, code string, risk *riskPendding) error {
	for kind, event := range risk.riskEvent {
		if kind == Key_RiskEventMail {
			if event.doneMail == false {
				return nil
			}
		}
		if kind == Key_RiskEventPhone {
			if event.donePhone == false {
				return nil
			}
		}
	}
	//done
	for kind, event := range risk.riskEvent {
		if kind == Key_RiskEventMail {
			event.afterMailFunc()
		}
		if kind == Key_RiskEventPhone {
			event.afterMailFunc()
		}
	}
	return nil
}
