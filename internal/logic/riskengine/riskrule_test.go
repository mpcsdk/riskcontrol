package riskengine

import (
	"riskcontrol/internal/model"
	"testing"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

func Test_nft(t *testing.T) {
	re := New()
	_, err := re.UpRules("nft", "nft", `
	chainIds = Slice.New(123)
	chainIds.Append(Context.ChainId)
	///
	befor24h = Time.TimeStampBefore(Time.Day)
	forRange i := SignTxs {
	  SignTx = SignTxs[i]
	  if Contract.IsNft(SignTx.Contract) {
		cnt = AggDB.NftAssetCnt(chainIds, SignTx.From, SignTx.Contract, befor24h)
		contractName = Contract.ContractName(SignTx.Contract)
		if contractName == "Weapon" {
		  if cnt.Cmp(100) > 0  {
			return RiskCode.Verify
		  }
		}else if contractName == "Miner"{
		  if cnt.Cmp(100) > 0  {
			return RiskCode.Verify
		  }
		}else if contractName == "Finshon"{
		  if cnt.Cmp(100) > 0  {
			return RiskCode.Verify
		  }
		}
	  }
	}
	
	return RiskCode.Pass
	`, 0)
	if err != nil {
		t.Fatal(err)
	}
	//
	ok, err := re.Exec(gctx.GetInitCtx(), "nft", &model.RiskExecData{
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
			ChainId: 111,
			CurTfa: &entity.Tfa{
				PhoneUpdatedAt: gtime.Now().Add(-2 * time.Hour),
				MailUpdatedAt:  gtime.Now().Add(-2 * time.Hour),
			},
		},
	})
	if err != nil {
		code := gerror.Code(err)
		if code != gcode.CodeNil {
			t.Error(code.Detail())
		}
	}
	if ok != mpccode.RiskCodePass {
		t.Fatal("riskCode:", ok)
	}
	t.Log("riskCode:", ok)
}

func Test_ft(t *testing.T) {
	re := New()
	_, err := re.UpRules("ft", "ft", `
	chainIds = Slice.New(123)
	chainIds.Append(Context.ChainId)
	///
	befor24h = Time.TimeStampBefore(Time.Day)
	forRange i := SignTxs {
	  SignTx = SignTxs[i]
	  if Contract.IsFt(SignTx.Contract) {
		cnt = AggDB.NftAssetCnt(chainIds, SignTx.From, SignTx.Contract, befor24h)
		contractName = Contract.ContractName(SignTx.Contract)
		if contractName == "usdt" {
		  if cnt.Cmp(100) == 0  {
			return RiskCode.Verify
		  }
		}else if contractName == "RPG"{
		  if cnt.Cmp(100) <  0  {
			return RiskCode.Verify
		  }
		}else if contractName == "MUD"{
		  if cnt.Cmp(100) > 0  {
			return RiskCode.Verify
		  }
		}
	  }
	}
	
	return RiskCode.Pass
	`, 0)
	if err != nil {
		t.Fatal(err)
	}
	//
	ok, err := re.Exec(gctx.GetInitCtx(), "ft", &model.RiskExecData{
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
			ChainId: 111,
			CurTfa: &entity.Tfa{
				PhoneUpdatedAt: gtime.Now().Add(-2 * time.Hour),
				MailUpdatedAt:  gtime.Now().Add(-2 * time.Hour),
			},
		},
	})
	if err != nil {
		code := gerror.Code(err)
		if code != gcode.CodeNil {
			t.Error(code.Detail())
		}
	}
	if ok != mpccode.RiskCodePass {
		t.Fatal("riskCode:", ok)
	}
	t.Log("riskCode:", ok)
}

func Test_bigint(t *testing.T) {
	re := New()
	_, err := re.UpRules("bigint", "bigint", `
	chainIds = NumberList.New(0)
	///
	a = BigInt.NewDecimal(1.2, 18)
	b = BigInt.NewDecimal(0.1, 18)
	c = BigInt.NewDecimal(1.1, 18)
	d = BigInt.NewInt(10000)
	if a.Cmp(b) < 0  {
		print("a>b")
		return RiskCode.Verify
	}
	if a.Cmp(c) < 0 {
		print("a>c")
		return RiskCode.Verify
	}
	if c.Cmp(b) < 0 {
		print("c<b")
		return RiskCode.Verify
	}
	if c.Cmp(d) < 0 {
		print("c<d")
		return RiskCode.Verify
	}
	if d.Cmp(c) > 0 {
		print("d>c")
		return RiskCode.Verify
	}
	if d.Cmp(BigInt.NewDecimal(1.1, 17)) > 0 {
		print("d>1.1")
		return RiskCode.Verify
	}

	return RiskCode.Pass
	`, 0)
	if err != nil {
		t.Fatal(err)
	}
	//
	ok, err := re.Exec(gctx.GetInitCtx(), "bigint", &model.RiskExecData{
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
			ChainId: 111,
			CurTfa: &entity.Tfa{
				PhoneUpdatedAt: gtime.Now().Add(-2 * time.Hour),
				MailUpdatedAt:  gtime.Now().Add(-2 * time.Hour),
			},
		},
	})
	if err != nil {
		code := gerror.Code(err)
		if code != gcode.CodeNil {
			t.Error(code.Detail())
		}
	}
	if ok != mpccode.RiskCodePass {
		t.Fatal("riskCode:", ok)
	}
	t.Log("riskCode:", ok)
}
