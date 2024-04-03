package tfa

import (
	"context"
	"riskcontral/internal/model"

	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

func (s *sTFA) TfaSetPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string, riskKind model.RiskKind) (string, error) {
	if riskKind == model.RiskKind_BindPhone {
		return s.TfaBindPhone(ctx, tfaInfo, phone, riskSerial)
	} else {
		return s.TfaUpPhone(ctx, tfaInfo, phone, riskSerial)
	}
}

// //
func (s *sTFA) TfaBindPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string) (string, error) {
	if tfaInfo == nil || tfaInfo.Phone != "" {
		return "", mpccode.CodeParamInvalid()
	}
	verifier := newVerifierPhone(model.RiskKind_BindPhone, phone)
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_BindPhone)

	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		err := s.recordPhone(ctx, tfaInfo.UserId, phone, false)
		return err
	})

	//
	return riskSerial, nil
}

func (s *sTFA) TfaUpPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string) (string, error) {
	if tfaInfo == nil || tfaInfo.Phone == "" {
		return "", mpccode.CodeParamInvalid()
	}
	verifier := newVerifierPhone(model.RiskKind_UpPhone, phone)
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_UpPhone)

	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		err := s.recordPhone(ctx, tfaInfo.UserId, phone, true)
		return err
	})

	//
	return riskSerial, nil
	//
}
