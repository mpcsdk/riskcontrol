package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/model"
	"riskcontral/internal/model/do"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/userInfoGeter"
)

func (s *sTFA) TfaRequest(ctx context.Context, info *userInfoGeter.UserInfo, riskKind model.RiskKind) (*v1.TfaRequestRes, error) {
	g.Log().Notice(ctx, "RpcTfaRequest:", "info:", info, "riskKind:", riskKind)

	/////
	//
	///
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, info.UserId)
	if err != nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	///
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

	code, err := s.RiskTfaRequest(ctx, tfaInfo, riskKind)
	if err != nil {
		g.Log().Warning(ctx, "RpcTfaRequest:", "info:", info, "riskKind:", riskKind, "err:", err)
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
	return &v1.TfaRequestRes{
		RiskSerial: riskSerial,
		VList:      vl,
	}, nil
	///
}
