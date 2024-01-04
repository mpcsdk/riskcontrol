package risk

import (
	"context"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sRisk) checkTfaUpPhone(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
	/////
	if tfaInfo == nil {
		g.Log().Warning(ctx, "checkTFAUpPhone tfaInfo not exists:", tfaInfo)
		return mpccode.RiskCodeNeedVerification, nil
	}
	if tfaInfo.Mail == "" && tfaInfo.Phone == "" {
		g.Log().Warning(ctx, "checkTFAUpPhone tfaInfo not exists:", tfaInfo)
		return mpccode.RiskCodeError, nil
	}
	if tfaInfo.PhoneUpdatedAt == nil {
		g.Log().Notice(ctx, "checkTFAUpPhone notuptime :",
			"tfaInfo:", tfaInfo,
			"info.PhoneUpdatedAt:", tfaInfo.PhoneUpdatedAt)
		return mpccode.RiskCodeNeedVerification, nil
	}

	befor24h := gtime.Now().Add(BeforH24)
	g.Log().Notice(ctx, "checkTFAUpPhone check phoneUpTime:",
		"tfaInfo:", tfaInfo,
		"info.PhoneUpdatedAt:", tfaInfo.PhoneUpdatedAt.Local(),
		"beforAt:", befor24h,
	)
	if tfaInfo.PhoneUpdatedAt.Before(befor24h) {
		return mpccode.RiskCodeNeedVerification, nil
	}
	return mpccode.RiskCodeForbidden, nil
}

func (s *sRisk) checkTfaUpMail(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
	if tfaInfo == nil {
		g.Log().Warning(ctx, "checkTfaUpMail userinfo not exists:", tfaInfo)
		return mpccode.RiskCodeNeedVerification, nil
	}
	if tfaInfo.Mail == "" && tfaInfo.Phone == "" {
		g.Log().Warning(ctx, "checkTfaUpMail tfaInfo not exists:", tfaInfo)
		return mpccode.RiskCodeError, nil
	}

	///
	if tfaInfo.MailUpdatedAt == nil {
		g.Log().Notice(ctx, "checkTfaUpMail notuptime :",
			"tfaInfo:", tfaInfo,
			"info.MailUpdatedAt:", tfaInfo.MailUpdatedAt)
		return mpccode.RiskCodeNeedVerification, nil
	}
	befor24h := gtime.Now().Add(BeforH24)
	g.Log().Notice(ctx, "checkTfaUpMail check phoneUpTime:",
		"tfaInfo:", tfaInfo,
		"info.MailUpdatedAt:", tfaInfo.MailUpdatedAt,
		"beforAt:", befor24h,
	)
	if tfaInfo.MailUpdatedAt.Before(befor24h) {
		return mpccode.RiskCodeNeedVerification, nil
	}
	return mpccode.RiskCodeForbidden, nil
}

func (s *sRisk) checkTfaBindPhone(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
	return mpccode.RiskCodePass, nil
}
func (s *sRisk) checkTfaBindMail(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
	return mpccode.RiskCodePass, nil
}
