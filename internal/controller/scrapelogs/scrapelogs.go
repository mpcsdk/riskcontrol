package scrapelogs

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang/protobuf/ptypes/empty"
)

type Controller struct {
	v1.UnimplementedTFAServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterTFAServer(s.Server, &Controller{})
}

func (*Controller) PerformNftCnt(ctx context.Context, req *v1.NftCntReq) (res *v1.NftCntRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformFtCnt(ctx context.Context, req *v1.FtCntReq) (res *v1.FtCntRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformAlive(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
