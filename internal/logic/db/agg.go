package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// /
func (s *sDB) UpAggFT(ctx context.Context, ft *entity.AggFt24H) error {
	aggdo := &do.AggFt24H{
		From:       ft.From,
		Contract:   ft.Contract,
		MethodName: ft.MethodName,
	}
	dao.AggFt24H.Ctx(ctx).Delete(aggdo)
	///
	_, err := dao.AggFt24H.Ctx(ctx).Insert(ft)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("do", aggdo),
		))
		return err
	}

	return nil
}
func (s *sDB) UpAggNFT(ctx context.Context, nft *entity.AggNft24H) error {
	aggdo := &do.AggNft24H{
		From:       nft.From,
		Contract:   nft.Contract,
		MethodName: nft.MethodName,
	}
	dao.AggNft24H.Ctx(ctx).Delete(aggdo)
	///
	_, err := dao.AggNft24H.Ctx(ctx).Insert(nft)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("do", aggdo),
		))
		return err
	}

	return nil
}

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
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("do", aggdo),
		))
		return nil, err
	}
	if rst.IsEmpty() {
		err = gerror.Wrap(errEmpty, mpccode.ErrDetails(
			mpccode.ErrDetail("do", aggdo),
		))
		return nil, err
	}
	////
	var data *entity.AggFt24H = nil
	err = rst.Struct(&data)
	return data, err
}
func (s *sDB) GetAggNFT(ctx context.Context, from, contract, methodName string) (int, error) {
	cnt, err := dao.AggNft24H.Ctx(ctx).
		Where(do.AggNft24H{
			From:       from,
			Contract:   contract,
			MethodName: methodName,
		}).
		Count()
	if err != nil {
		return 0, err
	}
	if cnt == 0 {
		return 0, errEmpty
	}
	////
	return cnt, err
}
