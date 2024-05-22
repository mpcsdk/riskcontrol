package pendding

import (
	"context"
	"riskcontral/internal/logic/tfa/tfaconst"
	"riskcontral/internal/logic/tfa/verifier"

	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	"github.com/mpcsdk/mpcCommon/rand"
)

func (s *RiskPendding) RiskTxTidy(ctx context.Context, tfaInfo *entity.Tfa) (string, []string, error) {
	riskSerial := rand.GenNewSid()
	if tfaInfo.Mail == "" && tfaInfo.Phone == "" {
		return riskSerial, nil, mpccode.CodeTFANotExist()
	}
	//
	// risk := s.NewRiskPendding(tfaInfo, tfaconst.RiskKind_Tx, nil)
	if tfaInfo.Phone != "" {
		verifier := verifier.NewVerifierPhone(tfaconst.RiskKind_Tx, tfaInfo.Phone)
		s.AddVerifier(verifier)
	}

	if tfaInfo.Mail != "" {
		verifer := verifier.NewVerifierMail(tfaconst.RiskKind_Tx, tfaInfo.Mail)
		s.AddVerifier(verifer)
	}

	return s.riskSerial, s.verifyKind, nil
}
