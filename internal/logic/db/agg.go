package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// /
// func (s *sDB) UpAggFT(ctx context.Context, ft *entity.AggFt24H) error {
// 	aggdo := &do.AggFt24H{
// 		From:       ft.From,
// 		Contract:   ft.Contract,
// 		MethodName: ft.MethodName,
// 	}
// 	dao.AggFt24H.Ctx(ctx).Delete(aggdo)
// 	///
// 	_, err := dao.AggFt24H.Ctx(ctx).Insert(ft)
// 	if err != nil {
// 		g.Log().Error(ctx, "UpAggFT:", "ft:", ft, "err:", err)
// 		return mpccode.CodeInternalError()
// 	}

// 	return nil
// }
// func (s *sDB) UpAggNFT(ctx context.Context, nft *entity.AggNft24H) error {
// 	aggdo := &do.AggNft24H{
// 		From:       nft.From,
// 		Contract:   nft.Contract,
// 		MethodName: nft.MethodName,
// 	}
// 	dao.AggNft24H.Ctx(ctx).Delete(aggdo)
// 	///
// 	_, err := dao.AggNft24H.Ctx(ctx).Insert(nft)
// 	if err != nil {
// 		g.Log().Error(ctx, "UpAggNFT:", "nft:", nft, "err:", err)
// 		return mpccode.CodeInternalError()
// 	}

// 	return nil
// }

func (s *sDB) GetAggFT(ctx context.Context, from, contract, methodName string) (*entity.AggFt24H, error) {
	aggdo := &do.AggFt24H{
		From:       from,
		Contract:   contract,
		MethodName: methodName,
	}
	rst, err := dao.AggFt24H.Ctx(ctx).
		Where(aggdo).
		One()
	if err != nil {
		g.Log().Error(ctx, "GetAggFT:", "from:", from,
			"contract:", contract, "methodName:", methodName, "err:", err)
		return nil, mpccode.CodeInternalError()
	}
	if rst.IsEmpty() {
		return nil, nil
	}
	////
	var data *entity.AggFt24H = nil
	err = rst.Struct(&data)
	if err != nil {
		g.Log().Error(ctx, "GetAggFT:", "from:", from,
			"contract:", contract, "methodName:", methodName, "data:", data, "err:", err)
		return nil, mpccode.CodeInternalError()
	}
	return data, nil
}

// /
func (s *sDB) GetAggNFT(ctx context.Context, from, contract, methodName string) (int, error) {
	cnt, err := dao.AggNft24H.Ctx(ctx).
		Where(do.AggNft24H{
			From:       from,
			Contract:   contract,
			MethodName: methodName,
		}).
		Count()
	if err != nil {
		g.Log().Error(ctx, "GetAggNFT:", "from:", from,
			"contract:", contract, "methodName:", methodName, "err:", err)
		return 0, mpccode.CodeInternalError()
	}
	if cnt == 0 {
		return 0, nil
	}
	////
	return cnt, nil
}
