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

// /
type verifierPhone struct {
	ctx        context.Context
	riskKind   tfaconst.RISKKIND
	verifyKind tfaconst.VERIFYKIND
	code       string
	phone      string
	verified   bool
}

func NewVerifierPhone(riskKind tfaconst.RISKKIND, phone string) tfaconst.IVerifier {
	return &verifierPhone{
		ctx:        gctx.GetInitCtx(),
		riskKind:   riskKind,
		phone:      phone,
		verifyKind: tfaconst.VerifierKind_Phone,
	}
}
func (s *verifierPhone) Destination() string {
	return s.phone
}
func (s *verifierPhone) SendCompletion() error {
	switch s.riskKind {
	case tfaconst.RiskKind_Tx:
	case tfaconst.RiskKind_BindPhone:
		return service.SmsCode().SendBindingCompletionPhone(s.ctx, s.phone)
	case tfaconst.RiskKind_BindMail:
	case tfaconst.RiskKind_UpPhone:
		return service.SmsCode().SendUpCompletionPhone(s.ctx, s.phone)
	case tfaconst.RiskKind_UpMail:
	}
	return nil
}
func (s *verifierPhone) SendVerificationCode() (string, error) {
	switch s.riskKind {
	case tfaconst.RiskKind_Tx:
		return service.SmsCode().SendVerificationCode(s.ctx, s.phone)
	case tfaconst.RiskKind_BindPhone:
		return service.SmsCode().SendBindingPhoneCode(s.ctx, s.phone)
	case tfaconst.RiskKind_BindMail:
		return service.SmsCode().SendVerificationCode(s.ctx, s.phone)
	case tfaconst.RiskKind_UpPhone:
		return service.SmsCode().SendUpPhoneCode(s.ctx, s.phone)
	case tfaconst.RiskKind_UpMail:
		return service.SmsCode().SendVerificationCode(s.ctx, s.phone)
	}
	return "", nil
}

func (s *verifierPhone) VerifyKind() tfaconst.VERIFYKIND {
	return tfaconst.VerifierKind_Phone
}
func (s *verifierPhone) RiskKind() tfaconst.RISKKIND {
	return s.riskKind
}
func (s *verifierPhone) IsDone() bool {
	return s.verified
}

func (s *verifierPhone) SetCode(code string) {
	s.code = code
}
func (s *verifierPhone) Verify(ctx context.Context, verifierCode *model.VerifyCode) (tfaconst.RISKKIND, error) {
	if s.code == verifierCode.PhoneCode && verifierCode.PhoneCode != "" {
		s.verified = true
		return "", nil
	} else {
		s.verified = false
		g.Log().Warning(ctx, "verifierPhone:", "codePhone:", s.code, "verifierPhoneCode:", verifierCode.PhoneCode)
		return tfaconst.VerifierKind_Phone, mpccode.CodeRiskVerifyPhoneInvalid()
	}
}
