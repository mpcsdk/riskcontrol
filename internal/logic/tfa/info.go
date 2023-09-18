package tfa

import (
	"context"
	"errors"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) TFAInfo(ctx context.Context, userId string) (*entity.Tfa, error) {
	///
	///
	rst := entity.Tfa{}
	err := dao.Tfa.Ctx(ctx).Where(dao.Tfa.Columns().UserId, userId).Scan(&rst)
	if err != nil {
		g.Log().Error(ctx, "tfainfo:", err, userId)
		return nil, gerror.NewCode(gcode.CodeOperationFailed)
	}
	return &rst, nil
}

func (s *sTFA) hasTFA(ctx context.Context, userId string) error {

	cnt, err := dao.Tfa.Ctx(ctx).Where(dao.Tfa.Columns().UserId, userId).CountColumn(dao.Tfa.Columns().UserId)
	if err != nil {
		return err
	}
	if cnt == 0 {
		//todo:
		return errors.New("")
	}
	return nil
}
