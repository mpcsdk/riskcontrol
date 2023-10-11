package risk

import (
	"context"
	"math/big"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
)

// 矿机、装备、时装、武器
func rule_nftcnt(ctx context.Context, address string, contract string, method string) (int, error) {
	cnt, err := dao.AggNft24H.Ctx(ctx).Where(do.AggNft24H{
		From:       address,
		Contract:   contract,
		MethodName: method,
	}).
		Where(do.AggNft24H{
			From:       address,
			Contract:   contract,
			MethodName: method,
		}).Count()
	g.Log().Debug(ctx, "AggNft24H:", address, contract, method, cnt)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

// MUD、MAK、USDT、RPG
func rule_ftcnt(ctx context.Context, address string, contract string, method string) (*big.Int, error) {
	rst, err := dao.AggFt24H.Ctx(ctx).Where(do.AggFt24H{
		From:       address,
		Contract:   contract,
		MethodName: method,
	}).
		Where(do.AggFt24H{
			From:       address,
			Contract:   contract,
			MethodName: method,
		}).
		Fields(
			dao.AggFt24H.Columns().Value,
		).One()
	g.Log().Debug(ctx, "AggFt24H:", address, contract, method, rst)
	if err != nil {
		return big.NewInt(0), err
	}
	if rst == nil {
		return big.NewInt(0), nil
	}
	///

	val := big.NewInt(0)
	err = rst.Struct(val)
	return val, err

}