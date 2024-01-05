package nats

import (
	"context"
	"riskcontral/api/riskserver"
	"riskcontral/internal/model"
	"riskcontral/internal/model/do"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *NrpcServer) RpcTfaRequest(ctx context.Context, req *riskserver.TfaRequestReq) (*riskserver.TfaRequestRes, error) {
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
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, info.UserId)
	if err != nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	///
	///
	riskKind := model.CodeType2RiskKind(req.CodeType)
	//
	switch riskKind {
	case model.RiskKind_BindPhone:
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
	case model.RiskKind_BindMail:
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
	case model.RiskKind_UpPhone:
		if tfaInfo == nil || tfaInfo.Phone == "" {
			return nil, mpccode.CodeTFANotExist()
		}
	case model.RiskKind_UpMail:
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
	// res, err := service.NrpcClient().RiskTfaRequest(ctx, &riskctrl.TfaRiskReq{
	// 	Token:    req.Token,
	// 	CodeType: req.CodeType,
	// })
	code, err := service.TFA().RiskTfaRequest(ctx, tfaInfo, &model.RiskTfa{
		RiskKind: riskKind,
	})
	if err != nil {
		g.Log().Warning(ctx, "RpcRiskTxs:", "req:", req, "err:", err)
		return nil, mpccode.CodeInternalError()
	}
	///

	if code == mpccode.RiskCodeForbidden {
		return nil, mpccode.CodePerformRiskForbidden()
	}
	if code == mpccode.RiskCodeError {
		return nil, mpccode.CodePerformRiskError()
	}
	///
	riskSerial, vl, _ := service.TFA().RiskTfaTidy(ctx, tfaInfo, riskKind)
	return &riskserver.TfaRequestRes{
		RiskSerial: riskSerial,
		VList:      vl,
	}, nil
	///
}
