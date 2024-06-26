package verifier

import (
	"context"
	"riskcontrol/internal/logic/tfa/tfaconst"
	"riskcontrol/internal/model"
	"riskcontrol/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

type verifierMail struct {
	ctx        context.Context
	code       string
	riskKind   tfaconst.RISKKIND
	verifyKind tfaconst.VERIFYKIND
	mail       string
	verified   bool
}

func NewVerifierMail(riskKind tfaconst.RISKKIND, mail string) tfaconst.IVerifier {
	return &verifierMail{
		ctx:        gctx.GetInitCtx(),
		riskKind:   riskKind,
		verifyKind: tfaconst.VerifierKind_Mail,
		mail:       mail,
	}
}
func (s *verifierMail) Destination() string {
	return s.mail
}
func (s *verifierMail) SendCompletion() error {
	switch s.riskKind {
	case tfaconst.RiskKind_Tx:
	case tfaconst.RiskKind_BindPhone:
	case tfaconst.RiskKind_BindMail:
		return service.MailCode().SendBindingCompletionMail(s.ctx, s.mail)
	case tfaconst.RiskKind_UpPhone:
	case tfaconst.RiskKind_UpMail:
		return service.MailCode().SendUpCompletionMail(s.ctx, s.mail)
	}
	return nil
}
func (s *verifierMail) SendVerificationCode() (string, error) {
	switch s.riskKind {
	case tfaconst.RiskKind_Tx:
		return service.MailCode().SendVerificationCode(s.ctx, s.mail)
	case tfaconst.RiskKind_BindPhone:
		return service.MailCode().SendVerificationCode(s.ctx, s.mail)
	case tfaconst.RiskKind_BindMail:
		return service.MailCode().SendBindingMailCode(s.ctx, s.mail)
	case tfaconst.RiskKind_UpPhone:
		return service.MailCode().SendVerificationCode(s.ctx, s.mail)
	case tfaconst.RiskKind_UpMail:
		return service.MailCode().SendUpMailCode(s.ctx, s.mail)
		///todo: risk mail
	case tfaconst.RiskKind_TxNeedVerify:
		return service.MailCode().SendUpMailCode(s.ctx, s.mail)
	}
	return "", nil
}
func (s *verifierMail) Verify(ctx context.Context, verifierCode *model.VerifyCode) (tfaconst.RISKKIND, error) {
	if s.code == verifierCode.MailCode && verifierCode.MailCode != "" {
		s.verified = true

		return "", nil
	} else {
		s.verified = false
		g.Log().Warning(ctx, "verifierMail:", "codeMail:", s.code, "verifierMailCode:", verifierCode.MailCode)
		return tfaconst.VerifierKind_Phone, mpccode.CodeRiskVerifyMailInvalid()
	}
}
func (s *verifierMail) IsDone() bool {
	return s.verified
}
func (s *verifierMail) VerifyKind() tfaconst.VERIFYKIND {
	return s.verifyKind
}
func (s *verifierMail) RiskKind() tfaconst.RISKKIND {
	return s.riskKind
}
func (s *verifierMail) SetCode(code string) {
	s.code = code
}
