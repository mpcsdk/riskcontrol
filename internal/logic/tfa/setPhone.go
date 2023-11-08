package tfa

import (
	"context"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sTFA) TfaSetPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string, codetype string) (string, error) {
	if codetype == model.Type_TfaBindPhone {
		return s.TfaBindPhone(ctx, tfaInfo, phone, riskSerial)
	} else {
		return s.TfaUpPhone(ctx, tfaInfo, phone, riskSerial)
	}
}

// //
func (s *sTFA) TfaBindPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string) (string, error) {
	if tfaInfo == nil || tfaInfo.Phone != "" {
		return "", mpccode.CodeParamInvalid.Error()
	}
	verifier := newVerifierPhone(RiskKind_BindPhone, phone)
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, RiskKind_BindPhone)

	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		err := s.recordPhone(ctx, tfaInfo.UserId, phone, false)
		if err != nil {
			g.Log().Warning(ctx, "TfaBindPhone recordMail err:", "userid:", tfaInfo.UserId, "phone:", phone, "err:", err)
			return err
		}
		return nil
	})
	///tfa phone if
	if tfaInfo.Phone != "" {
		verifier := newVerifierPhone(RiskKind_UpPhone, tfaInfo.Phone)
		risk.AddVerifier(verifier)
	}
	//
	return riskSerial, nil
}

func (s *sTFA) TfaUpPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string) (string, error) {
	if tfaInfo == nil || tfaInfo.Phone == "" {
		return "", mpccode.CodeParamInvalid.Error()
	}
	verifier := newVerifierPhone(RiskKind_BindPhone, phone)
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, RiskKind_UpPhone)

	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		err := s.recordMail(ctx, tfaInfo.UserId, phone, true)
		if err != nil {
			g.Log().Warning(ctx, "TfaUpPhone recordMail err:", "userid:", tfaInfo.UserId, "phone:", phone, "err:", err)
			return err
		}
		return nil
	})
	///tfa phone if
	if tfaInfo.Phone != "" {
		verifier := newVerifierPhone(RiskKind_UpPhone, tfaInfo.Phone)
		risk.AddVerifier(verifier)
	}
	//
	return riskSerial, nil
	//
}
