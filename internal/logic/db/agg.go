package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"
)

// /
func (s *sDB) UpAggFT(ctx context.Context, ft *entity.AggFt24H) error {
	dao.AggFt24H.Ctx(ctx).Delete(do.AggFt24H{
		From:       ft.From,
		Contract:   ft.Contract,
		MethodName: ft.MethodName,
	})
	///
	_, err := dao.AggFt24H.Ctx(ctx).Insert(ft)
	// _, err := dao.EthTx.Ctx(ctx).Insert(txs)
	// if err != nil {
	// 	g.Log().Warning(ctx, "RecordTxs :", err, txs)
	// }
	return err
}
func (s *sDB) UpAggNFT(ctx context.Context, nft *entity.AggNft24H) error {
	dao.AggNft24H.Ctx(ctx).Delete(do.AggNft24H{
		From:       nft.From,
		Contract:   nft.Contract,
		MethodName: nft.MethodName,
	})
	///
	_, err := dao.AggNft24H.Ctx(ctx).Insert(nft)
	// _, err := dao.EthTx.Ctx(ctx).Insert(txs)
	// if err != nil {
	// 	g.Log().Warning(ctx, "RecordTxs :", err, txs)
	// }
	return err
}

func (s *sDB) GetAggFT(ctx context.Context, from, contract, methodName string) (*entity.AggFt24H, error) {
	rst, err := dao.AggFt24H.Ctx(ctx).
		Where(do.AggFt24H{
			From:       from,
			Contract:   contract,
			MethodName: methodName,
		}).
		One()
	if err != nil {
		return nil, err
	}
	if rst.IsEmpty() {
		return nil, errEmpty
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
