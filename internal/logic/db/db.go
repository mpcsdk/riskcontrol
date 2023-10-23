package db

import (
	"context"
	"errors"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

type sDB struct {
	// cache *gcache.Cache
	ctx context.Context
}

func new() *sDB {
	// return &sDB{
	// 	cache: gcache.New(),
	// }
	// g.Redis().Exists(gctx.GetInitCtx())
	s := &sDB{
		ctx: gctx.GetInitCtx(),
	}
	//todo: notify
	go s.listenNotify([]string{RuleChName, AbiChName})
	return s
}

var errArg error = errors.New("arg err")
var errEmpty error = errors.New("empty db")
var errDataExists error = errors.New("empty data exists")

// 初始化
func init() {
	_, err := g.Redis().Del(gctx.GetInitCtx(), "test")
	if err != nil {
		panic(err)
	}
	///
	service.RegisterDB(new())
	redisCache := gcache.NewAdapterRedis(g.Redis())
	g.DB().GetCache().SetAdapter(redisCache)
}
