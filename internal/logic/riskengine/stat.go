package riskengine

import (
	"context"
	v1 "riskcontrol/api/riskctrl/v1"
)

func (s *sRiskEngine) Stat(ctx context.Context, req *v1.StateReq) interface{} {
	res := map[string]interface{}{}
	for _, rule := range s.TxEnginePool {
		res[rule.EngineName()] = rule

	}
	return res
}
