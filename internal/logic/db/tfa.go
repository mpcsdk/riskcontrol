package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

///
///

func (s *sDB) TfaMailNotExists(ctx context.Context, mail string) (bool, error) {
	rst, err := dao.Tfa.Ctx(ctx).Where(do.Tfa{
		Mail: mail,
	}).Count()
	if err != nil {
		g.Log().Error(ctx, "TfaMailNotExists:", "mail", mail, "err", err)
		return false, mpccode.CodeInternalError()
	}
	if rst > 0 {
		return false, nil
	}
	return true, nil
}
func (s *sDB) TfaPhoneNotExists(ctx context.Context, phone string) (bool, error) {
	rst, err := dao.Tfa.Ctx(ctx).Where(do.Tfa{
		Phone: phone,
	}).CountColumn(dao.Tfa.Columns().Phone)
	if err != nil {
		g.Log().Error(ctx, "TfaPhoneNotExists:", "phone", phone, "err", err)
		return false, mpccode.CodeInternalError()
	}
	if rst > 0 {
		return false, nil
	}
	return true, nil
}
func (s *sDB) InsertTfaInfo(ctx context.Context, userId string, data *do.Tfa) error {
	cnt, err := dao.Tfa.Ctx(ctx).Where(do.Tfa{
		UserId: data.UserId,
	}).CountColumn(dao.Tfa.Columns().UserId)

	if err != nil {
		g.Log().Error(ctx, "InsertTfaInfo:", "userId", userId, "data:", data, "err", err)
		return mpccode.CodeInternalError()
	}
	if cnt != 0 {
		return nil
	}

	_, err = dao.Tfa.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.Tfa.Table() + userId,
		Force:    false,
	}).Data(data).Insert()
	if err != nil {
		g.Log().Error(ctx, "InsertTfaInfo:", "userId", userId, "data:", data, "err", err)
		return mpccode.CodeInternalError()
	}

	return nil
}

// //
func (s *sDB) UpdateTfaInfo(ctx context.Context, userId string, data *do.Tfa) error {
	_, err := dao.Tfa.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.Tfa.Table() + userId,
		Force:    false,
	}).Data(data).Where(do.Tfa{
		UserId: data.UserId,
	}).Update()
	if err != nil {
		g.Log().Error(ctx, "UpdateTfaInfo:", "userId", userId, "data:", data, "err", err)
		return mpccode.CodeInternalError()
	}
	return nil
}

func (s *sDB) ExistsTfaInfo(ctx context.Context, userId string) (bool, error) {
	if userId == "" {
		return false, mpccode.CodeParamInvalid()
	}
	cnt, err := dao.Tfa.Ctx(ctx).Where(do.Tfa{
		UserId: userId,
	}).CountColumn(dao.Tfa.Columns().UserId)

	if err != nil {
		g.Log().Error(ctx, "ExistsTfaInfo:", "userId", userId, "err", err)
		return false, mpccode.CodeInternalError()
	}
	if cnt != 0 {
		return true, nil
	}
	return false, nil
}

func (s *sDB) FetchTfaInfo(ctx context.Context, userId string) (*entity.Tfa, error) {
	if userId == "" {
		return nil, mpccode.CodeParamInvalid()
	}

	aggdo := &do.Tfa{
		UserId: userId,
	}
	var data *entity.Tfa
	///
	rst, err := dao.Tfa.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dbDuration,
		Name:     dao.Tfa.Table() + userId,
		Force:    false,
	}).Where(aggdo).One()
	if err != nil {
		g.Log().Error(ctx, "ExistsTfaInfo:", "userId", userId, "agg:", aggdo, "err", err)
		return nil, mpccode.CodeInternalError()
	}
	if rst.IsEmpty() {
		return nil, nil
	}
	err = rst.Struct(&data)
	if err != nil {
		return nil, mpccode.CodeInternalError()
	}

	return data, nil
}
