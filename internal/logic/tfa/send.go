package tfa

import (
	"context"
	"riskcontral/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// /
func (s *sTFA) SendPhoneCode(ctx context.Context, userId string, riskSerial string) (string, error) {
	risk := s.riskPenddingContainer.GetRiskVerify(userId, riskSerial)
	if risk == nil {
		return "", mpccode.CodeRiskSerialNotExist()
	}

	v := risk.Verifier(model.VerifierKind_Phone)
	if v == nil {
		return "", mpccode.CodeRiskSerialNotExist()
	}
	code, err := v.SendVerificationCode()
	if err != nil {
		return string(model.VerifierKind_Phone), err
	}
	////

	g.Log().Notice(ctx, "SendPhoneCode:", "userId:", userId, "riskSerial:", riskSerial, "code:", code)
	v.SetCode(code)

	return "", nil

}

func (s *sTFA) SendMailCode(ctx context.Context, userId string, riskSerial string) (string, error) {
	risk := s.riskPenddingContainer.GetRiskVerify(userId, riskSerial)
	if risk == nil {
		return "", mpccode.CodeRiskSerialNotExist()
	}

	v := risk.Verifier(model.VerifierKind_Mail)
	if v == nil {
		return "", mpccode.CodeRiskSerialNotExist()
	}
	code, err := v.SendVerificationCode()
	if err != nil {

		return string(model.VerifierKind_Mail), err
	}
	////
	g.Log().Notice(ctx, "SendMailCode:", "userId:", userId, "riskSerial:", riskSerial, "code:", code)
	v.SetCode(code)

	return "", nil
}
