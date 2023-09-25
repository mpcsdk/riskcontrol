package tfa

import (
	"context"
	"riskcontral/internal/consts"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) riskEventPhone(ctx context.Context, phone string, after func() error) *riskEvent {
	return &riskEvent{
		Kind:           Key_RiskEventPhone,
		Phone:          phone,
		afterPhoneFunc: after,

		DonePhone: false,
		DoneMail:  true,
	}
}
func (s *sTFA) riskEventMail(ctx context.Context, mail string, after func() error) *riskEvent {
	return &riskEvent{
		Kind:          Key_RiskEventMail,
		Mail:          mail,
		afterMailFunc: after,

		////
		DonePhone: true,
		DoneMail:  false,
	}
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

func (s *sTFA) upRiskEventCode(ctx context.Context, event *riskEvent, code string) {
	if event.Kind == Key_RiskEventMail {
		event.VerifyMailCode = code
		event.DoneMail = false
	}
	if event.Kind == Key_RiskEventPhone {
		event.VerifyPhoneCode = code
		event.DonePhone = false
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
			if event.VerifyMailCode == code || event.VerifyMailCode != "" {
				event.DoneMail = true
				return nil
			}
		}
		if kind == Key_RiskEventPhone {
			if event.VerifyPhoneCode == code || event.VerifyPhoneCode != "" {
				event.DonePhone = true
				return nil
			}
		}
	}

	return gerror.NewCode(consts.CodeRiskVerifyCodeInvalid)
}
func (s *sTFA) doneRiskPendding(ctx context.Context, userId string, riskSerial string, code string, risk *riskPendding) error {
	for kind, event := range risk.riskEvent {
		if kind == Key_RiskEventMail {
			if event.DoneMail != true {
				return gerror.NewCode(consts.CodeRiskVerifyMailInvalid)
			}
		}
		if kind == Key_RiskEventPhone {
			if event.DonePhone != true {
				return gerror.NewCode(consts.CodeRiskVerifyPhoneInvalid)
			}
		}
	}
	//done
	for kind, event := range risk.riskEvent {
		if kind == Key_RiskEventMail {
			if event.afterMailFunc != nil {
				err := event.afterMailFunc()
				if err != nil {
					return err
				}
				g.Log().Debug(ctx, "doneRiskPendding:", event)

			}
		}
		if kind == Key_RiskEventPhone {
			if event.afterPhoneFunc != nil {
				err := event.afterPhoneFunc()
				if err != nil {
					return err
				}
				g.Log().Debug(ctx, "doneRiskPendding:", event)
			}
		}
	}
	return nil
}
