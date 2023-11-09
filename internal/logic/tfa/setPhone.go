package tfa

import (
	"context"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
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
		return "", mpccode.CodeParamInvalid.Error()
	}
	verifier := newVerifierPhone(model.RiskKind_BindPhone, phone)
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_BindPhone)

	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		err := s.recordPhone(ctx, tfaInfo.UserId, phone, false)
		if err != nil {
			g.Log().Warning(ctx, "TfaBindPhone recordMail err:", "userid:", tfaInfo.UserId, "phone:", phone, "err:", err)
			return err
		}
		return nil
	})
	///tfa mail if
	if tfaInfo.Mail != "" {
		verifier := newVerifierMail(model.RiskKind_BindPhone, tfaInfo.Mail)
		risk.AddVerifier(verifier)
	}
	//
	return riskSerial, nil
}

func (s *sTFA) TfaUpPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string) (string, error) {
	if tfaInfo == nil || tfaInfo.Phone == "" {
		return "", mpccode.CodeParamInvalid.Error()
	}
	verifier := newVerifierPhone(model.RiskKind_UpPhone, phone)
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_UpPhone)

	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		err := s.recordPhone(ctx, tfaInfo.UserId, phone, true)
		if err != nil {
			g.Log().Warning(ctx, "TfaUpPhone recordMail err:", "userid:", tfaInfo.UserId, "phone:", phone, "err:", err)
			return err
		}
		return nil
	})
	///tfa mail if
	if tfaInfo.Mail != "" {
		verifier := newVerifierMail(model.RiskKind_UpPhone, tfaInfo.Mail)
		risk.AddVerifier(verifier)
	}
	//
	return riskSerial, nil
	//
}
