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

func Test_NativeCoin(t *testing.T) {
	re := New()
	ruleName, err := re.UpRules("Test_NativeCoin", "Test_NativeCoin", `
	chainIds = NumberList.New(Context.ChainId)
	///
	befor24h = Time.TimeStampBefore(Time.Day)
	forRange i := SignTxs {
		SignTx = SignTxs[i]
		asset = AggDB.AssetAttr(Context.ChainId, SignTx)
		cnt = AggDB.AssetFtSum(chainIds, SignTx.From, asset, befor24h)

		if asset.Symbol == "RPG" {
			print(cnt)
			print(asset.Value)
			if asset.Value > 0.000000001{
				return RiskCode.Verify
			}
			// if cnt.Cmp(BigInt.NewDecimal(5, 18)) < 0  {
			// 	return RiskCode.Verify
			// }
	  	}else if asset.Symbol == "USDT"{
			if cnt + asset.Value > 5 {
				return RiskCode.Verify
			}
		}else if asset.Symbol == "MUD"{
			if cnt + asset.Value > 10  {
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
			//rpc
			From:       "0x12345678901234567890123456789",
			MethodName: "transfer",
			Target:     "",
			Value:      (*analzyer.BigInt)(big.NewInt(1000000000)),
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

func Test_ftcnt(t *testing.T) {
	re := New()
	ruleName, err := re.UpRules("Test_ftcnt", "Test_ftcnt", `
	chainIds = NumberList.New(Context.ChainId)
	///
	befor24h = Time.TimeStampBefore(Time.Day)
	forRange i := SignTxs {
		SignTx = SignTxs[i]

		asset = AggDB.AssetAttr(Context.ChainId, SignTx.Target)
		cnt = AggDB.NftSendCnt(chainIds, SignTx.From, asset, befor24h)

		if asset.Symbol == "USDT" {
			if cnt.CmpInt(0) == 0  {
			return RiskCode.Verify
			}
		}else if asset.Symbol == "RPG"{
			if cnt.CmpInt(100) <  0  {
			return RiskCode.Verify
			}
		}else if asset.Symbol == "MUD"{
			if cnt.CmpInt(100) > 0  {
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
	if ok != mpccode.RiskCodePass {
		t.Fatal("riskCode:", ok)
	}
	t.Log("riskCode:", ok)
}
