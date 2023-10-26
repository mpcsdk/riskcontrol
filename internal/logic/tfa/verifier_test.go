package tfa

import (
	"testing"

	"github.com/gogf/gf/v2/errors/gerror"
)

func Test_Verifier(t *testing.T) {
	verifier := newEmptyVerifier(RiskKind_Tx)
	_, err := verifier.SendVerificationCode()
	if !gerror.Equal(err, errRiskKindTx) {
		t.Error(err)
	}
}
