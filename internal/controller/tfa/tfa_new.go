// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package tfa

import (
	"riskcontral/api/tfa"
	"riskcontral/internal/controller/limiter"
)




type ControllerV1 struct{
	limiter *limiter.Limiter	
}

func NewV1() tfa.ITfaV1 {
	s := &ControllerV1{
		limiter: limiter.Instance(),
	}
	return s
}

