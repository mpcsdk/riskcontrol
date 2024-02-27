package nrpcclient

import (
	"riskcontral/api/riskengine"
	"riskcontral/internal/config"
	"riskcontral/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-io/nats.go"
)

type sNrpcClient struct {
	nc         *nats.Conn
	riskengine *riskengine.RiskEngineClient
}

func init() {

	// Connect to the NATS server.
	nc, err := nats.Connect(config.Config.Nrpc.NatsUrl, nats.Timeout(3*time.Second))
	if err != nil {
		panic(err)
	}
	// defer nc.Close()

	// This is our generated client.
	// cli := tfav1.NewTFAClient(nc)

	// Contact the server and print out its response.
	// _, err = cli.RpcAlive(&empty.Empty{})
	// if err != nil {
	// 	panic(err)
	// }
	ctx := gctx.GetInitCtx()
	riskengine := riskengine.NewRiskEngineClient(nc)
	_, err = riskengine.RpcAlive(&empty.Empty{})
	if err != nil {
		g.Log().Error(ctx, err)
		panic(err)
	}
	///
	s := &sNrpcClient{
		// cli: cli,
		riskengine: riskengine,
		nc:         nc,
	}
	service.RegisterNrpcClient(s)
}
func (s *sNrpcClient) Flush() {
	err := s.nc.Flush()
	if err != nil {
		panic(err)
	}
	s.riskengine = riskengine.NewRiskEngineClient(s.nc)
	// s.cli = tfav1.NewTFAClient(s.nc)
}
