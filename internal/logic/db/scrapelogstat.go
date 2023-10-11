package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sDB) GetScrapeStat(ctx context.Context, chainId string) (*entity.ScrapeLogsStat, error) {
	rst, err := g.Model(dao.ScrapeLogsStat.Table()).Ctx(ctx).Where(do.ScrapeLogsStat{
		ChainId: chainId,
	}).One()
	if err != nil {
		return nil, err
	}

	if rst.IsEmpty() {
		s.insert(ctx, chainId, rpc.LatestBlockNumber)
		return &entity.ScrapeLogsStat{
			ChainId:   chainId,
			LastBlock: rpc.LatestBlockNumber.Int64(),
		}, nil
	}
	var data *entity.ScrapeLogsStat = nil
	err = rst.Struct(&data)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errEmpty
	}
	return data, nil
}

func (s *sDB) insert(ctx context.Context, chainId string, bn rpc.BlockNumber) error {
	_, err := dao.ScrapeLogsStat.Ctx(ctx).Insert(&entity.ScrapeLogsStat{
		ChainId:   chainId,
		LastBlock: bn.Int64(),
		UpdatedAt: gtime.Now(),
	})
	if err != nil {
		g.Log().Warning(ctx, "ScrapeLogsStatinsert :", err)
	}
	return err
}
func (s *sDB) SetScrapeStat(ctx context.Context, chainId string, nr rpc.BlockNumber) error {

	_, err := g.Model(dao.ScrapeLogsStat.Table()).Ctx(ctx).
		Data(do.ScrapeLogsStat{
			LastBlock: nr.Int64(),
			UpdatedAt: gtime.Now(),
		}).Where(do.ScrapeLogsStat{
		ChainId: chainId,
	}).Update()
	return err
}
