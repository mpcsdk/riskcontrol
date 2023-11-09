package tfa

import (
	"context"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sTFA) TfaSetMail(ctx context.Context, tfaInfo *entity.Tfa, mail string, riskSerial string, riskKind model.RiskKind) (string, error) {
	if riskKind == model.RiskKind_BindMail {
		return s.TfaBindMail(ctx, tfaInfo, mail, riskSerial)
	} else {
		return s.TfaUpMail(ctx, tfaInfo, mail, riskSerial)
	}
}

// //
func (s *sTFA) TfaBindMail(ctx context.Context, tfaInfo *entity.Tfa, mail string, riskSerial string) (string, error) {
	if tfaInfo == nil || tfaInfo.Mail != "" {
		return "", mpccode.CodeParamInvalid.Error()
	}
	verifier := newVerifierMail(model.RiskKind_BindMail, mail)
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_UpMail)

	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		err := s.recordMail(ctx, tfaInfo.UserId, mail, false)
		if err != nil {
			g.Log().Warning(ctx, "TfaBindMail recordMail err:", "userid:", tfaInfo.UserId, "mail:", mail, "err:", err)
			return err
		}
		return nil
	})
	///tfa phone if
	if tfaInfo.Phone != "" {
		verifier := newVerifierPhone(model.RiskKind_BindMail, tfaInfo.Phone)
		risk.AddVerifier(verifier)
	}

	return riskSerial, nil
}

func (s *sTFA) TfaUpMail(ctx context.Context, tfaInfo *entity.Tfa, mail string, riskSerial string) (string, error) {
	if tfaInfo == nil || tfaInfo.Mail == "" {
		return "", mpccode.CodeParamInvalid.Error()
	}
	verifier := newVerifierMail(model.RiskKind_BindMail, mail)
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_UpMail)

	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		err := s.recordMail(ctx, tfaInfo.UserId, mail, true)
		if err != nil {
			g.Log().Warning(ctx, "TfaUpMail recordMail err:", "userid:", tfaInfo.UserId, "mail:", mail, "err:", err)
			return err
		}
		return nil
	})
	///tfa phone if
	if tfaInfo.Phone != "" {
		verifier := newVerifierPhone(model.RiskKind_UpMail, tfaInfo.Phone)
		risk.AddVerifier(verifier)
	}
	//
	return riskSerial, nil
	//
}
