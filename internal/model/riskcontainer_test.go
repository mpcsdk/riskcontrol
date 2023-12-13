package model

import (
	"testing"
	"time"
)

func Test_newRiskPenddingContainer(t *testing.T) {

	rc := NewRiskPenddingContainer(10)
	rc.NewRiskPendding("userId", "riskSerial", RiskKind_BindMail)
	rv := rc.GetRiskVerify("userId", "riskSerial")
	if rv == nil {
		t.Error("riskPendding not exists")
	}
	time.Sleep(10 * time.Second)
	rv = rc.GetRiskVerify("userId", "riskSerial")
	if rv != nil {
		t.Error("riskPendding exists")
	}
	///
	time.Sleep(5 * time.Second)
	rc.NewRiskPendding("userId2", "riskSerial2",RiskKind_BindMail)
	time.Sleep(10 * time.Second)
	rv = rc.GetRiskVerify("userId2", "riskSerial2")
	if rv == nil {
		t.Error("riskPendding not exists")
	}
}
