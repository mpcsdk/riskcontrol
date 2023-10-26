package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sTFA) TFAUpPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string) (string, error) {

	/// need verification
	phoneExists := false
	if tfaInfo.Phone == "" {
		phoneExists = false
	} else {
		phoneExists = true
	}
	// //
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, RiskKind_UpPhone)
	verifier := newVerifierPhone(RiskKind_BindPhone, phone)
	//
	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		return s.recordPhone(ctx, tfaInfo.UserId, phone, phoneExists)
	})

	// //
	// ///tfa mailif
	if tfaInfo.Mail != "" {
		verifier := newVerifierMail(RiskKind_BindPhone, tfaInfo.Mail)
		risk.AddVerifier(verifier)
	}
	///
	return riskSerial, gerror.NewCode(consts.CodePerformRiskNeedVerification)

}
