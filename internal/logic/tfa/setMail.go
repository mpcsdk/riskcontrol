package tfa

import (
	"context"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"

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
		return "", mpccode.CodeParamInvalid()
	}
	verifier := newVerifierMail(model.RiskKind_BindMail, mail)
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_UpMail)

	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		err := s.recordMail(ctx, tfaInfo.UserId, mail, false)
		return err
	})

	return riskSerial, nil
}

func (s *sTFA) TfaUpMail(ctx context.Context, tfaInfo *entity.Tfa, mail string, riskSerial string) (string, error) {
	if tfaInfo == nil || tfaInfo.Mail == "" {
		return "", mpccode.CodeParamInvalid()
	}
	verifier := newVerifierMail(model.RiskKind_UpMail, mail)
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_UpMail)

	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		err := s.recordMail(ctx, tfaInfo.UserId, mail, true)
		return err
	})

	//
	return riskSerial, nil
	//
}
