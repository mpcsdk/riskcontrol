package tfa

import (
	"context"
	"riskcontral/internal/model/entity"
)

func (s *sTFA) TFABindPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string) (string, error) {

	return "", nil
}
func (s *sTFA) TFAUpPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string) (string, error) {

	/// need verification
	var verifier IVerifier = nil
	phoneExists := false

	if tfaInfo == nil {
		verifier = newVerifierPhone(RiskKind_BindPhone, phone)
		phoneExists = false
	} else {
		verifier = newVerifierPhone(RiskKind_UpPhone, phone)
		if tfaInfo.Phone != "" {
			phoneExists = true
		}
	}

	// //
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, RiskKind_UpPhone)
	//
	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		return s.recordPhone(ctx, tfaInfo.UserId, phone, phoneExists)
	})

	// //
	// ///tfa mailif
	if tfaInfo.Mail != "" {
		verifier := newVerifierMail(RiskKind_UpPhone, tfaInfo.Mail)
		risk.AddVerifier(verifier)
	}
	///
	return riskSerial, nil
}
