package risk

import (
	"context"
	"riskcontral/internal/service"
)

type sRisk struct{}

func (s *sRisk) PerformRisk(ctx context.Context, riskName string, riskData interface{}) (interface{}, error) {
	return nil, nil
}

// new 创建一个新的sRisk
func new() *sRisk {
	return &sRisk{}
}
func init() {
	service.RegisterRisk(new())
}
