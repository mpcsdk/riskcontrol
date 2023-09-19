package tfa

import (
	"context"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/os/gctx"
)

type sTFA struct {
	ctx context.Context
	// riskClient riskv1.UserClient
	pendding     map[string]func()
	sendpendding map[string]func()
	url          string
	////
}

func new() *sTFA {

	ctx := gctx.GetInitCtx()
	// addr, err := gcfg.Instance().Get(ctx, "etcd.address")
	// if err != nil {
	// 	panic(err)
	// }
	// grpcx.Resolver.Register(etcd.New(addr.String()))
	// conn, err := grpcx.Client.NewGrpcClientConn("rulerpc")
	// if err != nil {
	// 	panic(err)
	// }
	// client := risk.NewUserClient(conn)
	///
	//
	s := &sTFA{
		pendding:     map[string]func(){},
		sendpendding: map[string]func(){},
		ctx:          ctx,
	}
	///

	return s
}

func init() {
	service.RegisterTFA(new())
}
