// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package tfa

import (
	"riskcontral/api/tfa"
	nats "riskcontral/internal/controller/nrpcserver"
)

type ControllerV1 struct{
	nrpc *nats.NrpcServer
}

func NewV1() tfa.ITfaV1 {
	return &ControllerV1{
		nrpc : nats.Instance(),
	}
}

