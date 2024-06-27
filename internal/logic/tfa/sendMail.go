package tfa

import (
	"context"
	"riskcontrol/internal/logic/tfa/tfaconst"
	"riskcontrol/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sTFA) SendMailCode(ctx context.Context, userId string, riskSerial string) error {
	tfaInfo, err := service.DB().TfaDB().FetchTfaInfo(ctx, userId)
	if err != nil || tfaInfo == nil {
		return mpccode.CodeTFANotExist()
	}
	///
	risk := s.riskPenddingContainer.GetRiskPendding(userId, riskSerial)
	if risk == nil {
		return mpccode.CodeRiskSerialNotExist()
	}
	// //
	v := risk.GetVerifier(tfaconst.VerifierKind_Mail)
	if v == nil {
		return mpccode.CodeRiskSerialNotExist()
	}
	if err := s.limitSendMailCnt(ctx, tfaInfo.UserId, v.Destination()); err != nil {
		return mpccode.CodeLimitSendMailCode()
	}
	////
	code, err := v.SendVerificationCode()
	if err != nil {
		g.Log().Warning(ctx, "SendMailCode:", "tfaInfo:", tfaInfo, "risk:", risk, "err:", err)
		return mpccode.CodeTFASendMailFailed()
	}
	////
	g.Log().Notice(ctx, "SendMailCode:", "tfaInfo:", tfaInfo, "risk:", risk, "code:", code)
	v.SetCode(code)

	return nil
}
