package tfa

import (
	"context"
	"riskcontral/internal/config"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

type UserRiskId string

func keyUserRiskId(userId string, riskSerial string) UserRiskId {
	return UserRiskId(userId + "keyUserRiskId" + riskSerial)
}

type sTFA struct {
	// riskClient riskv1.UserClient
	ctx                   context.Context
	riskPenddingContainer *model.RiskPenddingContainer
	////
}

// /
func new() *sTFA {

	ctx := gctx.GetInitCtx()
	//
	t := config.Config.Cache.VerificationCodeDuration
	s := &sTFA{
		//todo:
		riskPenddingContainer: model.NewRiskPenddingContainer(t),
		ctx:                   ctx,
	}
	///

	return s
}

///

func init() {
	service.RegisterTFA(new())
}

func (s *sTFA) TfaRiskKind(ctx context.Context, tfaInfo *entity.Tfa, riskSerial string) (model.RiskKind, error) {
	risk := s.riskPenddingContainer.GetRiskVerify(tfaInfo.UserId, riskSerial)
	if risk == nil {
		g.Log().Warning(ctx, "TfaRiskKind:", "tfaInfo:", tfaInfo, "riskSerial:", riskSerial)
		return "", mpccode.CodeParamInvalid()
	}
	return risk.RiskKind, nil
}
func (s *sTFA) TfaRiskTidy(ctx context.Context, tfaInfo *entity.Tfa, riskSerial string, riskKind model.RiskKind) ([]string, error) {
	///
	vlist := []string{}

	switch riskKind {
	case model.RiskKind_BindPhone:
		risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_BindPhone)
		verifier := newVerifierPhone(model.RiskKind_BindPhone, "")
		risk.AddVerifier(verifier)
		vlist = append(vlist, "phone")
		if tfaInfo.Mail != "" {
			verifier := newVerifierMail(model.RiskKind_BindPhone, tfaInfo.Mail)
			risk.AddVerifier(verifier)
			vlist = append(vlist, "mail")
		}
	case model.RiskKind_BindMail:
		risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_BindMail)
		verifier := newVerifierMail(model.RiskKind_BindMail, "")
		risk.AddVerifier(verifier)
		vlist = append(vlist, "mail")
		if tfaInfo.Phone != "" {
			verifier := newVerifierPhone(model.RiskKind_BindMail, tfaInfo.Phone)
			risk.AddVerifier(verifier)
			vlist = append(vlist, "phone")
		}
	case model.RiskKind_UpMail:
		risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_UpMail)
		verifier := newVerifierMail(model.RiskKind_UpMail, "")
		risk.AddVerifier(verifier)
		vlist = append(vlist, "mail")
		if tfaInfo.Phone != "" {
			verifier := newVerifierPhone(model.RiskKind_UpMail, tfaInfo.Phone)
			risk.AddVerifier(verifier)
			vlist = append(vlist, "phone")
		}
	case model.RiskKind_UpPhone:
		risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_UpPhone)
		verifier := newVerifierPhone(model.RiskKind_UpPhone, "")
		risk.AddVerifier(verifier)
		vlist = append(vlist, "phone")
		if tfaInfo.Mail != "" {
			verifier := newVerifierMail(model.RiskKind_UpPhone, tfaInfo.Mail)
			risk.AddVerifier(verifier)
			vlist = append(vlist, "mail")
		}
	}
	///
	return vlist, nil
}

func (s *sTFA) TFATx(ctx context.Context, tfaInfo *entity.Tfa, riskSerial string) ([]string, error) {

	if tfaInfo.Mail == "" && tfaInfo.Phone == "" {
		return nil, mpccode.CodeTFANotExist()
	}

	//
	kind := []string{}
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_Tx)
	if tfaInfo.Phone != "" {
		verifier := newVerifierPhone(model.RiskKind_Tx, tfaInfo.Phone)
		risk.AddVerifier(verifier)
		kind = append(kind, "phone")
	}

	if tfaInfo.Mail != "" {
		verifer := newVerifierMail(model.RiskKind_Tx, tfaInfo.Mail)
		risk.AddVerifier(verifer)
		kind = append(kind, "mail")
	}

	return kind, nil
}
