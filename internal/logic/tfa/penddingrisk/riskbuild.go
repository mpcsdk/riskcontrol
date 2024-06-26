package pendding

import (
	"context"
	v1 "riskcontrol/api/tfa/v1"
	"riskcontrol/internal/logic/tfa/tfaconst"
	"riskcontrol/internal/logic/tfa/verifier"

	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	"github.com/mpcsdk/mpcCommon/rand"
)

// ///
// ////
func (s *RiskPendding) build(tfaInfo *entity.Tfa, data *v1.RequestData) (string, []string, error) {
	///
	vlist := []string{}
	riskSerial := rand.GenNewSid()

	switch s.riskKind {
	case tfaconst.RiskKind_BindPhone:
		ver := verifier.NewVerifierPhone(tfaconst.RiskKind_BindPhone, "")
		s.AddVerifier(ver)
		vlist = append(vlist, "phone")
		if tfaInfo.Mail != "" {
			ver := verifier.NewVerifierMail(tfaconst.RiskKind_BindPhone, tfaInfo.Mail)
			s.AddVerifier(ver)
			vlist = append(vlist, "mail")
		}
	case tfaconst.RiskKind_BindMail:
		ver := verifier.NewVerifierMail(tfaconst.RiskKind_BindMail, "")
		s.AddVerifier(ver)
		vlist = append(vlist, "mail")
		if tfaInfo.Phone != "" {
			ver := verifier.NewVerifierPhone(tfaconst.RiskKind_BindMail, tfaInfo.Phone)
			s.AddVerifier(ver)
			vlist = append(vlist, "phone")
		}
	case tfaconst.RiskKind_UpMail:
		ver := verifier.NewVerifierMail(tfaconst.RiskKind_UpMail, "")
		s.AddVerifier(ver)
		vlist = append(vlist, "mail")
		if tfaInfo.Phone != "" {
			ver := verifier.NewVerifierPhone(tfaconst.RiskKind_UpMail, tfaInfo.Phone)
			s.AddVerifier(ver)
			vlist = append(vlist, "phone")
		}
	case tfaconst.RiskKind_UpPhone:
		ver := verifier.NewVerifierPhone(tfaconst.RiskKind_UpPhone, "")
		s.AddVerifier(ver)
		vlist = append(vlist, "phone")
		if tfaInfo.Mail != "" {
			ver := verifier.NewVerifierMail(tfaconst.RiskKind_UpPhone, tfaInfo.Mail)
			s.AddVerifier(ver)
			vlist = append(vlist, "mail")
		}
	case tfaconst.RiskKind_TfaRisk:
		if tfaInfo.Mail != "" {
			verifer := verifier.NewVerifierMail(tfaconst.RiskKind_TfaRisk, tfaInfo.Mail)
			s.AddVerifier(verifer)
			vlist = append(vlist, "mail")
		}
		if tfaInfo.Phone != "" {
			ver := verifier.NewVerifierPhone(tfaconst.RiskKind_TfaRisk, tfaInfo.Phone)
			s.AddVerifier(ver)
			vlist = append(vlist, "phone")
		}
	case tfaconst.RiskKind_TxNeedVerify:
		if tfaInfo.Mail != "" {
			verifer := verifier.NewVerifierMail(tfaconst.RiskKind_TxNeedVerify, tfaInfo.Mail)
			s.AddVerifier(verifer)
			vlist = append(vlist, "mail")
			//todo: enable
			s.AddAfterFunc(func(ctx context.Context) error {
				return recordPerson(ctx, tfaInfo.UserId, data.Enable)
			})
		}
	case tfaconst.RiskKind_Tx:
		if tfaInfo.Phone != "" {
			verifier := verifier.NewVerifierPhone(tfaconst.RiskKind_Tx, tfaInfo.Phone)
			s.AddVerifier(verifier)
			vlist = append(vlist, "phone")
		}

		if tfaInfo.Mail != "" {
			verifer := verifier.NewVerifierMail(tfaconst.RiskKind_Tx, tfaInfo.Mail)
			s.AddVerifier(verifer)
			vlist = append(vlist, "mail")
		}

	}
	///
	return riskSerial, vlist, nil
}
