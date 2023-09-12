package risk

import (
	"context"
	"encoding/json"
	"fmt"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sRisk) checkTx(ctx context.Context, riskName string, riskData interface{}) (bool, error) {
	if _, ok := riskData.(*conrisk.RiskTx); !ok {
		return false, gerror.NewCode(gcode.CodeInvalidParameter)
	}
	data := riskData.(*conrisk.RiskTx)
	////has contractrisk cfg
	if rcfg, ok := contractRiskMap[data.Contract]; ok {
		if rcfg.Kind == "erc20" {
			cnt, err := rule_Token(ctx, rcfg.Contract, rcfg.MethodName, data)
			if err != nil {
				return false, err
			}
			if cnt > rcfg.Threshold {
				return false, nil
			}
			return true, nil
		} else if rcfg.Kind == "erc721" {
			cnt, err := rule_nftcnt(ctx, rcfg.Contract, rcfg.MethodName, data)
			if err != nil {
				return false, err
			}
			if cnt > rcfg.Threshold {
				return false, nil
			}
			return true, nil
		} else {
			return false, gerror.NewCode(gcode.CodeInvalidParameter)
		}
	}

	return false, gerror.NewCode(gcode.CodeInvalidParameter)
}

// func rule_NFTCnt(ctx context.Context, data *conrisk.RiskTx) (int, error) {
// 	switch data.Contract {
// 	case "矿机":
// 		return rule_nftcnt(ctx, "usdt", "transfer", data)
// 	case "装备":
// 		return rule_nftcnt(ctx, "MUD", "transfer", data)
// 	case "时装":
// 		return rule_nftcnt(ctx, "MAK", "transfer", data)
// 	case "武器":
// 		return rule_nftcnt(ctx, "RPG", "safeTransferFrom", data)
// 	}
// 	return 0, gerror.New("rule_Token24HCnt")
// }

// 矿机、装备、时装、武器
func rule_nftcnt(ctx context.Context, tokenAddress string, methdoName string, data *conrisk.RiskTx) (int, error) {
	// values := []string{}
	rst, err := dao.EthTx.Ctx(ctx).Where(do.EthTx{
		Address:    data.Address,
		Target:     tokenAddress,
		MethodName: methdoName,
	}).
		WhereGT(dao.EthTx.Columns().CreatedAt, gtime.Now().Add(BeforH24)).Fields(
		dao.EthTx.Columns().Value,
	).
		Fields(dao.EthTx.Columns().Value).
		All()
	fmt.Println("rule_USDT24HCnt:", rst)
	if err != nil {
		return 0, err
	}
	///
	val := len(rst)
	return val, nil
}

// func rule_Token24HCnt(ctx context.Context, data *conrisk.RiskTx) (int, error) {
// 	switch data.Contract {
// 	case "USDT":
// 		return rule_Token(ctx, "usdt", "transfer", data)
// 	case "MUD":
// 		return rule_Token(ctx, "MUD", "transfer", data)
// 	case "MAK":
// 		return rule_Token(ctx, "MAK", "transfer", data)
// 	case "RPG":
// 		return rule_Token(ctx, "RPG", "safeTransferFrom", data)
// 	}
// 	return 0, gerror.New("rule_Token24HCnt")
// }

// MUD、MAK、USDT、RPG
func rule_Token(ctx context.Context, tokenAddress string, methdoName string, data *conrisk.RiskTx) (int, error) {
	rst, err := dao.EthTx.Ctx(ctx).Where(do.EthTx{
		Address:    data.Address,
		Target:     tokenAddress,
		MethodName: methdoName,
	}).
		WhereGT(dao.EthTx.Columns().CreatedAt, gtime.Now().Add(BeforH24)).Fields(
		dao.EthTx.Columns().Value,
	).
		Fields(dao.EthTx.Columns().Value).
		All()
	fmt.Println("rule_USDT24HCnt:", rst)
	if err != nil {
		return 0, err
	}
	///
	val := 0
	for _, v := range rst {
		val += v[dao.EthTx.Columns().Value].Int()
	}
	return val, nil

}

// var rulesFuncList map[string]func(ctx context.Context, data *conrisk.RiskTx) (int, error)

type contractRisk struct {
	Contract   string
	Kind       string
	MethodName string
	Threshold  int
}

var contractRiskMap map[string]*contractRisk

func init() {
	// rulesFuncList = make(map[string]func(ctx context.Context, data *conrisk.RiskTx) (int, error))
	// rulesFuncList["Token24HCnt"] = rule_Token24HCnt
	// rulesFuncList["Nft24HCnt"] = rule_NFTCnt
	////
	contractRiskMap := map[string]*contractRisk{}
	ctx := context.Background()
	riskcfg, err := gcfg.Instance().Get(ctx, "contractRisk")
	if err != nil {
		panic(err)
	}
	for _, val := range riskcfg.Array() {
		if valrisk, ok := val.(map[string]interface{}); !ok {
			panic(fmt.Errorf("contractRisk:%v", val))
		} else {
			Threshold, _ := valrisk["threshold"].(json.Number).Int64()
			r := &contractRisk{
				Contract:   valrisk["contract"].(string),
				Kind:       valrisk["kind"].(string),
				MethodName: valrisk["methodName"].(string),
				Threshold:  int(Threshold),
			}
			contractRiskMap[r.Contract] = r
		}
	}
}
