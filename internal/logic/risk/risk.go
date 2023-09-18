package risk

import (
	"context"
	"riskcontral/common"
	"riskcontral/internal/consts"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/service"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type sRisk struct {
}

func (s *sRisk) PerformRiskTxs(ctx context.Context, userId string, address string, txs []*conrisk.RiskTx) (string, int32, error) {
	//
	g.Log().Debug(ctx, "PerformRiskTxs:", "userId:", userId, "address:", address, "txs:", txs)
	///
	riskserial := common.GenNewSid()
	//todo:
	return riskserial, 0, nil
	///
	code, err := s.checkTxs(ctx, address, txs)
	if err != nil {
		g.Log().Warning(ctx, "PerformRiskTxs:", "checkTxs:", err)
		return riskserial, -1, gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	//
	service.Cache().Set(ctx, riskserial+consts.KEY_RiskUId, userId, 0)
	return riskserial, code, err
}
func (s *sRisk) PerformRiskTFA(ctx context.Context, userId string, riskData *conrisk.RiskTfa) (string, int32, error) {
	g.Log().Debug(ctx, "PerformRiskTFA:", "userId:", userId, "riskData:", riskData)
	//
	riskserial := common.GenNewSid()
	//todo:
	return riskserial, 0, nil
	///
	var code int32 = -1
	var err error
	///
	switch riskData.Kind {
	case consts.KEY_TFAKindUpPhone:
		code, err = s.checkTFAUpPhone(ctx, userId)
	case consts.KEY_TFAKindUpMail:
		code, err = s.checkTfaUpMail(ctx, userId)
	case consts.KEY_TFAKindCreate:
		code, err = s.checkTfaCreate(ctx, userId)
	default:
		g.Log().Error(ctx, "PerformRiskTFA:", "kind:", riskData.Kind, "not support")
		return riskserial, -1, gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	if err != nil {
		g.Log().Error(ctx, "PerformRiskTFA:", err)
		return riskserial, code, gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	///
	service.Cache().Set(ctx, riskserial+consts.KEY_RiskUId, userId, 0)
	return riskserial, code, nil
}

func new() *sRisk {
	return &sRisk{}
}

var BeforH24 time.Duration
var BeforM1 time.Duration

func init() {
	var err error
	BeforH24, err = gtime.ParseDuration("-24h")
	if err != nil {
		panic(err)
	}

	BeforM1, err = gtime.ParseDuration("-1m")
	if err != nil {
		panic(err)
	}

	///
	service.RegisterRisk(new())
}
