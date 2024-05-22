package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/logic/tfa/tfaconst"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/do"
)

func (s *sTFA) TfaRequest(ctx context.Context, userId string, riskKind tfaconst.RISKKIND, data *v1.RequestData) (*v1.TfaRequestRes, error) {
	g.Log().Notice(ctx, "RpcTfaRequest:", "userId:", userId, "riskKind:", riskKind)
	/////
	//
	// riskKind := tfaconst.CodeType2RiskKind(codeType)
	///
	tfaInfo, err := service.DB().TfaDB().FetchTfaInfo(ctx, userId)
	if err != nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	if tfaInfo == nil {
		err = service.DB().TfaDB().InsertTfaInfo(ctx, userId, &do.Tfa{
			UserId:    userId,
			TokenData: userId,
			CreatedAt: gtime.Now(),
		})
		if err != nil {
			return nil, mpccode.CodeInternalError()
		}
		///
		tfaInfo, err = service.DB().TfaDB().FetchTfaInfo(ctx, userId)
		if err != nil || tfaInfo == nil {
			g.Log().Warning(ctx, "RpcTfaRequest:", "userId:", userId, "riskKind:", riskKind, "err:", err)
			return nil, mpccode.CodeTokenInvalid()
		}
	}
	///check riskKind
	code, err := s.checker.CheckKind(ctx, tfaInfo, riskKind, data)
	///
	if err != nil {
		g.Log().Warning(ctx, "RpcTfaRequest:", "userId:", userId, "riskKind:", riskKind, "err:", err)
		return nil, err
	}
	///
	if code == mpccode.RiskCodeForbidden {
		return nil, mpccode.CodePerformRiskForbidden()
	}
	if code == mpccode.RiskCodeError {
		return nil, mpccode.CodePerformRiskError()
	}
	/////
	// if riskKind == tfaconst.RiskKind_PersonRisk {
	// 	if tfaInfo.Enable == data.Enable {
	// 		return nil, mpccode.CodeParamInvalid()
	// 	}
	// }
	////new risk
	risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo, riskKind, data)
	return &v1.TfaRequestRes{
		RiskSerial: risk.RiskSerial(),
		VList:      risk.VerifyKind(),
	}, nil
	///
}
