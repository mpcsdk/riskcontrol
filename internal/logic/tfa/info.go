package tfa

import (
	"context"
	"errors"
	"riskcontral/internal/consts"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) TFAInfo(ctx context.Context, userId string) (*entity.Tfa, error) {
	/// cache
	rst, err := s.getTfaCache(ctx, userId)
	if err != nil {
		g.Log().Error(ctx, "TFAInfo:", userId, err)
		return nil, err
	}
	///
	err = dao.Tfa.Ctx(ctx).Where(dao.Tfa.Columns().UserId, userId).Scan(&rst)
	if err != nil {
		g.Log().Error(ctx, "tfainfo:", err, userId)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	//set cache
	s.setTfaCache(ctx, userId, rst)
	return rst, nil
}

func (s *sTFA) getTfaCache(ctx context.Context, userId string) (*entity.Tfa, error) {
	if v, ok := service.Cache().Get(ctx, userId+consts.KEY_TFAInfoCache); ok == nil && !v.IsEmpty() {
		info := &entity.Tfa{}
		err := v.Struct(info)
		if err != nil {
			return nil, gerror.NewCode(consts.CodeInternalError)
		}
		return info, nil
	}
	return nil, gerror.NewCode(consts.CodeTFANotExist)
}

func (s *sTFA) setTfaCache(ctx context.Context, userId string, info *entity.Tfa) error {
	return service.Cache().Set(ctx, userId+consts.KEY_TFAInfoCache, info, consts.SessionDur)
}
func (s *sTFA) hasTFA(ctx context.Context, userId string) error {
	///cache
	_, err := s.getTfaCache(ctx, userId)
	if err == nil {
		return nil
	}
	cnt, err := dao.Tfa.Ctx(ctx).Where(dao.Tfa.Columns().UserId, userId).CountColumn(dao.Tfa.Columns().UserId)
	if err != nil {
		return err
	}
	if cnt == 0 {
		return errors.New("no tfa")
	}
	return nil
}
