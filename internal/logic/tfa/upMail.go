package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) TFAUpMail(ctx context.Context, tfaInfo *entity.Tfa, mail string, riskSerial string) (string, error) {
	//
	////
	var verifier IVerifier = nil
	mailExists := false
	if tfaInfo.Mail == "" {
		mailExists = false
		verifier = newVerifierMail(RiskKind_BindMail, mail)
	} else {
		mailExists = true
		verifier = newVerifierMail(RiskKind_UpMail, mail)
	}

	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, RiskKind_UpMail)
	///
	risk.AddVerifier(verifier)
	risk.AddAfterFunc(func(ctx context.Context) error {
		err := s.recordMail(ctx, tfaInfo.UserId, mail, mailExists)
		if err != nil {
			g.Log().Warning(ctx, "TFAUpMail recordMail err:", "userid:", tfaInfo.UserId, "mail:", mail, "mailExists:", mailExists, "err:", err)
			return err
		}
		return nil
	})
	///tfa phone if
	if tfaInfo.Phone != "" {
		verifier := newVerifierPhone(RiskKind_UpMail, tfaInfo.Phone)
		risk.AddVerifier(verifier)
	}
	//
	return riskSerial, gerror.NewCode(consts.CodePerformRiskNeedVerification)
	//
}
