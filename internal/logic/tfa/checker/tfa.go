package check

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

func (s *Checker) upTimeAfterForbiddent(ctx context.Context, uptime *gtime.Time) bool {
	if uptime == nil {
		return true
	}
	///
	forbiddentTime := gtime.Now().Add(s.forbiddentTime)
	if uptime.Before(forbiddentTime) {
		return true
	}
	return false
}
func (s *Checker) CheckTfaRisk(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
	/////
	if tfaInfo == nil {
		g.Log().Warning(ctx, "checkTFAUpPhone tfaInfo not exists:", tfaInfo)
		return mpccode.RiskCodeNeedVerification, nil
	}
	if tfaInfo.Mail == "" && tfaInfo.Phone == "" {
		g.Log().Warning(ctx, "checkTFAUpPhone tfaInfo not exists:", tfaInfo)
		return mpccode.RiskCodeError, nil
	}
	/////
	if tfaInfo.Mail != "" {
		g.Log().Notice(ctx, "checkTfaRisk check mailUpTime:",
			"tfaInfo:", tfaInfo,
		)
		if s.upTimeAfterForbiddent(ctx, tfaInfo.MailUpdatedAt) {
			return mpccode.RiskCodeForbidden, nil
		}
	}
	if tfaInfo.Phone != "" {
		g.Log().Notice(ctx, "checkTfaRisk check phoneUpTime:",
			"tfaInfo:", tfaInfo,
		)
		if s.upTimeAfterForbiddent(ctx, tfaInfo.PhoneUpdatedAt) {
			return mpccode.RiskCodeForbidden, nil
		}
	}
	/////
	return mpccode.RiskCodeNeedVerification, nil
}

// /
// /
func (s *Checker) CheckTfaUpPhone(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
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
	g.Log().Notice(ctx, "checkTfaUpPhone check phoneUpTime:",
		"tfaInfo:", tfaInfo,
	)
	if s.upTimeAfterForbiddent(ctx, tfaInfo.PhoneUpdatedAt) {
		return mpccode.RiskCodeForbidden, nil
	}
	return mpccode.RiskCodeNeedVerification, nil
}

func (s *Checker) CheckTfaUpMail(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
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

	g.Log().Notice(ctx, "checkTfaUpMail check mailUpTime:",
		"tfaInfo:", tfaInfo,
	)
	if s.upTimeAfterForbiddent(ctx, tfaInfo.MailUpdatedAt) {
		return mpccode.RiskCodeForbidden, nil
	}
	////

	return mpccode.RiskCodeNeedVerification, nil

}

// /
func (s *Checker) CheckTfaBindPhone(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {

	return mpccode.RiskCodeNeedVerification, nil
}
func (s *Checker) CheckTfaBindMail(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
	return mpccode.RiskCodeNeedVerification, nil
}
func (s *Checker) CheckPersonRisk(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
	return mpccode.RiskCodeNeedVerification, nil
}
