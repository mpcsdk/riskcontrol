package risk

import (
	"context"
	v1 "riskcontral/api/risk/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/model"
	"riskcontral/internal/service"

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

	code := &model.VerifyCode{
		MailCode:  req.MailCode,
		PhoneCode: req.PhoneCode,
	}
	err = service.TFA().VerifyCode(ctx, info.UserId, req.RiskSerial, code)
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
	///
	g.Log().Debug(ctx, "PerformRiskTxs:", req)

	/////
	serial, code := service.Risk().PerformRiskTxs(ctx, req.UserId, req.SignTxData)
	if code == consts.RiskCodeError {
		return nil, gerror.NewCode(consts.CodePerformRiskError)
	}
	///: pass or forbidden
	g.Log().Debug(ctx, "PerformRiskTxs:", req, serial, code)
	//

	if code == consts.RiskCodePass {
		return &v1.TxRiskRes{
			Ok: code,
		}, nil
	}
	if code == consts.RiskCodeForbidden {
		return &v1.TxRiskRes{
			Ok: code,
		}, nil
	}
	///
	//
	//notice:  tfatx  need verification
	kinds, err := service.TFA().TFATx(ctx, req.UserId, serial)
	if err != nil {
		g.Log().Warning(ctx, "PerformRiskTxs:", "PerformRiskTFA:", req.UserId, serial)
		return nil, gerror.NewCode(consts.CodePerformRiskError)
	}
	///
	g.Log().Info(ctx, "PerformRiskTFA:", req.UserId, serial, kinds)
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

func (*Controller) PerformAllAbi(ctx context.Context, req *v1.AllAbiReq) (res *v1.AllAbiRes, err error) {
	rst, err := service.DB().GetAbiAll(ctx)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	res = &v1.AllAbiRes{
		Abis: map[string]string{},
	}
	for _, v := range rst {
		res.Abis[v.Addr] = v.Abi
	}

	return res, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformAllNftRules(ctx context.Context, req *v1.NftRulesReq) (res *v1.NftRulesRes, err error) {
	rst, err := service.DB().GetNftRules(ctx)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	res = &v1.NftRulesRes{
		NftRules: map[string]*v1.NftRules{},
	}
	///
	for k, v := range rst {
		res.NftRules[k] = &v1.NftRules{
			Contract:           v.Contract,
			Name:               v.Name,
			MethodName:         v.MethodName,
			MethodSig:          v.MethodSig,
			MethodFromField:    v.MethodFromField,
			MethodToField:      v.MethodToField,
			MethodTokenIdField: v.MethodTokenIdField,

			EventName:         v.EventName,
			EventSig:          v.EventSig,
			EventTopic:        v.EventTopic,
			EventFromField:    v.EventFromField,
			EventToField:      v.EventToField,
			EventTokenIdField: v.EventTokenIdField,

			SkipToAddr: v.SkipToAddr,
			Threshold:  int32(v.Threshold),
		}
	}
	return res, nil
}

func (*Controller) PerformAllFtRules(ctx context.Context, req *v1.FtRulesReq) (res *v1.FtRulesRes, err error) {
	rst, err := service.DB().GetFtRules(ctx)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	res = &v1.FtRulesRes{
		FtRules: map[string]*v1.Ftrules{},
	}
	///
	for k, v := range rst {
		res.FtRules[k] = &v1.Ftrules{
			Contract:         v.Contract,
			Name:             v.Name,
			MethodName:       v.MethodName,
			MethodSig:        v.MethodSig,
			MethodFromField:  v.MethodFromField,
			MethodToField:    v.MethodToField,
			MethodValueField: v.MethodValueField,

			EventName:       v.EventName,
			EventSig:        v.EventSig,
			EventTopic:      v.EventTopic,
			EventFromField:  v.EventFromField,
			EventToField:    v.EventToField,
			EventValueField: v.EventValueField,

			SkipToAddr:           v.SkipToAddr,
			ThresholdBigintBytes: v.Threshold.Bytes(),
		}
	}
	return res, nil
}
