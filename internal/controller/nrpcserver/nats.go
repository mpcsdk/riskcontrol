package nrpcserver

import (
	"context"
	"riskcontral/api/riskctrl"
	"riskcontral/internal/conf"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-io/nats.go"
)

type NrpcServer struct {
	ctx   context.Context
	sub   *nats.Subscription
	nc    *nats.Conn
	cache *gcache.Cache
}

var once sync.Once
var nrpcServer *NrpcServer

func Init() *NrpcServer {
	return Instance()
}
func Instance() *NrpcServer {
	once.Do(func() {
		nrpcServer = new()
	})
	return nrpcServer
}
func new() *NrpcServer {
	//
	nc, err := nats.Connect(conf.Config.Nats.NatsUrl, nats.Timeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	// defer nc.Close()
	redisCache := gcache.NewAdapterRedis(g.Redis())
	s := &NrpcServer{
		cache: gcache.New(),
	}
	s.cache.SetAdapter(redisCache)
	///
	h := riskctrl.NewRiskCtrlHandler(gctx.GetInitCtx(), nc, s)
	sub, err := nc.QueueSubscribe(h.Subject(), "riskcontrol", h.Handler)
	if err != nil {
		panic(err)
	}
	s.sub = sub
	s.nc = nc
	s.ctx = gctx.GetInitCtx()

	///
	return s
}

func (*NrpcServer) RpcAlive(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "RpcAlive")
	defer span.End()
	//
	return &empty.Empty{}, nil
}
