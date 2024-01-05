package tfa

import (
	"context"
	"riskcontral/internal/config"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

type UserRiskId string

func keyUserRiskId(userId string, riskSerial string) UserRiskId {
	return UserRiskId(userId + "keyUserRiskId" + riskSerial)
}

type sTFA struct {
	// riskClient riskv1.UserClient
	ctx                   context.Context
	riskPenddingContainer *model.RiskPenddingContainer
	////
}

// /
func new() *sTFA {

	ctx := gctx.GetInitCtx()
	//
	t := config.Config.Cache.VerificationCodeDuration
	s := &sTFA{
		//todo:
		riskPenddingContainer: model.NewRiskPenddingContainer(t),
		ctx:                   ctx,
	}
	///

	return s
}

///

func init() {
	service.RegisterTFA(new())
}

func (s *sTFA) TfaRiskKind(ctx context.Context, tfaInfo *entity.Tfa, riskSerial string) (model.RiskKind, error) {
	risk := s.riskPenddingContainer.GetRiskVerify(tfaInfo.UserId, riskSerial)
	if risk == nil {
		g.Log().Warning(ctx, "TfaRiskKind:", "tfaInfo:", tfaInfo, "riskSerial:", riskSerial)
		return "", mpccode.CodeParamInvalid()
	}
	return risk.RiskKind, nil
}
