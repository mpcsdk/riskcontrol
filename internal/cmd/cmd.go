package cmd

import (
	"context"

	"riskcontral/internal/controller/rules"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			///
			//todo: etcd
			addr, err := gcfg.Instance().Get(ctx, "etcd.address")
			if err != nil {
				panic(err)
			}
			////grpc
			grpcx.Resolver.Register(etcd.New(addr.String()))
			// c := grpcx.Server.NewConfig()
			rpcs := grpcx.Server.New()
			rules.Register(rpcs)
			rpcs.Start()
			// // http
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					rules.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
