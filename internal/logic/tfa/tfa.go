package tfa

import (
	"context"
	"riskcontrol/internal/conf"
	check "riskcontrol/internal/logic/tfa/checker"
	pendding "riskcontrol/internal/logic/tfa/penddingrisk"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

type UserRiskId string

func keyUserRiskId(userId string, riskSerial string) UserRiskId {
	return UserRiskId(userId + "keyUserRiskId" + riskSerial)
}

type sTFA struct {
	forbiddentTime        time.Duration
	ctx                   context.Context
	riskPenddingContainer *pendding.RiskPenddingContainer
	////
	cache *gcache.Cache
	///
	checker *check.Checker
}

// /
func New() *sTFA {

	ctx := gctx.GetInitCtx()
	limitSendPhoneDurationCnt = conf.Config.Cache.LimitSendPhoneCount
	limitSendPhoneDuration = time.Duration(conf.Config.Cache.LimitSendPhoneDuration) * time.Second
	limitSendMailDurationCnt = conf.Config.Cache.LimitSendMailCount
	limitSendMailDuration = time.Duration(conf.Config.Cache.LimitSendMailDuration) * time.Second
	///
	forbiddentTime, err := gtime.ParseDuration(conf.Config.UserRisk.ForbiddenTime)
	if err != nil {
		panic(err)
	}
	//
	t := conf.Config.Cache.VerificationCodeDuration
	s := &sTFA{
		riskPenddingContainer: pendding.NewRiskPenddingContainer(t),
		ctx:                   ctx,
		cache:                 gcache.New(),
		forbiddentTime:        forbiddentTime,
		checker:               check.NewChecker(forbiddentTime),
	}
	redisCache := gcache.NewAdapterRedis(g.Redis())
	s.cache.SetAdapter(redisCache)
	///

	return s
}

///
