package riskengine

import (
	"riskcontrol/internal/model"
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

func Test_slice(t *testing.T) {
	re := New()
	_, err := re.UpRules("Test_slice", "Test_slice", `
	ids = NumberList.New(123)
	ids.Append(444)
	print(ids)
	`, 0)
	if err != nil {
		t.Fatal(err)
	}
	_, err = re.ExecTx(ctx, &model.RiskExecData{
		SignTxs: []*analzyer.AnalzyedSignTx{{
			//rpg
			From:       "0x12345678901234567890123456789",
			MethodName: "transfer",
			Target:     "0x71d9CFd1b7AdB1E8eb4c193CE6FFbe19B4aeE0dB",
		}, {
			//weapon
			From:       "0x12345678901234567890123456789",
			MethodName: "transfer",
			Target:     "0xb1682c08BEb47328D4f98AC08d3Cd01679ff5C3b",
		}},
		Context: &model.RiskContext{
			ChainId: 9527,
			CurTfa: &entity.Tfa{
				PhoneUpdatedAt: gtime.Now().Add(-2 * time.Hour),
				MailUpdatedAt:  gtime.Now().Add(-2 * time.Hour),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}
