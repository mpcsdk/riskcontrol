package tfa

import (
	"context"
	"riskcontrol/internal/logic/tfa/tfaconst"
	"riskcontrol/internal/service"

	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sTFA) TfaTidyMail(ctx context.Context, userId string, riskSerial string, mail string) error {
	risk := s.riskPenddingContainer.GetRiskPendding(userId, riskSerial)
	if risk == nil {
		return mpccode.CodeRiskSerialNotExist()
	}
	riskKind := risk.RiskKind()
	if riskKind != tfaconst.RiskKind_BindMail && riskKind != tfaconst.RiskKind_UpMail {
		return nil
	}
	if mail == "" {
		return mpccode.CodeParamInvalid()
	}
	return risk.TfaTidyMail(ctx, userId, mail)
}

func (s *sTFA) TfaTidyPhone(ctx context.Context, userId string, phone string, riskSerial string) error {
	////
	if phone == "" {
		return mpccode.CodeParamInvalid()
	}
	notexists, err := service.DB().TfaDB().TfaPhoneNotExists(ctx, phone)
	if err != nil || !notexists {
		return mpccode.CodeTFAPhoneExists()
	}
	/////
	risk := s.riskPenddingContainer.GetRiskPendding(userId, riskSerial)
	riskKind := risk.RiskKind()
	if riskKind != tfaconst.RiskKind_BindPhone && riskKind != tfaconst.RiskKind_UpPhone {
		return nil
	}
	return risk.TfaTidyPhone(ctx, userId, phone)
}
