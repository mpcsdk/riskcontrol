package tfa

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// /
func (s *sTFA) SendPhoneCode(ctx context.Context, userId string, riskSerial string) (string, error) {
	risk := s.riskPenddingContainer.GetRiskVerify(userId, riskSerial)
	if risk == nil {
		return "", errRiskNotExist
	}

	v := risk.Verifier(VerifierKind_Phone)
	if v == nil {
		return "", errRiskNotExist
	}
	code, err := v.SendVerificationCode()
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userId", userId),
			mpccode.ErrDetail("riskSerial", riskSerial),
		))
		return string(VerifierKind_Phone), err
	}
	////

	g.Log().Notice(ctx, "SendPhoneCode:", "userId:", userId, "riskSerial:", riskSerial, "code:", code)
	v.SetCode(code)

	return "", nil

}

func (s *sTFA) SendMailCode(ctx context.Context, userId string, riskSerial string) (string, error) {
	risk := s.riskPenddingContainer.GetRiskVerify(userId, riskSerial)
	if risk == nil {
		return "", errRiskNotExist
	}

	v := risk.Verifier(VerifierKind_Mail)
	if v == nil {
		return "", errRiskNotExist
	}
	code, err := v.SendVerificationCode()
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userId", userId),
			mpccode.ErrDetail("riskSerial", riskSerial),
		))
		return string(VerifierKind_Mail), err
	}
	////
	g.Log().Notice(ctx, "SendMailCode:", "userId:", userId, "riskSerial:", riskSerial, "code:", code)
	v.SetCode(code)

	return "", nil
}
