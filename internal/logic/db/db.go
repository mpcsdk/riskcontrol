package db

import (
	"context"
	"riskcontral/internal/conf"
	"riskcontral/internal/service"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

type sDB struct {
	// cache *gcache.Cache
	ctx        context.Context
	dbDuration time.Duration
	///
	riskCtrlTfa *mpcdao.RiskTfa
}

func new() *sDB {
	///
	r := g.Redis("")
	_, err := r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	///
	s := &sDB{
		ctx:         gctx.GetInitCtx(),
		dbDuration:  time.Duration(conf.Config.Cache.DBCacheDuration) * time.Second,
		riskCtrlTfa: mpcdao.NewRiskTfa(r, conf.Config.Cache.DBCacheDuration),
	}

	return s
}

func (s *sDB) TfaDB() *mpcdao.RiskTfa {
	return s.riskCtrlTfa
}

// 初始化
func init() {
	service.RegisterDB(new())
}
