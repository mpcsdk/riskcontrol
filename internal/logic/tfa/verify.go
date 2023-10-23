package tfa

import (
	"context"
	"riskcontral/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sTFA) VerifyCode(ctx context.Context, userId string, riskSerial string, code *model.VerifyCode) error {

	k, err := s.riskPenddingContainer.VerifierCode(userId, riskSerial, code)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userid", userId),
			mpccode.ErrDetail("riskSerial", riskSerial),
			mpccode.ErrDetail("code", code),
			mpccode.ErrDetail("kind", k),
		))
		return err
	}
	err = s.riskPenddingContainer.DoAfter(ctx, userId, riskSerial)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userid", userId),
			mpccode.ErrDetail("riskSerial", riskSerial),
			mpccode.ErrDetail("code", code),
		))

		// return gerror.NewCode(consts.CodeRiskVerifyCodeInvalid)
		return err
	}

	return nil
}
