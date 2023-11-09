package tfa

import (
	"context"
	"riskcontral/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sTFA) VerifyCode(ctx context.Context, userId string, riskSerial string, code *model.VerifyCode) error {
	risk := s.riskPenddingContainer.GetRiskVerify(userId, riskSerial)
	if risk == nil {
		return errRiskNotExist
	}
	k, err := risk.VerifierCode(code)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userid", userId),
			mpccode.ErrDetail("riskSerial", riskSerial),
			mpccode.ErrDetail("code", code),
			mpccode.ErrDetail("kind", k),
			mpccode.ErrDetail("risk", risk),
		))
		return err
	}
	k, err = risk.DoFunc(ctx)
	// err = s.riskPenddingContainer.DoAfter(ctx, userId, riskSerial)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userid", userId),
			mpccode.ErrDetail("riskSerial", riskSerial),
			mpccode.ErrDetail("code", code),
			mpccode.ErrDetail("kind", k),
		))
		// return gerror.NewCode(consts.CodeRiskVerifyCodeInvalid)
		return err
	}

	return nil
}
