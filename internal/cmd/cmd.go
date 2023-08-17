package cmd

import (
	"context"

	"riskcontral/internal/controller/rules"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			///
			// l, err := net.Listen("tcp", ":50051")
			// if err != nil {
			// 	log.Panic(err)
			// }
			// m := cmux.New(l)
			// // We first match the connection against HTTP2 fields. If matched, the
			// // connection will be sent through the "grpcl" listener.
			// grpcl := m.Match(cmux.HTTP2HeaderFieldPrefix("content-type", "application/grpc"))
			// //Otherwise, we match it againts a websocket upgrade request.
			// // wsl := m.Match(cmux.HTTP1HeaderField("Upgrade", "websocket"))
			// // Otherwise, we match it againts HTTP1 methods. If matched,
			// // it is sent through the "httpl" listener.
			// httpl := m.Match(cmux.HTTP1Fast())
			////grpc
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			rpcs := grpcx.Server.New(c)
			// rpcs.Server.Serve(grpcl)
			rules.Register(rpcs)
			rpcs.Start()
			/// http
			s := g.Server()
			// s.SetListener(httpl)
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
