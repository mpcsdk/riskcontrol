package tfa

import (
	"context"
	"riskcontral/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	"github.com/mpcsdk/mpcCommon/rand"
)

func (s *sTFA) RiskTfaRequest(ctx context.Context,
	tfaInfo *entity.Tfa,
	riskKind model.RiskKind) (int32, error) {

	//

	///
	code := mpccode.RiskCodePass
	var err error
	///
	switch riskKind {
	case model.RiskKind_UpPhone:
		code, err = s.checkTfaUpPhone(ctx, tfaInfo)
	case model.RiskKind_UpMail:
		code, err = s.checkTfaUpMail(ctx, tfaInfo)
	case model.RiskKind_BindPhone:
		code, err = s.checkTfaBindPhone(ctx, tfaInfo)
	case model.RiskKind_BindMail:
		code, err = s.checkTfaBindMail(ctx, tfaInfo)
	default:
		g.Log().Error(ctx, "RiskTFA:", "req:", riskKind, "not support")
		return code, mpccode.CodeParamInvalid()
	}
	if err != nil {
		g.Log().Warning(ctx, "RiskTFA:", "tfaInfo:", tfaInfo, "riskKind:", riskKind, "err:", err)
		return mpccode.RiskCodeError, mpccode.CodeInternalError()
	}

	return code, nil
}

func (s *sTFA) checkTfaUpPhone(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
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

	forbiddentTime := gtime.Now().Add(s.forbiddentTime)
	g.Log().Notice(ctx, "checkTFAUpPhone check phoneUpTime:",
		"tfaInfo:", tfaInfo,
		"info.PhoneUpdatedAt:", tfaInfo.PhoneUpdatedAt.Local(),
		"beforAt:", forbiddentTime,
	)
	if tfaInfo.PhoneUpdatedAt.Before(forbiddentTime) {
		return mpccode.RiskCodeNeedVerification, nil
	}
	return mpccode.RiskCodeForbidden, nil
}

func (s *sTFA) checkTfaUpMail(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
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
	forbiddentTime := gtime.Now().Add(s.forbiddentTime)
	g.Log().Notice(ctx, "checkTfaUpMail check phoneUpTime:",
		"tfaInfo:", tfaInfo,
		"info.MailUpdatedAt:", tfaInfo.MailUpdatedAt,
		"beforAt:", forbiddentTime,
	)
	if tfaInfo.MailUpdatedAt.Before(forbiddentTime) {
		return mpccode.RiskCodeNeedVerification, nil
	}
	return mpccode.RiskCodeForbidden, nil
}

func (s *sTFA) checkTfaBindPhone(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
	return mpccode.RiskCodePass, nil
}
func (s *sTFA) checkTfaBindMail(ctx context.Context, tfaInfo *entity.Tfa) (int32, error) {
	return mpccode.RiskCodePass, nil
}

// /
func (s *sTFA) RiskTfaTidy(ctx context.Context, tfaInfo *entity.Tfa, riskKind model.RiskKind) (string, []string, error) {
	///
	vlist := []string{}
	riskSerial := rand.GenNewSid()

	switch riskKind {
	case model.RiskKind_BindPhone:
		risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_BindPhone)
		verifier := newVerifierPhone(model.RiskKind_BindPhone, "")
		risk.AddVerifier(verifier)
		vlist = append(vlist, "phone")
		if tfaInfo.Mail != "" {
			verifier := newVerifierMail(model.RiskKind_BindPhone, tfaInfo.Mail)
			risk.AddVerifier(verifier)
			vlist = append(vlist, "mail")
		}
	case model.RiskKind_BindMail:
		risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_BindMail)
		verifier := newVerifierMail(model.RiskKind_BindMail, "")
		risk.AddVerifier(verifier)
		vlist = append(vlist, "mail")
		if tfaInfo.Phone != "" {
			verifier := newVerifierPhone(model.RiskKind_BindMail, tfaInfo.Phone)
			risk.AddVerifier(verifier)
			vlist = append(vlist, "phone")
		}
	case model.RiskKind_UpMail:
		risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_UpMail)
		verifier := newVerifierMail(model.RiskKind_UpMail, "")
		risk.AddVerifier(verifier)
		vlist = append(vlist, "mail")
		if tfaInfo.Phone != "" {
			verifier := newVerifierPhone(model.RiskKind_UpMail, tfaInfo.Phone)
			risk.AddVerifier(verifier)
			vlist = append(vlist, "phone")
		}
	case model.RiskKind_UpPhone:
		risk := s.riskPenddingContainer.NewRiskPendding(tfaInfo.UserId, riskSerial, model.RiskKind_UpPhone)
		verifier := newVerifierPhone(model.RiskKind_UpPhone, "")
		risk.AddVerifier(verifier)
		vlist = append(vlist, "phone")
		if tfaInfo.Mail != "" {
			verifier := newVerifierMail(model.RiskKind_UpPhone, tfaInfo.Mail)
			risk.AddVerifier(verifier)
			vlist = append(vlist, "mail")
		}
	}
	///
	return riskSerial, vlist, nil
}
