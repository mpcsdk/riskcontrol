package nats

import (
	"context"
	"riskcontral/api/risk/nrpc"
	"riskcontral/internal/config"
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

var apiInterval = time.Second * 1
var limitSendInterval = time.Second * 60
var limitSendPhoneDurationCnt = 50
var limitSendPhoneDuration = time.Hour
var limitSendMailDurationCnt = 10
var limitSendMailDuration = time.Hour

var once sync.Once
var nrpcServer *NrpcServer

func Instance() *NrpcServer {
	once.Do(func() {
		nrpcServer = new()
	})
	return nrpcServer
}
func new() *NrpcServer {
	apiInterval = time.Duration(config.Config.Cache.ApiInterval) * time.Second
	limitSendInterval = time.Duration(config.Config.Cache.LimitSendInterval) * time.Second
	//
	limitSendPhoneDurationCnt = config.Config.Cache.LimitSendPhoneCount
	limitSendPhoneDuration = time.Duration(config.Config.Cache.LimitSendPhoneDuration) * time.Second
	limitSendMailDurationCnt = config.Config.Cache.LimitSendMailCount
	limitSendMailDuration = time.Duration(config.Config.Cache.LimitSendMailDuration) * time.Second

	//
	nc, err := nats.Connect(config.Config.Nrpc.NatsUrl, nats.Timeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	// defer nc.Close()
	// defer nc.Close()
	redisCache := gcache.NewAdapterRedis(g.Redis())
	s := &NrpcServer{
		cache: gcache.New(),
	}
	s.cache.SetAdapter(redisCache)
	///
	h := nrpc.NewRiskHandler(gctx.GetInitCtx(), nc, s)
	sub, err := nc.QueueSubscribe(h.Subject(), "riskcontrol", h.Handler)
	if err != nil {
		panic(err)
	}
	// defer sub.Unsubscribe()
	s.sub = sub
	s.nc = nc
	s.ctx = gctx.GetInitCtx()

	///
	s.NatsPub()
	// service.RegisterNrpcServer(s)
	return s
}

func (*NrpcServer) RpcAlive(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "RpcAlive")
	defer span.End()
	//
	return &empty.Empty{}, nil
}
