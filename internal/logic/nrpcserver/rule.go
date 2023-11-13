package nats

import (
	"context"
	v1 "riskcontral/api/risk/nrpc/v1"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (*sNrpcServer) RpcAllAbi(ctx context.Context, req *v1.AllAbiReq) (res *v1.AllAbiRes, err error) {
	rst, err := service.DB().GetAbiAll(ctx)
	if err != nil {
		return nil, gerror.NewCode(mpccode.CodeInternalError)
	}
	res = &v1.AllAbiRes{
		Abis: map[string]string{},
	}
	for _, v := range rst {
		res.Abis[v.Addr] = v.Abi
	}

	return res, nil
}

func (*sNrpcServer) RpcAllNftRules(ctx context.Context, req *v1.NftRulesReq) (res *v1.NftRulesRes, err error) {
	rst, err := service.DB().GetNftRules(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeInternalError)
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

func (*sNrpcServer) RpcAllFtRules(ctx context.Context, req *v1.FtRulesReq) (res *v1.FtRulesRes, err error) {
	rst, err := service.DB().GetFtRules(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeInternalError)
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
