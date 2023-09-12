package risk

import (
	"context"
	v1 "riskcontral/api/risk/v1"
	"riskcontral/common/ethtx"
	"riskcontral/internal/service"
	"strings"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang/protobuf/ptypes/empty"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

// func (*Controller) PerformRisk(ctx context.Context, req *v1.RiskReq) (res *v1.RiskRes, err error) {
// 	if req.RuleName == "phone" {
// 		return nil, nil
// 	} else if req.RuleName == "mail" {
// 		return nil, gerror.NewCode(gcode.CodeNotImplemented)
// 	} else if req.RuleName == "tx" {
// 		return nil, gerror.NewCode(gcode.CodeNotImplemented)
// 	}

// 	return nil, gerror.NewCode(gcode.CodeNotImplemented)
// }

func (*Controller) PerformSmsCode(ctx context.Context, req *v1.SmsCodeReq) (res *v1.SmsCodeRes, err error) {
	// service.SmsCode().SendCode(ctx)
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformMailCode(ctx context.Context, req *v1.MailCodekReq) (res *v1.MailCodekRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformAlive(context.Context, *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (*Controller) PerformRiskTx(ctx context.Context, req *v1.TxRiskReq) (res *v1.TxRiskRes, err error) {
	//
	req.Address = strings.ToLower(req.Address)
	req.Contract = strings.ToLower(req.Contract)
	///
	contractabi, err := service.RulesDb().GetAbi(ctx, req.Contract)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	tx, err := ethtx.AnalzyTxData(contractabi, req.TxData)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	////
	g.Log().Debug(ctx, "PerformRiskTx:", tx)

	///
	return &v1.TxRiskRes{
		Ok: true,
	}, nil
}

func (*Controller) PerformRiskTFA(ctx context.Context, req *v1.TFARiskReq) (res *v1.TFARiskRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
