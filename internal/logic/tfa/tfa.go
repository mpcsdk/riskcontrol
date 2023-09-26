package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type UserRiskId string
type RiskKind string

func keyUserRiskId(userId string, riskSerial string) UserRiskId {
	return UserRiskId(userId + "keyUserRiskId" + riskSerial)
}

type verifier interface {
	exec(risk *riskPendding, code string)
	setNext(verifier)
}
type verifierPhone struct {
	next verifier
}

func (s *verifierPhone) exec(risk *riskPendding, code string) {
	for k, e := range risk.riskEvent {
		if k == Key_RiskEventPhone {
			if e.VerifyPhoneCode == code {
				e.DoneEvent = true
				return
			}
		}
	}
	return
}
func (s *verifierPhone) setNext(v verifier) {
	s.next = v
}

type verifierMail struct {
	next verifier
}

func (s *verifierMail) exec(risk *riskPendding, code string) {
	for k, e := range risk.riskEvent {
		if k == Key_RiskEventMail {
			if e.VerifyMailCode == code {
				e.DoneEvent = true
				return
			}
		}
	}
	return
}
func (s *verifierMail) setNext(v verifier) {
	s.next = v
}

type sTFA struct {
	// riskClient riskv1.UserClient
	ctx context.Context
	// verifyPendding map[string]func()
	// mailVerifyPendding  map[string]func()
	// phoneVerifyPendding map[string]func()
	///
	riskPendding map[UserRiskId]*riskPendding
	url          string
	////
}

func new() *sTFA {

	ctx := gctx.GetInitCtx()
	// addr, err := gcfg.Instance().Get(ctx, "etcd.address")
	// if err != nil {
	// 	panic(err)
	// }
	// grpcx.Resolver.Register(etcd.New(addr.String()))
	// conn, err := grpcx.Client.NewGrpcClientConn("rulerpc")
	// if err != nil {
	// 	panic(err)
	// }
	// client := risk.NewUserClient(conn)
	///
	//
	s := &sTFA{
		// verifyPendding: map[string]func(){},
		// mailVerifyPendding:  map[string]func(){},
		// phoneVerifyPendding: map[string]func(){},
		riskPendding: map[UserRiskId]*riskPendding{},
		ctx:          ctx,
	}
	///

	return s
}

type riskPendding struct {
	//风控序号
	RiskSerial string
	//用户id
	UserId string

	///
	riskEvent map[RiskKind]*riskEvent
}

const (
	Key_RiskEventPhone = "Key_RiskEventPhone"
	Key_RiskEventMail  = "Key_RiskEventMail"
)

///

func init() {
	service.RegisterTFA(new())
}

func (s *sTFA) TFACreate(ctx context.Context, userId string, phone string, mail string) (string, []string, error) {
	// create nft
	riskData := &conrisk.RiskTfa{
		UserId: userId,
		Kind:   consts.KEY_TFAKindCreate,
		Phone:  phone,
		Mail:   mail,
	}
	riskSerial, _, err := service.Risk().PerformRiskTFA(ctx, userId, riskData)
	g.Log().Debug(ctx, "CreateTFA:", userId, phone, mail)
	// if err != nil || code != 0 {
	kind := []string{}
	/// need verification
	if phone != "" {
		event := newRiskEventPhone(phone, func(ctx context.Context) error {
			return s.insertPhone(ctx, userId, phone)
		})
		s.addRiskEvent(ctx, userId, riskSerial, event)
		kind = append(kind, "phone")
		g.Log().Debug(ctx, "TFACreate:", userId, riskSerial, event)
	}

	/// need verification
	if mail != "" {
		event := newRiskEventMail(mail, func(ctx context.Context) error {
			return s.insertMail(ctx, userId, mail)
		})
		s.addRiskEvent(ctx, userId, riskSerial, event)
		kind = append(kind, "mail")
		g.Log().Debug(ctx, "TFACreate:", userId, riskSerial, event)
	}
	return riskSerial, kind, err
	// }
}

func (s *sTFA) TFAUpPhone(ctx context.Context, userId string, phone string) (string, error) {
	info, err := s.TFAInfo(ctx, userId)
	if err != nil {
		g.Log().Error(ctx, "TFAUpPhone:", userId, phone, err)
		return "", gerror.NewCode(consts.CodeTFANotExist)
	}
	//
	//upphone
	//
	riskData := &conrisk.RiskTfa{
		UserId: userId,
		Kind:   consts.KEY_TFAKindUpPhone,
		Phone:  phone,
	}
	riskSerial, _, err := service.Risk().PerformRiskTFA(ctx, userId, riskData)
	if err != nil {
		return "", err
	}

	/// need verification
	event := newRiskEventPhone(phone, func(ctx context.Context) error {
		return s.recordPhone(ctx, userId, phone)
	})
	s.addRiskEvent(ctx, userId, riskSerial, event)
	//
	///tfa mailif
	if info.Mail != "" {
		event := newRiskEventMail(info.Mail, nil)
		s.addRiskEvent(ctx, userId, riskSerial, event)
	}
	///
	return riskSerial, gerror.NewCode(consts.CodeRiskNeedVerification)

}

func (s *sTFA) TFAUpMail(ctx context.Context, userId string, mail string) (string, error) {
	info, err := s.TFAInfo(ctx, userId)
	if err != nil {
		g.Log().Error(ctx, "TFAUpMail:", userId, mail, err)
		return "", gerror.NewCode(consts.CodeTFANotExist)
	}
	//
	riskData := &conrisk.RiskTfa{
		UserId: userId,
		Kind:   consts.KEY_TFAKindUpMail,
		Mail:   mail,
	}

	riskSerial, _, err := service.Risk().PerformRiskTFA(ctx, userId, riskData)
	if err != nil {
		return "", err
	}

	// modtidy mail
	/// need verification
	event := newRiskEventMail(mail, func(ctx context.Context) error {
		return s.recordMail(ctx, userId, mail)
	})
	s.addRiskEvent(ctx, userId, riskSerial, event)
	///tfa phone if
	if info.Phone != "" {
		event := newRiskEventPhone(info.Phone, nil)
		s.addRiskEvent(ctx, userId, riskSerial, event)
	}

	//
	return riskSerial, gerror.NewCode(consts.CodeRiskNeedVerification)
	//

}
func (s *sTFA) TFATx(ctx context.Context, userId string, riskSerial string) ([]string, error) {
	info, err := s.TFAInfo(ctx, userId)
	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", userId, riskSerial, err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}

	//
	kind := []string{}
	if info.Phone != "" {
		event := newRiskEventPhone(info.Phone, nil)
		s.addRiskEvent(ctx, userId, riskSerial, event)
		kind = append(kind, "phone")
	}

	if info.Mail != "" {
		event := newRiskEventMail(info.Mail, nil)
		s.addRiskEvent(ctx, userId, riskSerial, event)
		kind = append(kind, "mail")
	}

	g.Log().Debug(ctx, "PerformRiskTFA:",
		"userId:", userId,
		"riskSerial:", riskSerial,
		"kind:", kind,
		"info:", info,
	)
	return kind, nil
}
