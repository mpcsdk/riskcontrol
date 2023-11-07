package tfa

import (
	"context"
	"riskcontral/internal/config"
	"riskcontral/internal/model/do"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

type UserRiskId string
type RiskKind string
type VerifyKind string

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

///

func init() {
	service.RegisterTFA(new())
}

func (s *sTFA) TFACreate(ctx context.Context, userId string, phone string, mail string, riskSerial string) ([]string, error) {
	///
	err := service.DB().InsertTfaInfo(ctx, userId, &do.Tfa{
		UserId:    userId,
		CreatedAt: gtime.Now(),
	})
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userId", userId),
			mpccode.ErrDetail("phone", phone),
			mpccode.ErrDetail("mail", mail),
		))
		return nil, err
	}
	return nil, nil

	// if err != nil || code != 0 {
	// kind := []string{}
	// var risk *riskVerifyPendding = nil
	// /// need verification
	// if phone != "" {
	// 	verifier := newVerifierPhone(RiskKind_BindPhone, phone)
	// 	risk = s.riskPenddingContainer.NewRiskPendding(userId, riskSerial, RiskKind_BindPhone)
	// 	risk.AddVerifier(verifier)
	// 	risk.AddAfterFunc(nil)
	// 	risk.AddAfterFunc(func(ctx context.Context) error {
	// 		return s.recordPhone(ctx, userId, phone, false)
	// 	})

	// 	kind = append(kind, "phone")
	// } else if mail != "" {
	// 	risk = s.riskPenddingContainer.NewRiskPendding(userId, riskSerial, RiskKind_BindMail)
	// 	verifier := newVerifierMail(RiskKind_BindPhone, mail)
	// 	risk.AddVerifier(verifier)
	// 	risk.AddAfterFunc(nil)
	// 	risk.AddAfterFunc(func(ctx context.Context) error {
	// 		return s.recordMail(ctx, userId, mail, false)
	// 	})

	// 	kind = append(kind, "mail")
	// }
	// ///

	// risk.AddBeforFunc(func(ctx context.Context) error {
	// 	return s.createTFA(ctx, userId, mail, phone)
	// })
	///
	// return kind, nil
}

func (s *sTFA) TFATx(ctx context.Context, userId string, riskSerial string) ([]string, error) {
	info, err := s.TFAInfo(ctx, userId)
	if err != nil {
		g.Log().Warning(ctx, "TFATx:", "userid:", userId, "riskSerial:", riskSerial)
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeInternalError)
	}
	if info == nil {
		g.Log().Warning(ctx, "TFATx:", "userid:", userId, "riskSerial:", riskSerial)
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}

	//
	kind := []string{}
	risk := s.riskPenddingContainer.NewRiskPendding(userId, riskSerial, RiskKind_Tx)
	if info.Phone != "" {
		verifier := newVerifierPhone(RiskKind_Tx, info.Phone)
		risk.AddVerifier(verifier)
		kind = append(kind, "phone")
	}

	if info.Mail != "" {
		verifer := newVerifierMail(RiskKind_Tx, info.Mail)
		risk.AddVerifier(verifer)
		kind = append(kind, "mail")
	}

	return kind, nil
}
