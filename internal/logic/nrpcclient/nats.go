package nrpcclient

import (
	"riskcontrol/api/riskengine"

	"github.com/nats-io/nats.go"
)

type sNrpcClient struct {
	nc         *nats.Conn
	riskengine *riskengine.RiskEngineClient
}

// func init() {

// 	// Connect to the NATS server.
// 	nc, err := nats.Connect(conf.Config.Nats.NatsUrl, nats.Timeout(3*time.Second))
// 	if err != nil {
// 		panic(err)
// 	}
// 	// defer nc.Close()

//		riskengine := riskengine.NewRiskEngineClient(nc)
//		_, err = riskengine.RpcAlive(&empty.Empty{})
//		if err != nil {
//			panic(err)
//		}
//		///
//		s := &sNrpcClient{
//			riskengine: riskengine,
//			nc:         nc,
//		}
//		service.RegisterNrpcClient(s)
//	}
func (s *sNrpcClient) Flush() {
	err := s.nc.Flush()
	if err != nil {
		panic(err)
	}
	s.riskengine = riskengine.NewRiskEngineClient(s.nc)
	// s.cli = tfav1.NewTFAClient(s.nc)
}
