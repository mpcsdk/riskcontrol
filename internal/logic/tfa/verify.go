package tfa

import (
	"context"
	"riskcontral/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sTFA) VerifyCode(ctx context.Context, userId string, riskSerial string, code *model.VerifyCode) error {
	risk := s.riskPenddingContainer.GetRiskVerify(userId, riskSerial)
	if risk == nil {
		return gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	k, err := risk.VerifierCode(code)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userid", userId),
			mpccode.ErrDetail("riskSerial", riskSerial),
			mpccode.ErrDetail("code", code),
			mpccode.ErrDetail("kind", k),
		))
		return err
	}
	k, err = risk.DoFunc(ctx)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userid", userId),
			mpccode.ErrDetail("riskSerial", riskSerial),
			mpccode.ErrDetail("code", code),
			mpccode.ErrDetail("kind", k),
		))
		g.Log().Warningf(ctx, "%+v", err)
		return gerror.NewCode(mpccode.CodeRiskVerifyCodeInvalid)
		// return err
	}

	return nil
}
