package nrpcclient

import (
	tfav1 "riskcontral/api/tfa/nrpc/v1"
	"riskcontral/internal/config"
	"riskcontral/internal/service"
	"time"

	"github.com/nats-io/nats.go"
)

type sNrpcClient struct {
	cli *tfav1.TFAClient
	nc  *nats.Conn
}

func init() {

	// Connect to the NATS server.
	nc, err := nats.Connect(config.Config.Nrpc.NatsUrl, nats.Timeout(3*time.Second))
	if err != nil {
		panic(err)
	}
	// defer nc.Close()

	// This is our generated client.
	cli := tfav1.NewTFAClient(nc)

	// Contact the server and print out its response.
	// _, err = cli.RpcAlive(&empty.Empty{})
	// if err != nil {
	// 	panic(err)
	// }
	s := &sNrpcClient{
		cli: cli,
		nc:  nc,
	}
	service.RegisterNrpcClient(s)
}
func (s *sNrpcClient) Flush() {
	err := s.nc.Flush()
	if err != nil {
		panic(err)
	}
	s.cli = tfav1.NewTFAClient(s.nc)
}
