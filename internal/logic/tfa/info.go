package tfa

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type sTFA struct {
	ctx context.Context
	// riskClient riskv1.UserClient
	pendding map[string]func()
}

func (s *sTFA) TFAInfo(ctx context.Context, token string) (*entity.Tfa, error) {
	rst := entity.Tfa{}
	err := dao.Tfa.Ctx(ctx).Where(dao.Tfa.Columns().Token, token).Scan(&rst)
	if err != nil {
		g.Log().Error(ctx, "tfainfo:", err, token)
		return nil, gerror.NewCode(gcode.CodeOperationFailed)
	}
	return &rst, nil
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
	return &sTFA{
		pendding: map[string]func(){},
		ctx:      ctx,
	}

}

func init() {
	service.RegisterTFA(new())
}
