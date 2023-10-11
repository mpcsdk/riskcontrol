package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// /
func (s *sDB) InsertTx(ctx context.Context, d *entity.EthTx) error {
	_, err := dao.EthTx.Ctx(ctx).FieldsEx(dao.EthTx.Columns().Id).Insert(d)
	if err != nil {
		g.Log().Warning(ctx, "RecordTxs :", err, d)
	}
	return nil
}

// /
func (s *sDB) InsertTxs(ctx context.Context, txs []*entity.EthTx) error {
	_, err := dao.EthTx.Ctx(ctx).FieldsEx(dao.EthTx.Columns().Id).Insert(txs)
	if err != nil {
		g.Log().Warning(ctx, "RecordTxs :", err, txs)
	}
	return nil
}
