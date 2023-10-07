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

type sTFA struct {
	// riskClient riskv1.UserClient
	ctx context.Context
	// verifyPendding map[string]func()
	// mailVerifyPendding  map[string]func()
	// phoneVerifyPendding map[string]func()
	///
	riskPendding          map[UserRiskId]*riskPendding
	riskPenddingContainer *riskPenddingContainer
	// url                   string
	////
}

// func (s *sTFA) setRiskCache(ctx context.Context, key UserRiskId, risk *riskPendding) {
// 	dur, _ := gtime.ParseDuration("1m")
// 	service.Cache().Set(ctx, string(key), risk, dur)
// }

// func (s *sTFA) getRiskCache(ctx context.Context, key UserRiskId) *riskPendding {
// 	val, err := service.Cache().Get(ctx, string(key))
// 	if err != nil {
// 		return nil
// 	}
// 	var risk *riskPendding = nil
// 	val.Struct(*risk)
// 	return risk
// }

// /
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
	riskSerial, code := service.Risk().PerformRiskTFA(ctx, userId, riskData)
	if code == consts.RiskCodeError {
		return "", nil, gerror.NewCode(consts.CodePerformRiskError)
	}
	g.Log().Debug(ctx, "CreateTFA:", userId, phone, mail, code)
	// if err != nil || code != 0 {
	kind := []string{}
	/// need verification
	if phone != "" {
		event := newRiskEventPhone(phone, func(ctx context.Context) error {
			return s.createTFA(ctx, userId, nil, &phone)
		})
		s.addRiskEvent(ctx, userId, riskSerial, event)
		kind = append(kind, "phone")
		g.Log().Debug(ctx, "TFACreate:", userId, riskSerial, event)
	}

	/// need verification
	if mail != "" {
		event := newRiskEventMail(mail, func(ctx context.Context) error {
			return s.createTFA(ctx, userId, &mail, nil)
		})
		s.addRiskEvent(ctx, userId, riskSerial, event)
		kind = append(kind, "mail")
		g.Log().Debug(ctx, "TFACreate:", userId, riskSerial, event)
	}
	return riskSerial, kind, nil
	// }
}

func (s *sTFA) TFAUpPhone(ctx context.Context, userId string, phone string, riskSerial string) (string, error) {
	info, err := s.TFAInfo(ctx, userId)
	if err != nil || info == nil {
		g.Log().Error(ctx, "TFAUpPhone:", userId, phone, err)
		return "", gerror.NewCode(consts.CodeTFANotExist)
	}
	//
	/// need verification
	event := newRiskEventPhone(phone, func(ctx context.Context) error {
		return s.recordPhone(ctx, userId, &phone)
	})
	s.addRiskEvent(ctx, userId, riskSerial, event)
	//
	///tfa mailif
	if info.Mail != "" {
		event := newRiskEventMail(info.Mail, nil)
		s.addRiskEvent(ctx, userId, riskSerial, event)
	}
	///
	return riskSerial, gerror.NewCode(consts.CodePerformRiskNeedVerification)

}

func (s *sTFA) TFAUpMail(ctx context.Context, userId string, mail string, riskSerial string) (string, error) {
	info, err := s.TFAInfo(ctx, userId)
	if err != nil || info == nil {
		g.Log().Error(ctx, "TFAUpMail:", userId, mail, err)
		return "", gerror.NewCode(consts.CodeTFANotExist)
	}
	//
	// modtidy mail
	/// need verification
	event := newRiskEventMail(mail, func(ctx context.Context) error {
		err := s.recordMail(ctx, userId, &mail)
		if err != nil {
			return err
		}
		return service.MailCode().SendBindingMail(ctx, mail)
	})
	s.addRiskEvent(ctx, userId, riskSerial, event)
	///tfa phone if
	if info.Phone != "" {
		event := newRiskEventPhone(info.Phone, nil)
		s.addRiskEvent(ctx, userId, riskSerial, event)
	}

	//
	return riskSerial, gerror.NewCode(consts.CodePerformRiskNeedVerification)
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
