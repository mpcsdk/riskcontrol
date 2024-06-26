package riskengine

import (
	"fmt"
	"riskcontrol/internal/logic/riskengine/intrinsic"
	"riskcontrol/internal/model"
	"testing"
	"time"

	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"

	"github.com/bilibili/gengine/builder"

	gcontext "context"
)

var ctx = gcontext.Background()

// 定义想要注入的结构体
type User struct {
	Name string
	Age  int64
	Male bool
}

func (u *User) GetNum(i int64) int64 {
	return i
}

func (u *User) Print(s string) {
	fmt.Println(s)
}

func (u *User) Say() {
	fmt.Println("hello world")
}

func Test_BuildRule1(t *testing.T) {

	dataContext := context.NewDataContext()
	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	err := ruleBuilder.BuildRuleFromString(`
	rule "name test1" "i can" salience 0
	BEGIN
		return 1
	END
	
	rule "name test2" "i can" salience 2
	BEGIN
		return 2
	END
	`)
	if err != nil {
		panic(err)
	}

	rule := ruleBuilder.Kc.RuleEntities
	for k, v := range rule {
		fmt.Println(k)
		fmt.Println(v.RuleName)
		fmt.Println(v.RuleDescription)
		fmt.Println(v.Salience)
		fmt.Println(v.RuleContent)
	}
	eng := engine.NewGengine()
	//执行规则
	err = eng.Execute(ruleBuilder, true)
	rst, _ := eng.GetRulesResultMap()
	fmt.Println(rst)

}

func Test_BuildRule(t *testing.T) {

	dataContext := context.NewDataContext()
	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	err := ruleBuilder.BuildRuleFromString(`
	rule "name test" "i can" salience 0
	BEGIN
				User.Name = "abc"
	END`)
	if err != nil {
		panic(err)
	}

	rule := ruleBuilder.Kc.RuleEntities
	for k, v := range rule {
		fmt.Println(k)
		fmt.Println(v.RuleName)
		fmt.Println(v.RuleDescription)
		fmt.Println(v.Salience)
		fmt.Println(v.RuleContent)
	}
	// eng := engine.NewGengine()
	// //执行规则
	// err = eng.Execute(ruleBuilder, true)
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

// 定义规则
const rule1 = `
rule "name test" "i can"  salience 1
begin
        if 7 == User.GetNum(7){
            User.Age = User.GetNum(89767) + 10000000
            User.Print("6666")
        }else{
            User.Name = "yyyy"
        }
end


rule "name 1" "i can"  salience 0
begin
	if AccountType == 1 {
		return true
	}else{
		return false
	}
end
`

func Test_Multi(t *testing.T) {
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}

	dataContext := context.NewDataContext()
	//注入初始化的结构体
	dataContext.Add("User", user)
	dataContext.Add("AccountType", 2)

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//构建规则
	err := ruleBuilder.BuildRuleFromString(rule1) //string(bs)

	if err != nil {
		t.Fatal(err)
	} else {
		eng := engine.NewGengine()
		//执行规则
		err := eng.Execute(ruleBuilder, true)
		// eng.ExecuteSelectedRules()
		println(user.Age)
		if err != nil {
			t.Fatal("execute rule error: ", err)
		}
	}
}

func Test_badmannar(t *testing.T) {
	dataContext := context.NewDataContext()
	//注入初始化的结构体

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//构建规则
	err := ruleBuilder.BuildRuleFromString(`
		test = 234
		324jf + 234
	`) //string(bs)
	if err != nil {
		t.Error(err)
	}
	if err == nil {
		t.Fatal(err)
	}

}

const rule_intrinsic = `
rule "rule_intrinsic" "i can"  salience 0
begin
	cnt = AggDB.RuleCnt("address", "contract", "method")
	if cnt > 0 {
		return true
	}
	return false
end
`
const rule_intrinsic2 = `
rule "rule_intrinsic2" "i can"  salience 0
begin
	cnt = AggDB.RuleCnt("address", "contract", "method")
	if cnt > 0 {
		return true
	}
	return false
end
`

func Test_Intrinsic(t *testing.T) {

	p, err := engine.NewGenginePool(10, 100, 2, rule_intrinsic, nil) //, map[string]interface{}{"DB": &intrinsic.Intrinsic{}})
	if err != nil {
		t.Error(err)
	}

	//构建规则
	err, rst := p.Execute(nil, false)
	if err != nil {
		t.Errorf("err:%s ", err)
	} else {
		t.Logf("rst:%v ", rst)
	}

}
func Test_Pool_up(t *testing.T) {

	p, err := engine.NewGenginePool(10, 100, 2, rule_intrinsic, map[string]interface{}{"DB": &intrinsic.Intrinsic{}})
	if err != nil {
		t.Error(err)
	}

	// err = p.UpdatePooledRules(rule_intrinsic2)
	// if err != nil {
	// 	t.Error(err)
	// }
	err = p.UpdatePooledRulesIncremental(rule_intrinsic2)
	if err != nil {
		t.Error(err)
	}
	err, rst := p.ExecuteSelectedRules(nil, []string{"rule_intrinsic"})
	if err != nil {
		t.Error(err)
	}
	t.Log(rst)
	err, rst = p.ExecuteSelectedRules(nil, []string{"rule_intrinsic2"})
	if err != nil {
		t.Error(err)
	}
	t.Log(rst)
}

func Test_nilerr(t *testing.T) {
	re := New()
	_, err := re.UpRules("ruleName", "ruleName", `
	if !isNil(CurTfa.PhoneUpdatedAt) {
		if Time.After(CurTfa.PhoneUpdatedAt, Time.NowAdd(Time.H* -1)){
			return RiskCode.Forbidden
		}
	}
	return RiskCode.Pass
	`, 0)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := re.ExecTx(ctx, &model.RiskExecData{
		SignTxs: nil,
	})
	if err != nil {
		ok := mpccode.Equal(err, mpccode.CodePerformRiskTimeOut())
		t.Fatal(ok)
	}
	if ok != mpccode.RiskCodeForbidden {
		t.Fatal("riskCode:", ok)
	}
}
func Test_tfa_Forbidden(t *testing.T) {
	re := New()
	_, err := re.UpRules("ruleName", "ruleName", `
	if !isNil(CurTfa.PhoneUpdatedAt) {
		if Time.After(CurTfa.PhoneUpdatedAt, Time.NowAdd(Time.H* -1)){
			return RiskCode.Forbidden
		}
	}
	if !isNil(CurTfa.MailUpdatedAt){
		if Time.After(CurTfa.MailUpdatedAt, Time.NowAdd(Time.H* -1)){
			return RiskCode.Forbidden
		}
	}
	forRange i := SignTxs {
		SignTx = SignTxs[i]
		if !Contract.IsNft(SignTx.Target) {
			return RiskCode.Verify
		}
	}
	return RiskCode.Pass
	`, 0)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := re.ExecTx(ctx, &model.RiskExecData{
		SignTxs: []*analzyer.AnalzyedSignTx{{
			From:       "0x12345678901234567890123456789",
			MethodName: "transfer",
			Target:     "0x12345678901234567890123456",
		}},
		Context: &model.RiskContext{
			CurTfa: &entity.Tfa{
				PhoneUpdatedAt: gtime.Now(),
				MailUpdatedAt:  gtime.Now().Add(-1 * time.Hour),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if ok != mpccode.RiskCodeForbidden {
		t.Fatal("riskCode:", ok)
	}
}
func Test_ctrl(t *testing.T) {
	re := New()
	re.UpRules("ruleName", "ruleName", `
	forRange i := SignTxs {
		SignTx := SignTxs[i]
	if !Contract.IsNft(SignTx.Target) {
		return true
	}
	threshold = Contract.RuleThreshold(SignTx.Target)
	riskRule = Contract.ContractRule(SignTx.Target)
	cnt = AggDB.RuleCnt(SignTx.Target, SignTx.From, riskRule.EventName)
	if riskRule.Threshold.Cmp(cnt) > 0{
		return true
	}
	}
	return false
	`, 0)
	ok, err := re.ExecTx(ctx, &model.RiskExecData{
		SignTxs: []*analzyer.AnalzyedSignTx{{
			From:       "0x12345678901234567890123456789",
			MethodName: "transfer",
			Target:     "0x12345678901234567890123456",
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if ok != mpccode.RiskCodePass {
		t.Fatal(ok)
	}
	////
	re.UpRules("ruleName", "ruleName", `
	if !Contract.IsNft(SignTx.Target) {
		return true
	}
	threshold = Contract.RuleThreshold(SignTx.Target)
	riskRule = Contract.ContractRule(SignTx.Target)
	cnt = AggDB.RuleCnt(SignTx.Target, SignTx.From, riskRule.EventName)
	if riskRule.Threshold.Cmp(cnt) < 0{
		return true
	}
	return false
	`, 0)
	ok, err = re.ExecTx(ctx, &model.RiskExecData{
		SignTxs: []*analzyer.AnalzyedSignTx{{
			From:       "0x12345678901234567890123456789",
			MethodName: "transfer",
			Target:     "0x12345678901234567890123456",
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if ok == mpccode.RiskCodePass {
		t.Fatal(ok)
	}
}
func Test_tfa_fortxs(t *testing.T) {
	// m := map[string]string{
	// 	"sdf": "s",
	// }
	// d := m["sdf"]
	// fmt.Println(d)
	re := New()
	_, err := re.UpRules("Test_tfa_fortxs", "Test_tfa_fortxs", `
	print(m)
	if !isNil(CurTfa.PhoneUpdatedAt) {
		if Time.After(CurTfa.PhoneUpdatedAt, Time.NowAdd(Time.H * -1)){
			return RiskCode.Forbidden
		}
	}
	if !isNil(CurTfa.MailUpdatedAt){
		if Time.After(CurTfa.MailUpdatedAt, Time.NowAdd(Time.H * -1)){
			return RiskCode.Forbidden
		}
	}

	forRange i := SignTxs {
		SignTx = SignTxs[i]
		if Contract.IsNft(SignTx.Target) {
			threshold = Contract.RuleThreshold(SignTx.Target)
			riskRule = Contract.ContractRule(SignTx.Target)
			cnt = AggDB.AggNft24HCnt(SignTx.From, SignTx.Target, riskRule.EventName)
			print("IsNft:", SignTx.Target, "cnt:", cnt)
			if riskRule.Threshold.Cmp(cnt) < 0{
				return RiskCode.Verify
			}
			
		}else if Contract.IsFt(SignTx.Target) {
			threshold = Contract.RuleThreshold(SignTx.Target)
			riskRule = Contract.ContractRule(SignTx.Target)
			cnt = AggDB.AggFt24HCnt(SignTx.From, SignTx.Target, riskRule.EventName)
			print("IsFt:", SignTx.Target, "cnt:", cnt)
			if riskRule.Threshold.Cmp(cnt) < 0{
				return RiskCode.Verify
			}
		}else{
			print("isUnknow:", SignTx.Target)
		}
	}

	return RiskCode.Pass
	`, 0)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := re.ExecTx(ctx, &model.RiskExecData{
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
			CurTfa: &entity.Tfa{
				PhoneUpdatedAt: gtime.Now().Add(-2 * time.Hour),
				MailUpdatedAt:  gtime.Now().Add(-2 * time.Hour),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if ok != mpccode.RiskCodeForbidden {
		t.Fatal("riskCode:", ok)
	}
	t.Log("riskCode:", ok)
}

func Test_mul_txrule(t *testing.T) {
	re := New()
	_, err := re.UpRules("Test_tx1", "Test_tx1", `
	return RiskCode.Pass
	`, 0)
	if err != nil {
		t.Fatal(err)
	}
	//

	_, err = re.UpRules("Test_tx2", "Test_tx2", `
	return RiskCode.Verify
	`, 0)
	if err != nil {
		t.Fatal(err)
	}
	//
	ok, err := re.ExecTx(ctx, &model.RiskExecData{
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
func Test_time(t *testing.T) {
	re := New()
	_, err := re.UpRules("Test_time", "Test_time", `
		if !isNil(CurTfa.MailUpdatedAt){
		  if Time.After(CurTfa.MailUpdatedAt, Time.NowAdd(Time.Hour * -3)){
		    return RiskCode.Forbidden
		  }
		}
		return RiskCode.Pass
	`, 0)
	if err != nil {
		t.Fatal(err)
	}
	//
	ok, err := re.ExecTx(ctx, &model.RiskExecData{
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
