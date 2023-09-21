package risk

import (
	"context"
	v1 "riskcontral/api/risk/v1"
	"riskcontral/common/ethtx"
	"riskcontral/internal/consts"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/service"
	"strings"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/golang/protobuf/ptypes/empty"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) PerformSmsCode(ctx context.Context, req *v1.SmsCodeReq) (res *v1.SmsCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "PerformSmsCode")
	defer span.End()
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	// err = service.Risk().RiskPhoneCode(ctx, req.RiskSerial)
	_, err = service.TFA().SendPhoneCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {
		g.Log().Error(ctx, "PerformSmsCode:", req, err)
	}
	return nil, err
}

func (*Controller) PerformMailCode(ctx context.Context, req *v1.MailCodekReq) (res *v1.MailCodekRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "PerformMailCode")
	defer span.End()
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	// err = service.Risk().RiskMailCode(ctx, req.RiskSerial)
	_, err = service.TFA().SendMailCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {
		g.Log().Error(ctx, "PerformMailCode:", req, err)
	}
	return nil, err
}

func (*Controller) PerformVerifyCode(ctx context.Context, req *v1.VerifyCodekReq) (res *v1.VerifyCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "PerformVerifyCode")
	defer span.End()
	//
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	// err = service.Risk().VerifyCode(ctx, req.RiskSerial, req.Code)
	err = service.TFA().VerifyCode(ctx, info.UserId, req.RiskSerial, req.Code)
	if err != nil {
		g.Log().Error(ctx, "PerformVerifyCode:", req, err)
	}
	return nil, err
}

func (*Controller) PerformAlive(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "PerformAlive")
	defer span.End()
	//
	return &empty.Empty{}, nil
}

func (*Controller) PerformRiskTxs(ctx context.Context, req *v1.TxRiskReq) (res *v1.TxRiskRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "performRiskTxs")
	defer span.End()
	//
	///
	g.Log().Debug(ctx, "PerformRiskTxs:", req)
	req.Address = strings.ToLower(req.Address)
	///
	txs := []*conrisk.RiskTx{}
	for _, tx := range req.Txs {
		contract := strings.ToLower(tx.Contract)
		contractabi, err := service.RulesDb().GetAbi(ctx, contract)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError)
		}
		tx, err := ethtx.AnalzyTxData(contractabi, tx.TxData)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError)
		}
		////
		txs = append(txs, &conrisk.RiskTx{
			Address:  req.Address,
			Contract: contract,
			//
			MethodName: tx.MethodName,
			MethodId:   tx.MethodId,
			Args:       tx.Args,
			// From: tx.Args[]
		})
		g.Log().Debug(ctx, "PerformRiskTx:", tx)
	}
	serial, code, err := service.Risk().PerformRiskTxs(ctx, req.UserId, req.Address, txs)
	g.Log().Info(ctx, "PerformRiskTx:", req, serial)
	if err != nil {
		g.Log().Error(ctx, "PerformRiskTx", serial, err)
	}
	//
	//notice: wait verify
	kinds, err := service.TFA().PerformRiskTFA(ctx, req.UserId, serial)
	if err != nil {
		g.Log().Warning(ctx, "PerformRiskTxs:", "PerformRiskTFA:", req.UserId, serial)
		return nil, gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	///
	return &v1.TxRiskRes{
		Ok:         code,
		RiskSerial: serial,
		//todo:
		RiskKind: kinds,
	}, err
}

func (*Controller) PerformRiskTFA(ctx context.Context, req *v1.TFARiskReq) (res *v1.TFARiskRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
