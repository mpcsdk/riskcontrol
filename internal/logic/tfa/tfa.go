package tfa

import (
	"context"
	"riskcontral/internal/conf"
	"riskcontral/internal/model"
	"riskcontral/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type UserRiskId string

func keyUserRiskId(userId string, riskSerial string) UserRiskId {
	return UserRiskId(userId + "keyUserRiskId" + riskSerial)
}

type sTFA struct {
	forbiddentTime        time.Duration
	ctx                   context.Context
	riskPenddingContainer *model.RiskPenddingContainer
	////
	cache *gcache.Cache
}

// /
func new() *sTFA {

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
		//todo:
		riskPenddingContainer: model.NewRiskPenddingContainer(t),
		ctx:                   ctx,
		cache:                 gcache.New(),
		forbiddentTime:        forbiddentTime,
	}
	redisCache := gcache.NewAdapterRedis(g.Redis())
	s.cache.SetAdapter(redisCache)
	///

	return s
}

///

func init() {
	service.RegisterTFA(new())
}

func (s *sTFA) TfaRiskKind(ctx context.Context, tfaInfo *entity.Tfa, riskSerial string) (model.RiskKind, error) {
	risk := s.riskPenddingContainer.GetRiskVerify(tfaInfo.UserId, riskSerial)
	if risk == nil {
		g.Log().Warning(ctx, "TfaRiskKind:", "tfaInfo:", tfaInfo, "riskSerial:", riskSerial)
		return "", mpccode.CodeParamInvalid()
	}
	return risk.RiskKind, nil
}
