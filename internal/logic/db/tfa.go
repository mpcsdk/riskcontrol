package db

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

type sDB struct {
	cache *gcache.Cache
}

func new() *sDB {
	return &sDB{
		cache: gcache.New(),
	}
}

// 初始化
func init() {
	service.RegisterDB(new())
	redisCache := gcache.NewAdapterRedis(g.Redis())
	g.DB().GetCache().SetAdapter(redisCache)
}

///
///

func (s *sDB) InsertTfaInfo(ctx context.Context, data *entity.Tfa) error {

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
		Name:     dao.Tfa.Table() + data.UserId,
		Force:    false,
	}).Data(data).
		Insert()

	return err
}

// //
func (s *sDB) UpdateTfaInfo(ctx context.Context, data *entity.Tfa) error {
	_, err := g.Model(dao.Tfa.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.Tfa.Table() + data.UserId,
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

	err = rst.Struct(&data)
	return data, err
}
