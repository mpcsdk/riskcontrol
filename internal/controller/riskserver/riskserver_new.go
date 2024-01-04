// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package riskserver

import (
	"riskcontral/api/riskserver"
	nats "riskcontral/internal/controller/nrpcserver"
)

type ControllerV1 struct{
	nrpc *nats.NrpcServer
}

func NewV1() riskserver.IRiskserverV1 {
	return &ControllerV1{
		nrpc : nats.Instance(),
	}
}