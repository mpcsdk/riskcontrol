package tfa

import (
	"context"
	"riskcontral/internal/model"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// /
func (s *sTFA) SendPhoneCode(ctx context.Context, userId string, riskSerial string, phone string) error {
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, userId)
	if err != nil || tfaInfo == nil {
		return mpccode.CodeTFANotExist()
	}
	///
	risk := service.TFA().GetRiskVerify(ctx, userId, riskSerial)
	if risk == nil {
		return mpccode.CodeRiskSerialNotExist()
	} ////
	switch risk.RiskKind {
	case model.RiskKind_BindPhone, model.RiskKind_UpPhone:
		if phone == "" {
			return mpccode.CodeParamInvalid()
		}
		notexists, err := service.DB().TfaPhoneNotExists(ctx, phone)
		if err != nil || !notexists {
			return mpccode.CodeTFAPhoneExists()
		}
		////
		service.TFA().TfaSetPhone(ctx, tfaInfo, phone, risk.RiskSerial, risk.RiskKind)
		///
	case model.RiskKind_BindMail, model.RiskKind_UpMail:
	default:
		return mpccode.CodeRiskSerialNotExist()
	}
	///
	v := risk.GetVerifier(model.VerifierKind_Phone)
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
