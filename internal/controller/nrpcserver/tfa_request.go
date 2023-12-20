package nats

import (
	"context"
	"riskcontral/api/risk/nrpc"
	"riskcontral/internal/model"
	"riskcontral/internal/model/do"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *NrpcServer) RpcTfaRequest(ctx context.Context, req *nrpc.TfaRequestReq) (res *nrpc.TfaRequestRes, err error) {
	g.Log().Notice(ctx, "RpcTfaRequest:", "req:", req)
	// limit
	if err := s.apiLimit(ctx, req.Token, "TfaRequest"); err != nil {
		return nil, err
	}
	/////
	//trace
	ctx, span := gtrace.NewSpan(ctx, "TfaRequest")
	defer span.End()
	/////
	//
	info, _ := service.UserInfo().GetUserInfo(ctx, req.Token)
	if info == nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	///
	// tfaInfo, err := service.TFA().TFAInfo(ctx, info.UserId)
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, info.UserId)
	if err != nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	///
	///
	var riskKind model.RiskKind = model.RiskKind_Nil
	//
	switch req.CodeType {
	case model.Type_TfaBindPhone:
		riskKind = model.RiskKind_BindPhone
		if tfaInfo != nil && tfaInfo.Phone != "" {
			return nil, mpccode.CodeTFAExist()
		}
		if tfaInfo == nil {
			err = service.DB().InsertTfaInfo(ctx, info.UserId, &do.Tfa{
				UserId:    info.UserId,
				TokenData: info,
				CreatedAt: gtime.Now(),
			})
			if err != nil {
				return nil, mpccode.CodeInternalError()
			}
		}
		///
	case model.Type_TfaBindMail:
		riskKind = model.RiskKind_BindMail
		if tfaInfo != nil && tfaInfo.Mail != "" {
			return nil, mpccode.CodeTFAExist()
		}
		if tfaInfo == nil {
			err = service.DB().InsertTfaInfo(ctx, info.UserId, &do.Tfa{
				UserId:    info.UserId,
				TokenData: info,
				CreatedAt: gtime.Now(),
			})
			if err != nil {
				return nil, mpccode.CodeInternalError()
			}
		}
		////
	case model.Type_TfaUpdatePhone:
		riskKind = model.RiskKind_UpPhone
		if tfaInfo == nil || tfaInfo.Phone == "" {
			return nil, mpccode.CodeTFANotExist()
		}
	case model.Type_TfaUpdateMail:
		riskKind = model.RiskKind_UpMail
		if tfaInfo == nil || tfaInfo.Mail == "" {
			return nil, mpccode.CodeTFANotExist()
		}
	default:
		return nil, mpccode.CodeParamInvalid()
	}
	///
	// tfaInfo, err = service.TFA().TFAInfo(ctx, info.UserId)
	tfaInfo, err = service.DB().FetchTfaInfo(ctx, info.UserId)
	if err != nil || tfaInfo == nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	///
	riskSerial, code := service.Risk().RiskTFA(ctx, tfaInfo, &model.RiskTfa{
		UserId: tfaInfo.UserId,
		Type:   req.CodeType,
	})
	if code == mpccode.RiskCodeForbidden {
		return nil, mpccode.CodePerformRiskForbidden()
	}
	if code == mpccode.RiskCodeError {
		return nil, mpccode.CodePerformRiskError()
	}
	///
	vl, _ := service.TFA().TfaRiskTidy(ctx, tfaInfo, riskSerial, riskKind)
	res = &nrpc.TfaRequestRes{
		RiskSerial: riskSerial,
		VList:      vl,
	}
	return res, nil
	///
}
