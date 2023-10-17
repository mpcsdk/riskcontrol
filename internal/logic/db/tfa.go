package db

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

///
///

func (s *sDB) TfaMailExists(ctx context.Context, mail string) error {
	rst, err := dao.Tfa.Ctx(ctx).Where(do.Tfa{
		Mail: mail,
	}).CountColumn(dao.Tfa.Columns().Mail)
	if err != nil {
		return nil
	}
	if rst > 0 {
		return errEmpty
	}
	return nil
}
func (s *sDB) TfaPhoneExists(ctx context.Context, phone string) error {
	rst, err := dao.Tfa.Ctx(ctx).Where(do.Tfa{
		Phone: phone,
	}).CountColumn(dao.Tfa.Columns().Mail)
	if err != nil {
		return nil
	}
	if rst > 0 {
		return errEmpty
	}
	return nil
}
func (s *sDB) InsertTfaInfo(ctx context.Context, userId string, data *do.Tfa) error {

	cnt, err := g.Model(dao.Tfa.Table()).Ctx(ctx).Where(do.Tfa{
		UserId: data.UserId,
	}).CountColumn(dao.Tfa.Columns().UserId)
	if err != nil {
		return err
	}
	if cnt != 0 {
		return nil
	}

	_, err = g.Model(dao.Tfa.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.Tfa.Table() + userId,
		Force:    false,
	}).Data(data).
		Insert()

	return err
}

// //
func (s *sDB) UpdateTfaInfo(ctx context.Context, userId string, data *do.Tfa) error {
	_, err := g.Model(dao.Tfa.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.Tfa.Table() + userId,
		Force:    false,
	}).Data(data).Where(do.Tfa{
		UserId: data.UserId,
	}).Update()
	return err
}

func (s *sDB) FetchTfaInfo(ctx context.Context, userId string) (*entity.Tfa, error) {
	if userId == "" {
		return nil, nil
	}

	var data *entity.Tfa
	rst, err := g.Model(dao.Tfa.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     dao.Tfa.Table() + userId,
		Force:    false,
		// }).Where("user_id", 1).One()
	}).Where(do.Tfa{
		UserId: userId,
	}).One()
	if err != nil {
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	// if rst == nil {
	// 	g.Log().Warning(ctx, "FetchTfaInfo not exist:", userId)
	// 	return nil, gerror.NewCode(consts.CodeTFANotExist)
	// }
	err = rst.Struct(&data)
	return data, err
}
