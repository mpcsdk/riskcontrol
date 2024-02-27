package tfa

import (
	"context"
	"riskcontral/internal/model"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sTFA) SendMailCode(ctx context.Context, userId string, riskSerial string, mail string) error {
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, userId)
	if err != nil || tfaInfo == nil {
		return mpccode.CodeTFANotExist()
	}
	///
	risk := service.TFA().GetRiskVerify(ctx, userId, riskSerial)
	if risk == nil {
		return mpccode.CodeRiskSerialNotExist()
	}
	// //
	switch risk.RiskKind {
	case model.RiskKind_BindMail, model.RiskKind_UpMail:
		if mail == "" {
			return mpccode.CodeParamInvalid()
		}
		notexists, err := service.DB().TfaMailNotExists(ctx, mail)
		if err != nil || !notexists {
			return mpccode.CodeTFAMailExists()
		}
		////
		service.TFA().TfaSetMail(ctx, tfaInfo, mail, risk.RiskSerial, risk.RiskKind)
		///
	case model.RiskKind_BindPhone, model.RiskKind_UpPhone:
	default:
		return mpccode.CodeRiskSerialNotExist()
	}
	///
	v := risk.GetVerifier(model.VerifierKind_Mail)
	if v == nil {
		return mpccode.CodeRiskSerialNotExist()
	}
	if err := s.limitSendMailCnt(ctx, tfaInfo.UserId, v.Destination()); err != nil {
		return mpccode.CodeLimitSendMailCode()
	}
	////
	code, err := v.SendVerificationCode()
	if err != nil {
		return mpccode.CodeTFASendMailFailed()
	}
	////
	g.Log().Notice(ctx, "SendMailCode:", "tfaInfo:", tfaInfo, "risk:", risk, "code:", code)
	v.SetCode(code)

	return nil
}
