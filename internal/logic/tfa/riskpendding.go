package tfa

import (
	"context"
	"errors"
	"riskcontral/internal/consts"
	"riskcontral/internal/model"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
)

type riskPenddingContainer struct {
	riskPendding map[UserRiskId]*riskPendding
	ctx          context.Context
}

func newRiskPenddingContainer(times int) *riskPenddingContainer {
	s := &riskPenddingContainer{
		riskPendding: make(map[UserRiskId]*riskPendding),
		ctx:          context.Background(),
	}
	//
	gtimer.Add(s.ctx, time.Second*time.Duration(times), func(ctx context.Context) {
		for key, risk := range s.riskPendding {
			if risk.dealline.Before(gtime.Now()) {
				delete(s.riskPendding, key)
			}
		}
	})
	//
	return s
}

func (s *riskPenddingContainer) Get(userId, riskSerial string) *riskPendding {
	key := keyUserRiskId(userId, riskSerial)
	if risk, ok := s.riskPendding[key]; ok {
		return risk
	}
	return nil
}

func (s *riskPenddingContainer) GetEvent(userId, riskSerial string, kind RiskKind) *riskEvent {
	risk := s.Get(userId, riskSerial)
	if risk == nil {
		return nil
	}
	return risk.fetchEvent(kind)
}

func (s *riskPenddingContainer) Add(userId, riskSerial string, event *riskEvent) {
	risk := s.Get(userId, riskSerial)
	if risk == nil {
		risk = &riskPendding{
			UserId:     userId,
			RiskSerial: riskSerial,
			riskEvent: map[RiskKind]*riskEvent{
				event.Kind: event,
			},
			verifier: &emptyVerifier{},
			//todo:
			// deadline: gtime.Now().Add(BeforH24),
			dealline: gtime.Now(),
		}

		key := keyUserRiskId(userId, riskSerial)
		s.riskPendding[key] = risk
	} else {
		risk.riskEvent[event.Kind] = event
	}
	///
	if event.Kind == Key_RiskEventMail {
		risk.verifier.setNext(&verifierMail{})
	}
	if event.Kind == Key_RiskEventPhone {
		risk.verifier.setNext(&verifierPhone{})
	}
}

func (s *riskPenddingContainer) Del(userId, riskSerial string) {
	key := keyUserRiskId(userId, riskSerial)
	delete(s.riskPendding, key)
}
func (s *riskPenddingContainer) UpCode(userId, riskSerial string, kind RiskKind, code string) {
	risk := s.Get(userId, riskSerial)
	if risk == nil {
		return
	}
	risk.upCode(kind, code)
}

var errRiskNotExist error = errors.New("risk not exist")
var errRiskNotDone error = errors.New("risk not done")

func (s *riskPenddingContainer) VerifierCode(userId, riskSerial string, code *model.VerifyCode) (string, error) {
	risk := s.Get(userId, riskSerial)
	if risk == nil {
		return "", errRiskNotExist
	}
	k, err := risk.verifierCode(code)
	return string(k), err
	// risk.verifier()
	// return "", nil
}
func (s *riskPenddingContainer) AllDone(userId, riskSerial string) (string, error) {
	risk := s.Get(userId, riskSerial)
	if risk == nil {
		return "", errRiskNotExist
	}
	if risk.AllDone() {
		return "", nil
	}
	return "", nil
}
func (s *riskPenddingContainer) DoAfter(ctx context.Context, userId, riskSerial string) error {
	risk := s.Get(userId, riskSerial)
	if risk == nil {
		return errRiskNotExist
	}
	if !risk.AllDone() {
		return errRiskNotDone
	}

	risk.DoAfter(ctx, risk)
	return nil
}

// //
// //
type riskPendding struct {
	//风控序号
	RiskSerial string
	//用户id
	UserId string

	///
	riskEvent map[RiskKind]*riskEvent
	verifier  verifier
	dealline  *gtime.Time
}

func (s *riskPendding) verifierCode(code *model.VerifyCode) (RiskKind, error) {
	return s.verifier.exec(s, code)
}

func (s *riskPendding) upCode(kind RiskKind, code string) (string, error) {
	if event, ok := s.riskEvent[kind]; ok {
		event.setCode(code)
	}
	return string(kind), nil
}

func (s *riskPendding) DoAfter(ctx context.Context, risk *riskPendding) (string, error) {
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
	}
	return "", nil
}
func (s *riskPendding) AllDone() bool {
	for _, e := range s.riskEvent {
		if e.isDone() {
			continue
		}
		return false
	}
	return true
}
func (s *riskPendding) fetchEvent(kind RiskKind) *riskEvent {
	if e, ok := s.riskEvent[kind]; ok {
		return e
	}
	return nil
}

// //
// //
type verifier interface {
	exec(risk *riskPendding, verifierCode *model.VerifyCode) (RiskKind, error)
	setNext(verifier)
}
type emptyVerifier struct {
	next verifier
}

func (s *emptyVerifier) exec(risk *riskPendding, verifierCode *model.VerifyCode) (RiskKind, error) {
	if s.next == nil {
		return "", nil
	}
	return s.next.exec(risk, verifierCode)
}
func (s *emptyVerifier) setNext(v verifier) {
	if s.next == nil {
		s.next = v
	} else {
		s.next.setNext(v)
	}
}

// /
type verifierPhone struct {
	next verifier
}

func (s *verifierPhone) exec(risk *riskPendding, verifierCode *model.VerifyCode) (RiskKind, error) {
	for k, e := range risk.riskEvent {
		if k == Key_RiskEventPhone {
			if e.VerifyPhoneCode == verifierCode.PhoneCode && verifierCode.PhoneCode != "" {
				e.DoneEvent = true
				break
			} else {
				return Key_RiskEventPhone, gerror.NewCode(consts.CodeRiskVerifyPhoneInvalid)
			}
		}
	}
	if s.next == nil {
		return "", nil
	}
	return s.next.exec(risk, verifierCode)
}
func (s *verifierPhone) setNext(v verifier) {
	s.next = v
}

type verifierMail struct {
	next verifier
}

func (s *verifierMail) exec(risk *riskPendding, verifierCode *model.VerifyCode) (RiskKind, error) {
	for k, e := range risk.riskEvent {
		if k == Key_RiskEventMail {
			if e.VerifyMailCode == verifierCode.MailCode && verifierCode.MailCode != "" {
				e.DoneEvent = true
				break
			} else {
				return Key_RiskEventMail, gerror.NewCode(consts.CodeRiskVerifyMailInvalid)

			}
		}
	}
	if s.next == nil {
		return "", nil
	}
	return s.next.exec(risk, verifierCode)
}
func (s *verifierMail) setNext(v verifier) {
	s.next = v
}

// //
// //
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

// /
// /
// /
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

// /
func (s *riskEvent) setCode(code string) {
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
// func (s *sTFA) addRiskEvent(ctx context.Context, userId, riskSerial string, event *riskEvent) {
// 	key := keyUserRiskId(userId, riskSerial)
// 	if v, ok := s.riskPendding[key]; ok {
// 		v.riskEvent[event.Kind] = event
// 	} else {
// 		risk := &riskPendding{
// 			UserId:     userId,
// 			RiskSerial: riskSerial,
// 			riskEvent: map[RiskKind]*riskEvent{
// 				event.Kind: event,
// 			},
// 		}
// 		s.riskPendding[key] = risk
// 	}
// }

// func (s *sTFA) fetchRiskEvent(ctx context.Context, userId string, riskSerial string, kind RiskKind) *riskEvent {
// 	key := keyUserRiskId(userId, riskSerial)
// 	if r, ok := s.riskPendding[key]; ok {
// 		if e, ok := r.riskEvent[kind]; ok {
// 			return e
// 		}
// 	}
// 	return nil
// }

func (s *sTFA) verifyRiskPendding(ctx context.Context, userId string, riskSerial string, code *model.VerifyCode, risk *riskPendding) (RiskKind, error) {
	for _, event := range risk.riskEvent {
		// if event.isDone() {
		// 	continue
		// }
		if event.Kind == Key_RiskEventMail {
			ok, k := event.verify(code.MailCode)
			if ok {
				continue

			} else {
				return k, gerror.NewCode(consts.CodeRiskVerifyMailInvalid)
			}
		}
		if event.Kind == Key_RiskEventPhone {
			ok, k := event.verify(code.PhoneCode)
			if ok {
				continue

			} else {
				return k, gerror.NewCode(consts.CodeRiskVerifyPhoneInvalid)
			}
		}
	}
	return "", nil
}
