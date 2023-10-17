package rpc

import (
	"context"
	"errors"
	"math/big"
	"time"

	scrapev1 "riskcontral/api/scrapelogs/v1"
	"riskcontral/internal/config"
	"riskcontral/internal/service"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/protobuf/types/known/emptypb"
)

type sRPC struct {
	ctx    g.Ctx
	client scrapev1.ScrapeLogsAggClient
}

var timeout = 3 * time.Second
var errDeadLine = errors.New("context deadline exceeded")

func (s *sRPC) PerformNftCnt(ctx context.Context, addr string, contract string, method string) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	///
	rst, err := s.client.PerformNftCnt(ctx, &scrapev1.NftCntReq{
		Address:  addr,
		Contract: contract,
		Method:   method,
	})
	if err != nil {
		return 0, err
	}
	return int(rst.Cnt), err
}
func (s *sRPC) PerformFtCnt(ctx context.Context, addr string, contract string, method string) (*big.Int, error) {

	subctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	rst, err := s.client.PerformFtCnt(subctx, &scrapev1.FtCntReq{
		Address:  addr,
		Contract: contract,
		Method:   method,
	})
	if err != nil {
		return big.NewInt(0), err
	}

	return big.NewInt(0).SetBytes(rst.Cnt_BigInt_Bytes), nil
}
func (s *sRPC) PerformAlive(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	_, err := s.client.PerformAlive(ctx, &emptypb.Empty{})

	return err
}

func new() *sRPC {
	ctx := gctx.GetInitCtx()

	g.Log().Info(ctx, "etcd address...:", config.Config.Etcd.Address, config.Config.Etcd.ScrapeLogsRpc)
	grpcx.Resolver.Register(etcd.New(config.Config.Etcd.Address))

	conn, err := grpcx.Client.NewGrpcClientConn(
		config.Config.Etcd.ScrapeLogsRpc,
	)
	if err != nil {
		g.Log().Error(ctx, "etcd err:", err)
	}
	g.Log().Notice(ctx, "etcd RiskRpc stat:", conn.GetState().String())
	client := scrapev1.NewScrapeLogsAggClient(conn)

	s := &sRPC{
		ctx:    ctx,
		client: client,
	}
	err = s.PerformAlive(ctx)
	if err != nil {
		g.Log().Error(ctx, "PerformAlive:", err)
	}
	return s
}

func init() {
	service.RegisterRPC(new())
}
