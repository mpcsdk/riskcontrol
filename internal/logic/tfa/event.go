package tfa

import (
	"context"
	"riskcontral/internal/consts"

	"github.com/gogf/gf/v2/errors/gerror"
)

type riskEvent struct {
	Kind      RiskKind
	DoneEvent bool

	Phone           string
	VerifyPhoneCode string
	afterPhoneFunc  func(context.Context) error

	Mail           string
	VerifyMailCode string
	afterMailFunc  func(context.Context) error
}

func (s *riskEvent) afterFunc() func(context.Context) error {
	if s.Kind == Key_RiskEventMail {
		return s.afterMailFunc
	}
	if s.Kind == Key_RiskEventPhone {
		return s.afterPhoneFunc
	}
	return nil
}

// /
func (s *riskEvent) riskKind() RiskKind {
	return s.Kind
}
func (s *riskEvent) isDone() bool {
	return s.DoneEvent
}
func (s *riskEvent) verify(code string) (bool, RiskKind) {
	if s.Kind == Key_RiskEventPhone {
		if code == s.VerifyPhoneCode {
			s.DoneEvent = true
		}
	}
	///
	if s.Kind == Key_RiskEventMail {
		if code == s.VerifyMailCode {
			s.DoneEvent = true
		}
	}
	return s.DoneEvent, s.Kind
}
func (s *riskEvent) upCode(code string) {
	if s.Kind == Key_RiskEventMail {
		s.VerifyMailCode = code
		s.DoneEvent = false
	}
	if s.Kind == Key_RiskEventPhone {
		s.VerifyPhoneCode = code
		s.DoneEvent = false
	}
}

// /
func newRiskEventPhone(phone string, after func(context.Context) error) *riskEvent {
	return &riskEvent{
		Kind:           Key_RiskEventPhone,
		Phone:          phone,
		afterPhoneFunc: after,

		DoneEvent: false,
	}
}
func newRiskEventMail(mail string, after func(context.Context) error) *riskEvent {
	return &riskEvent{
		Kind:          Key_RiskEventMail,
		Mail:          mail,
		afterMailFunc: after,

		////
		DoneEvent: false,
	}
}

// //
// /
// /
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

func (s *sTFA) verifyRiskPendding(ctx context.Context, userId string, riskSerial string, code string, risk *riskPendding) (RiskKind, error) {
	for _, event := range risk.riskEvent {
		if event.isDone() {
			continue
		}
		if ok, k := event.verify(code); ok {
			return k, nil
		} else {
			if event.Kind == Key_RiskEventMail {
				return k, gerror.NewCode(consts.CodeRiskVerifyMailInvalid)
			}
			if event.Kind == Key_RiskEventPhone {
				return k, gerror.NewCode(consts.CodeRiskVerifyPhoneInvalid)
			}
		}
	}
	return "", gerror.NewCode(consts.CodeRiskVerifyCodeInvalid)
}
