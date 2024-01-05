package tfa

import (
	"context"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"

	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/rand"
)

func (s *sTFA) RiskTxTidy(ctx context.Context, tfaInfo *entity.Tfa) (string, []string, error) {
	riskSerial := rand.GenNewSid()
	if tfaInfo.Mail == "" && tfaInfo.Phone == "" {
		return riskSerial, nil, mpccode.CodeTFANotExist()
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

	return riskSerial, kind, nil
}
