package risk

import (
	"context"
	"riskcontral/internal/service"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
)

type sRisk struct{}

func (s *sRisk) PerformRisk(ctx context.Context, riskName string, riskData interface{}) (interface{}, error) {
	if riskName == "phone" {
		return nil, nil
	} else if riskName == "mail" {
		return nil, gerror.NewCode(gcode.CodeNotImplemented)
	}

	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

// new 创建一个新的sRisk
func new() *sRisk {
	return &sRisk{}
}
func init() {
	service.RegisterRisk(new())
}
