package tfa

import (
	"context"
	"riskcontral/internal/config"
	"riskcontral/internal/consts"
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
	// riskPendding          map[UserRiskId]*riskPendding
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
	// t := gcfg.Instance().Get(ctx, "userRisk.verificationDuration", 600)
	t := config.Config.UserRisk.VerificationCodeDuration
	s := &sTFA{
		// verifyPendding: map[string]func(){},
		// mailVerifyPendding:  map[string]func(){},
		// phoneVerifyPendding: map[string]func(){},
		// riskPendding: map[UserRiskId]*riskPendding{},
		//todo:
		riskPenddingContainer: newRiskPenddingContainer(t),
		ctx:                   ctx,
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

func (s *sTFA) TFACreate(ctx context.Context, userId string, phone string, mail string, riskSerial string) ([]string, error) {

	// if err != nil || code != 0 {
	kind := []string{}
	/// need verification
	if phone != "" {
		event := newRiskEventPhone(phone, func(ctx context.Context) error {

			return s.recordPhone(ctx, userId, phone, false)
		})
		s.riskPenddingContainer.Add(userId, riskSerial, event)
		// s.addRiskEvent(ctx, userId, riskSerial, event)
		kind = append(kind, "phone")
		g.Log().Debug(ctx, "TFACreate:", userId, riskSerial, event)
	}

	/// need verification
	if mail != "" {
		event := newRiskEventMail(mail, func(ctx context.Context) error {
			return s.recordMail(ctx, userId, mail, false)
		})
		// s.addRiskEvent(ctx, userId, riskSerial, event)
		s.riskPenddingContainer.Add(userId, riskSerial, event)
		kind = append(kind, "mail")
		g.Log().Debug(ctx, "TFACreate:", userId, riskSerial, event)
	}
	///
	s.riskPenddingContainer.AddBeforFunc(userId, riskSerial, func(ctx context.Context) error {
		return s.createTFA(ctx, userId, mail, phone)
	})
	///
	return kind, nil
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
		// s.addRiskEvent(ctx, userId, riskSerial, event)
		s.riskPenddingContainer.Add(userId, riskSerial, event)
		kind = append(kind, "phone")
	}

	if info.Mail != "" {
		event := newRiskEventMail(info.Mail, nil)
		// s.addRiskEvent(ctx, userId, riskSerial, event)
		s.riskPenddingContainer.Add(userId, riskSerial, event)
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
