package risk

import (
	"context"
	"fmt"
	"math/big"
	analyzsigndata "riskcontral/common/ethtx/analyzSignData"
	"riskcontral/internal/consts"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sRisk) checkTxs(ctx context.Context, signTx string) (int32, error) {
	atx, err := s.analzer.AnalzySignTxData(signTx)

	if err != nil {
		return consts.RiskCodeError, err
	}

	for _, tx := range atx.Txs {
		code, err := s.checkTx(ctx, atx.Address, tx)
		if err != nil {
			return consts.RiskCodeError, err
		}
		if code != consts.RiskCodePass {
			return code, nil
		}
	}
	return consts.RiskCodePass, nil
}

func (s *sRisk) checkTx(ctx context.Context, from string, riskTx *analyzsigndata.AnalzyedTxData) (int32, error) {

	////has contractrisk cfg
	if rcfg, ok := s.contractRiskMap[riskTx.Target]; ok {
		if rcfg.Kind == "erc20" {
			cnt, err := rule_Token(ctx, from, riskTx)
			if err != nil {
				g.Log().Warning(ctx, "checTx rule_Token:", riskTx, err)
				return consts.RiskCodeError, err
			}
			threshold := &big.Int{}
			//todo:
			threshold.UnmarshalText([]byte("1000000000000000000"))
			if cnt.Cmp(threshold) > 0 {
				g.Log().Warning(ctx, "checTx > threshold:", riskTx, cnt.String(), threshold.String())
				return consts.RiskCodeNeedVerification, nil
			}
			g.Log().Debug(ctx, "checTx < threshold:", riskTx, cnt.String(), threshold.String())
			return consts.RiskCodePass, nil
		} else if rcfg.Kind == "erc721" {
			if riskTx.MethodName != "tranferFrom" {
				return consts.RiskCodeNoRiskControl, nil
			}
			cnt, err := rule_nftcnt(ctx, from, riskTx)
			if err != nil {
				g.Log().Warning(ctx, "checTx rule_Token:", riskTx, err)
				return consts.RiskCodeError, err
			}
			if cnt > rcfg.Threshold {
				g.Log().Debug(ctx, "checTx > threshold:", riskTx, cnt, rcfg.Threshold)
				return consts.RiskCodeNeedVerification, nil
			}
			g.Log().Debug(ctx, "checTx < threshold:", riskTx, cnt, rcfg.Threshold)
			return consts.RiskCodePass, nil
		} else {
			g.Log().Warning(ctx, "checkTx unkonwo contract:", riskTx)
			// return consts.RiskCodeError, gerror.NewCode(gcode.CodeInvalidParameter)
			return consts.RiskCodePass, nil
		}
	}
	//notice: default
	return consts.RiskCodePass, nil
}

// 矿机、装备、时装、武器
func rule_nftcnt(ctx context.Context, address string, tx *analyzsigndata.AnalzyedTxData) (int, error) {
	// values := []string{}
	rst, err := dao.EthTx.Ctx(ctx).Where(do.EthTx{
		Address:    address,
		Target:     tx.Target,
		MethodName: tx.MethodName,
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

// MUD、MAK、USDT、RPG
func rule_Token(ctx context.Context, address string, tx *analyzsigndata.AnalzyedTxData) (*big.Int, error) {
	rst, err := dao.EthTx.Ctx(ctx).Where(do.EthTx{
		Address:    address,
		Target:     tx.Target,
		MethodName: tx.MethodName,
	}).
		WhereGT(dao.EthTx.Columns().CreatedAt, gtime.Now().Add(BeforH24)).Fields(
		dao.EthTx.Columns().Value,
	).
		Fields(dao.EthTx.Columns().Value).
		All()
	g.Log().Debug(ctx, "rule_Token:", rst, err,
		"tx.Target:", tx.Target,
		"methdoName:", tx.Target,
		"address:", address,
	)
	if err != nil {
		return big.NewInt(0), err
	}
	///
	val := big.NewInt(0)
	for _, v := range rst {
		c := &big.Int{}
		c.UnmarshalText(v[dao.EthTx.Columns().Value].Bytes())
		val = val.Add(val, c)
	}
	g.Log().Debug(ctx, "rule_Token:", rst, err,
		"tx.Target:", tx.Target,
		"methdoName:", tx.Target,
		"address:", address,
		"val:", val,
	)
	return val, nil

}

// var rulesFuncList map[string]func(ctx context.Context, data *conrisk.RiskTx) (int, error)

type contractRisk struct {
	Contract   string
	Kind       string
	MethodName string
	Threshold  int
}

// var contractRiskMap map[string]*contractRisk

// func init() {
// 	// rulesFuncList = make(map[string]func(ctx context.Context, data *conrisk.RiskTx) (int, error))
// 	// rulesFuncList["Token24HCnt"] = rule_Token24HCnt
// 	// rulesFuncList["Nft24HCnt"] = rule_NFTCnt
// 	////
// 	contractRiskMap = map[string]*contractRisk{}
// 	ctx := context.Background()
// 	riskcfg, err := gcfg.Instance().Get(ctx, "contractRisk")
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, val := range riskcfg.Array() {
// 		if valrisk, ok := val.(map[string]interface{}); !ok {
// 			panic(fmt.Errorf("contractRisk:%v", val))
// 		} else {
// 			Threshold, _ := valrisk["threshold"].(json.Number).Int64()
// 			r := &contractRisk{
// 				Contract:   strings.ToLower(valrisk["contract"].(string)),
// 				Kind:       valrisk["kind"].(string),
// 				MethodName: valrisk["methodName"].(string),
// 				Threshold:  int(Threshold),
// 			}
// 			contractRiskMap[r.Contract] = r
// 		}
// 	}
// 	////
// 	analzer = analyzsigndata.NewAnalzer()

// }
