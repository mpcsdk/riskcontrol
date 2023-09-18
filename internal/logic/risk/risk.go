package risk

import (
	"context"
	"riskcontral/common"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/service"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

type sRisk struct {
}

func (s *sRisk) PerformRiskTxs(ctx context.Context, userId string, address string, txs []*conrisk.RiskTx) (string, int32, error) {
	//
	//todo: record riskinfo
	riskserial := common.GenNewSid()
	///
	code, err := s.checkTxs(ctx, address, txs)
	if err != nil {
		return riskserial, -1, err
	}
	//todo: record serial
	service.Cache().Set(ctx, riskserial+"riskUserId", userId, 0)
	return riskserial, code, err
}
func (s *sRisk) PerformRiskTFA(ctx context.Context, userId string, riskData *conrisk.RiskTfa) (string, int32, error) {

	//todo: record riskinfo
	riskserial := common.GenNewSid()
	///
	var code int32 = -1
	var err error
	///
	switch riskData.Kind {
	case "upPhone":
		code, err = s.checkTFAUpPhone(ctx, userId)
	case "upMail":
		code, err = s.checkTfaUpMail(ctx, userId)
	default:
		return riskserial, -1, nil
	}
	if err != nil {
		return riskserial, code, err
	}
	// todo: record serial
	service.Cache().Set(ctx, riskserial+"riskUserId", userId, 0)
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
