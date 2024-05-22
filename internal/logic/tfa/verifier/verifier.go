package verifier

import (
	"context"
	"errors"
	"fmt"
	"riskcontral/internal/logic/tfa/tfaconst"
	"riskcontral/internal/model"
)

// //

var errRiskKindTx = errors.New("riskKindTx")
var errRiskKindBindPhone = errors.New("riskKindBindPhone")
var errRiskKindBindMail = errors.New("riskKindBindMail")
var errRiskKindUpPhone = errors.New("riskKindUpPhone")
var errRiskKindUpMail = errors.New("riskKindUpMail")

type emptyVerifier struct {
	riskKind tfaconst.RISKKIND
}

func newEmptyVerifier(riskKind tfaconst.RISKKIND) tfaconst.IVerifier {
	return &emptyVerifier{
		riskKind: riskKind,
	}
}
func (s *emptyVerifier) Destination() string {
	return "emptyVerifier"
}
func (s *emptyVerifier) SendCompletion() error {
	fmt.Println(s.riskKind)
	switch s.riskKind {
	case tfaconst.RiskKind_Tx:
		return errRiskKindTx
	case tfaconst.RiskKind_BindPhone:
		return errRiskKindBindPhone
	case tfaconst.RiskKind_BindMail:
		return errRiskKindBindMail
	case tfaconst.RiskKind_UpPhone:
		return errRiskKindUpPhone
	case tfaconst.RiskKind_UpMail:
		return errRiskKindUpMail
	}
	return nil
}
func (s *emptyVerifier) SendVerificationCode() (string, error) {
	switch s.riskKind {
	case tfaconst.RiskKind_Tx:
		return "", errRiskKindTx
	case tfaconst.RiskKind_BindPhone:
		return "", errRiskKindBindPhone
	case tfaconst.RiskKind_BindMail:
		return "", errRiskKindBindMail
	case tfaconst.RiskKind_UpPhone:
		return "", errRiskKindUpPhone
	case tfaconst.RiskKind_UpMail:
		return "", errRiskKindUpMail
	}
	return "", nil
}
func (s *emptyVerifier) IsDone() bool {
	return true
}
func (s *emptyVerifier) VerifyKind() tfaconst.VERIFYKIND {
	return tfaconst.VerifierKind_Nil
}
func (s *emptyVerifier) RiskKind() tfaconst.RISKKIND {
	return tfaconst.RiskKind_Nil
}

func (s *emptyVerifier) SetCode(code string) {
}
func (s *emptyVerifier) Verify(ctx context.Context, verifierCode *model.VerifyCode) (tfaconst.RISKKIND, error) {
	return "", nil
}

// /
