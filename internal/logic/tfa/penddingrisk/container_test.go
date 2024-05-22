package pendding

import (
	"riskcontral/internal/logic/tfa/tfaconst"
	"testing"
	"time"

	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

func Test_newRiskPenddingContainer(t *testing.T) {

	rc := NewRiskPenddingContainer(10)
	rc.NewRiskPendding(&entity.Tfa{
		UserId: "userId",
	}, tfaconst.RiskKind_BindMail, nil)
	rv := rc.GetRiskPendding("userId", "riskSerial")
	if rv == nil {
		t.Error("riskPendding not exists")
	}
	time.Sleep(10 * time.Second)
	rv = rc.GetRiskPendding("userId", "riskSerial")
	if rv != nil {
		t.Error("riskPendding exists")
	}
	///
	time.Sleep(5 * time.Second)
	rc.NewRiskPendding(&entity.Tfa{
		UserId: "userId2",
	}, tfaconst.RiskKind_BindMail, nil)
	time.Sleep(10 * time.Second)
	rv = rc.GetRiskPendding("userId2", "riskSerial2")
	if rv == nil {
		t.Error("riskPendding not exists")
	}
}
