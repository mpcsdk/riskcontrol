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

type sTFA struct {
	ctx context.Context
	// riskClient riskv1.UserClient
	verifyPendding map[string]func()
	// mailVerifyPendding  map[string]func()
	// phoneVerifyPendding map[string]func()
	///
	riskPendding map[string]*riskPendding
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
		verifyPendding: map[string]func(){},
		// mailVerifyPendding:  map[string]func(){},
		// phoneVerifyPendding: map[string]func(){},
		riskPendding: map[string]*riskPendding{},
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
	riskEvent map[string]*riskEvent
}

const (
	Key_RiskEventPhone = "Key_RiskEventPhone"
	Key_RiskEventMail  = "Key_RiskEventMail"
)

type riskEvent struct {
	Kind           string
	Phone          string
	afterPhoneFunc func()
	Mail           string
	afterMailFunc  func()
}

func init() {
	service.RegisterTFA(new())
}

func (s *sTFA) CreateTFA(ctx context.Context, userId string, phone string, mail string) (string, error) {
	// create nft
	riskData := &conrisk.RiskTfa{
		UserId: userId,
		Kind:   consts.KEY_TFAKindCreate,
		Phone:  phone,
	}
	riskSerial, code, err := service.Risk().PerformRiskTFA(ctx, userId, riskData)
	if err != nil || code != 0 {
		/// need verification
		event := s.riskEventPhone(ctx, phone, func() {
			s.recordPhone(ctx, userId, phone)
		})
		s.addRiskEvent(ctx, userId, riskSerial, event)

		/// need verification
		event = s.riskEventMail(ctx, mail, func() {
			s.recordPhone(ctx, userId, mail)
		})
		s.addRiskEvent(ctx, userId, riskSerial, event)

		return riskSerial, err
	} else {
		s.recordPhone(ctx, userId, phone)
	}
	g.Log().Info(ctx, "CreateTFA PerformRiskTFA:", riskSerial, code)
	return riskSerial, err
}

func (s *sTFA) UpPhone(ctx context.Context, userId string, phone string) (string, error) {
	_, err := s.TFAInfo(ctx, userId)
	if err != nil {
		return s.CreateTFA(ctx, userId, phone, "")
	} else {
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
		event := s.riskEventPhone(ctx, phone, func() {
			s.recordPhone(ctx, userId, phone)
		})
		s.addRiskEvent(ctx, userId, riskSerial, event)
		//
		return riskSerial, gerror.NewCode(consts.CodeRiskNeedVerification)

	}
}

func (s *sTFA) UpMail(ctx context.Context, userId string, mail string) (string, error) {
	_, err := s.TFAInfo(ctx, userId)
	if err != nil {
		return s.CreateTFA(ctx, userId, "", mail)
	} else {
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

		/// need verification
		event := s.riskEventMail(ctx, mail, func() {
			s.recordPhone(ctx, userId, mail)
		})
		s.addRiskEvent(ctx, userId, riskSerial, event)
		//
		return riskSerial, gerror.NewCode(consts.CodeRiskNeedVerification)
		//

	}
}
func (s *sTFA) PerformRiskTFA(ctx context.Context, userId string, riskSerial string) ([]string, error) {
	info, err := s.TFAInfo(ctx, userId)
	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", userId, riskSerial, err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}

	//
	kind := []string{}
	if info.Phone != "" {
		event := s.riskEventPhone(ctx, info.Phone, func() {
			//todo:
		})
		s.addRiskEvent(ctx, userId, riskSerial, event)
		kind = append(kind, "phone")
	}

	if info.Mail != "" {
		event := s.riskEventMail(ctx, info.Mail, func() {
			//todo:
		})
		s.addRiskEvent(ctx, userId, riskSerial, event)
		kind = append(kind, "mail")
	}

	return kind, nil
}
