package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedTFAServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterTFAServer(s.Server, &Controller{})
}

func (*Controller) CallTFAInfo(ctx context.Context, req *v1.TFAReq) (res *v1.TFARes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
