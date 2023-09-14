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
func (s *sRisk) PerformRiskTFA(ctx context.Context, userId string, riskData *conrisk.RiskTfa) (string, error) {

	//todo: record riskinfo
	riskserial := common.GenNewSid()
	///
	//todo: record serial
	return riskserial, nil
}

// 	switch riskName {
// 	case "upPhone":
// 		return false, gerror.NewCode(gcode.CodeNotImplemented)
// 	case "upMail":
// 		return s.checkTfaUpMail(ctx, "upMail", riskData)
// 	case "checkTx":
// 		return s.checkTx(ctx, "checkTx", riskData)
// 	}
// 	return false, gerror.NewCode(gcode.CodeNotImplemented)
// }

// new 创建一个新的sRisk
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
