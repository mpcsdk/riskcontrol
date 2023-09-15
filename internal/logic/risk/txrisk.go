package risk

import (
	"context"
	"encoding/json"
	"fmt"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"strings"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sRisk) checkTxs(ctx context.Context, address string, txs []*conrisk.RiskTx) (int32, error) {
	for _, tx := range txs {
		code, err := s.checkTx(ctx, tx)
		if err != nil {
			return -1, err
		}
		if code != 0 {
			return code, nil
		}
	}
	return 0, nil
}

func (s *sRisk) checkTx(ctx context.Context, riskTx *conrisk.RiskTx) (int32, error) {

	data := riskTx
	////has contractrisk cfg
	if rcfg, ok := contractRiskMap[data.Contract]; ok {
		if rcfg.Kind == "erc20" {
			cnt, err := rule_Token(ctx, rcfg.Contract, rcfg.MethodName, data)
			if err != nil {
				return 1, err
			}
			if cnt > rcfg.Threshold {
				return 1, nil
			}
			return 0, nil
		} else if rcfg.Kind == "erc721" {
			cnt, err := rule_nftcnt(ctx, rcfg.Contract, rcfg.MethodName, data)
			if err != nil {
				return 1, err
			}
			if cnt > rcfg.Threshold {
				return 1, nil
			}
			return 0, nil
		} else {
			return 1, gerror.NewCode(gcode.CodeInvalidParameter)
		}
	}
	//notice: default
	return 0, nil
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
	contractRiskMap = map[string]*contractRisk{}
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
				Contract:   strings.ToLower(valrisk["contract"].(string)),
				Kind:       valrisk["kind"].(string),
				MethodName: valrisk["methodName"].(string),
				Threshold:  int(Threshold),
			}
			contractRiskMap[r.Contract] = r
		}
	}
}
