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
	ctx                   context.Context
	riskPenddingContainer *riskPenddingContainer
	////
}

// /
func new() *sTFA {

	ctx := gctx.GetInitCtx()
	//
	t := config.Config.UserRisk.VerificationCodeDuration
	s := &sTFA{

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
	}

	/// need verification
	if mail != "" {
		event := newRiskEventMail(mail, func(ctx context.Context) error {
			return s.recordMail(ctx, userId, mail, false)
		})
		// s.addRiskEvent(ctx, userId, riskSerial, event)
		s.riskPenddingContainer.Add(userId, riskSerial, event)
		kind = append(kind, "mail")
	}
	///
	s.riskPenddingContainer.AddBeforFunc(userId, riskSerial, func(ctx context.Context) error {
		return s.createTFA(ctx, userId, mail, phone)
	})
	///
	return kind, nil
}

func (s *sTFA) TFATx(ctx context.Context, userId string, riskSerial string) ([]string, error) {
	info, err := s.TFAInfoErr(ctx, userId)
	if err != nil {
		g.Log().Warning(ctx, "TFATx:", "userid:", userId, "riskSerial:", riskSerial)
		g.Log().Errorf(ctx, "%+v", err)
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

	return kind, nil
}
