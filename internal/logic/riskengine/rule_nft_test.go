package riskengine

import (
	"math/big"
	"riskcontrol/internal/model"
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

func Test_nftcnt(t *testing.T) {
	re := New()
	ruleName, err := re.UpRules("Test_nftcnt", "Test_nftcnt", `
	chainIds = NumberList.New(Context.ChainId)
	///
	befor24h = Time.TimeStampBefore(Time.Day)
	forRange i := SignTxs {
		SignTx = SignTxs[i]

		asset = AggDB.AssetAttr(Context.ChainId, SignTx)
		cnt = AggDB.AssetNftCnt(chainIds, SignTx.From, asset, befor24h)
		print(cnt)
		print(asset.Value)
		if asset.Symbol == "Weapon" {
			if cnt > 5  {
				return RiskCode.Verify
			}
		}else if asset.Symbol == "Energy Core"{
			if cnt + asset.Value > 10 {
				return RiskCode.Verify
			}
		}else if asset.Symbol == "MUD"{
			if cnt > 10  {
				return RiskCode.Verify
			}
		}
	}
	
	return RiskCode.Pass
	`, 0)
	if err != nil {
		t.Fatal(err)
	}
	//
	ok, err := re.ExecRule(ctx, ruleName, &model.RiskExecData{
		SignTxs: []*analzyer.AnalzyedSignTx{{
			//usdt
			From:       "0x12345678901234567890123456789",
			MethodName: "transfer",
			Target:     "0x0F3A62dB02F743b549053cc8d538C65aB01E3618",
		}, {
			//weapon
			From:       "0x12345678901234567890123456789",
			MethodName: "transfer",
			Target:     "0x61b67130f8519347DCA83BA4A98Df54C01c5eeDa",
		}, {
			///energy core
			Target:  "0x9a381F5d5BEeBd7C04Efb90D5854798F6C96b715",
			From:    "0x12345678901234567890123456789",
			Value:   (*analzyer.BigInt)(big.NewInt(2)),
			TokenId: (*analzyer.BigInt)(big.NewInt(2)),
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
	if ok != mpccode.RiskCodePass {
		t.Fatal("riskCode:", ok)
	}
	t.Log("riskCode:", ok)
}
