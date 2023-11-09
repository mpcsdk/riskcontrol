package nats

import (
	"context"
	v1 "riskcontral/api/risk/nrpc/v1"
	"riskcontral/internal/config"
	"riskcontral/internal/service"
	"time"

	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-io/nats.go"
)

type sNats struct {
	sub *nats.Subscription
	nc  *nats.Conn
}

func init() {
	nc, err := nats.Connect(config.Config.Nrpc.NatsUrl, nats.Timeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	// defer nc.Close()

	s := &sNats{}
	h := v1.NewUserHandler(gctx.GetInitCtx(), nc, s)
	sub, err := nc.QueueSubscribe(h.Subject(), "riskcontrol", h.Handler)
	if err != nil {
		panic(err)
	}
	// defer sub.Unsubscribe()
	s.sub = sub
	s.nc = nc

	service.RegisterNats(s)
}

func (*sNats) PerformAlive(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "PerformAlive")
	defer span.End()
	//
	return &empty.Empty{}, nil
}
