// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package riskctrl

import (
	"riskcontrol/api/riskctrl"
	"riskcontrol/internal/controller/limiter"
)

type ControllerV1 struct{
	limiter *limiter.Limiter	
}

func NewV1() riskctrl.IRiskctrlV1 {
	return &ControllerV1{
		limiter: limiter.Instance(),
	}
}

