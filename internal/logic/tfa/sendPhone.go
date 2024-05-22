package tfa

import (
	"context"
	"riskcontral/internal/logic/tfa/tfaconst"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// /
func (s *sTFA) SendPhoneCode(ctx context.Context, userId string, riskSerial string) error {
	tfaInfo, err := service.DB().TfaDB().FetchTfaInfo(ctx, userId)
	if err != nil || tfaInfo == nil {
		return mpccode.CodeTFANotExist()
	}
	///
	risk := s.riskPenddingContainer.GetRiskPendding(userId, riskSerial)
	if risk == nil {
		return mpccode.CodeRiskSerialNotExist()
	} ////
	///
	v := risk.GetVerifier(tfaconst.VerifierKind_Phone)
	if v == nil {
		return mpccode.CodeRiskSerialNotExist()
	}
	if err := s.limitSendPhoneCnt(ctx, tfaInfo.UserId, v.Destination()); err != nil {
		return mpccode.CodeLimitSendPhoneCode()
	}
	///
	code, err := v.SendVerificationCode()
	if err != nil {
		return mpccode.CodeTFASendSmsFailed()
	}
	////

	g.Log().Notice(ctx, "SendPhoneCode:", "tfaInfo:", tfaInfo, "risk:", risk, "code:", code)
	v.SetCode(code)

	return nil

}
