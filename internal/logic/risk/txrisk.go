package risk

import (
	"context"
	"fmt"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sRisk) checkTx(ctx context.Context, riskName string, riskData interface{}) (bool, error) {
	if _, ok := riskData.(*conrisk.RiskTx); !ok {
		return false, gerror.NewCode(gcode.CodeInvalidParameter)
	}
	data := riskData.(*conrisk.RiskTx)
	////
	checkRst := false
	for _, f := range rulesList {
		switch data.Contract {
		case "0x60e4d786628fea6478f785a6d7e704777c86a7c6":
			rst, err := f(ctx, data)
			if err != nil {
				return false, err
			}
			if rst > data.Threshold {
				return false, nil
			}
			checkRst = true
		case "MUD、MAK、USDT、RPG":
		case "矿机、装备、时装、武器":
		}
	}

	return checkRst, nil
}

func rule_NFTCnt(ctx context.Context, data *conrisk.RiskTx) (int, error) {
	switch data.Contract {
	case "矿机":
		return rule_nftcnt(ctx, "usdt", "transfer", data)
	case "装备":
		return rule_nftcnt(ctx, "MUD", "transfer", data)
	case "时装":
		return rule_nftcnt(ctx, "MAK", "transfer", data)
	case "武器":
		return rule_nftcnt(ctx, "RPG", "safeTransferFrom", data)
	}
	return 0, gerror.New("rule_Token24HCnt")
}

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

func rule_Token24HCnt(ctx context.Context, data *conrisk.RiskTx) (int, error) {
	switch data.Contract {
	case "USDT":
		return rule_Token(ctx, "usdt", "transfer", data)
	case "MUD":
		return rule_Token(ctx, "MUD", "transfer", data)
	case "MAK":
		return rule_Token(ctx, "MAK", "transfer", data)
	case "RPG":
		return rule_Token(ctx, "RPG", "safeTransferFrom", data)
	}
	return 0, gerror.New("rule_Token24HCnt")
}

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

var rulesList map[string]func(ctx context.Context, data *conrisk.RiskTx) (int, error)

func init() {
	rulesList = make(map[string]func(ctx context.Context, data *conrisk.RiskTx) (int, error))
	rulesList["Token24HCnt"] = rule_Token24HCnt
	rulesList["Nft24HCnt"] = rule_NFTCnt
}
