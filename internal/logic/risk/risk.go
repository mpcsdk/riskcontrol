package risk

import (
	"context"
	"riskcontral/internal/service"
	"time"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type sRisk struct{}

func (s *sRisk) PerformRisk(ctx context.Context, riskName string, riskData interface{}) (bool, error) {
	switch riskName {
	case "upPhone":
		return false, gerror.NewCode(gcode.CodeNotImplemented)
	case "upMail":
		return s.checkTfaUpMail(ctx, "upMail", riskData)
	case "checkTx":
		return s.checkTx(ctx, "checkTx", riskData)
	}
	return false, gerror.NewCode(gcode.CodeNotImplemented)
}

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
