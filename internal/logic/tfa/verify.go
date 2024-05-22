package tfa

import (
	"context"
	"riskcontral/internal/model"

	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sTFA) VerifyCode(ctx context.Context, userId string, riskSerial string, code *model.VerifyCode) error {
	risk := s.riskPenddingContainer.GetRiskPendding(userId, riskSerial)
	if risk == nil {
		return mpccode.CodeRiskSerialNotExist()
	}
	_, err := risk.VerifierCode(ctx, code)
	if err != nil {

		return err
	}
	_, err = risk.DoFunc(ctx)
	if err != nil {

		return mpccode.CodeRiskVerifyCodeInvalid()
	}

	return nil
}
