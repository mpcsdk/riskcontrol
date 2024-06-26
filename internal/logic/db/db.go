package db

import (
	"context"
	"riskcontrol/internal/conf"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

type sDB struct {
	ctx          context.Context
	chainCfg     *mpcdao.ChainCfg
	riskCtrlRule *mpcdao.RiskCtrlRule
	riskCtrlTfa  *mpcdao.RiskTfa
}

func New() *sDB {
	///
	r := g.Redis()
	_, err := r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	///
	s := &sDB{
		ctx:          gctx.GetInitCtx(),
		chainCfg:     mpcdao.NewChainCfg(r, conf.Config.Cache.Duration),
		riskCtrlRule: mpcdao.NewRiskCtrlRule(r, conf.Config.Cache.Duration),
		riskCtrlTfa:  mpcdao.NewRiskTfa(r, conf.Config.Cache.Duration),
	}

	return s
}

func (s *sDB) TfaDB() *mpcdao.RiskTfa {
	return s.riskCtrlTfa
}
func (s *sDB) RiskCtrl() *mpcdao.RiskCtrlRule {
	return s.riskCtrlRule
}
func (s *sDB) ChainCfg() *mpcdao.ChainCfg {
	return s.chainCfg
}
