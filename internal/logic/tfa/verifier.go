package tfa

import (
	"context"
	"errors"
	"fmt"
	"riskcontral/internal/model"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// //
type IVerifier interface {
	Verify(verifierCode *model.VerifyCode) (model.RiskKind, error)
	SetCode(string)
	RiskKind() model.RiskKind
	VerifyKind() VerifyKind
	IsDone() bool
	///
	SendVerificationCode() (string, error)
	SendCompletion() error
}

var errRiskKindTx = errors.New("riskKindTx")
var errRiskKindBindPhone = errors.New("riskKindBindPhone")
var errRiskKindBindMail = errors.New("riskKindBindMail")
var errRiskKindUpPhone = errors.New("riskKindUpPhone")
var errRiskKindUpMail = errors.New("riskKindUpMail")

type emptyVerifier struct {
	riskKind model.RiskKind
}

func newEmptyVerifier(riskKind model.RiskKind) IVerifier {
	return &emptyVerifier{
		riskKind: riskKind,
	}
}

func (s *emptyVerifier) SendCompletion() error {
	fmt.Println(s.riskKind)
	switch s.riskKind {
	case model.RiskKind_Tx:
		return errRiskKindTx
	case model.RiskKind_BindPhone:
		return errRiskKindBindPhone
	case model.RiskKind_BindMail:
		return errRiskKindBindMail
	case model.RiskKind_UpPhone:
		return errRiskKindUpPhone
	case model.RiskKind_UpMail:
		return errRiskKindUpMail
	}
	return nil
}
func (s *emptyVerifier) SendVerificationCode() (string, error) {
	fmt.Println(s.riskKind)
	switch s.riskKind {
	case model.RiskKind_Tx:
		return "", errRiskKindTx
	case model.RiskKind_BindPhone:
		return "", errRiskKindBindPhone
	case model.RiskKind_BindMail:
		return "", errRiskKindBindMail
	case model.RiskKind_UpPhone:
		return "", errRiskKindUpPhone
	case model.RiskKind_UpMail:
		return "", errRiskKindUpMail
	}
	return "", nil
}
func (s *emptyVerifier) IsDone() bool {
	return true
}
func (s *emptyVerifier) VerifyKind() VerifyKind {
	return VerifierKind_Nil
}
func (s *emptyVerifier) RiskKind() model.RiskKind {
	return model.RiskKind_Nil
}

func (s *emptyVerifier) SetCode(code string) {
}
func (s *emptyVerifier) Verify(verifierCode *model.VerifyCode) (model.RiskKind, error) {
	return "", nil
}

// /
// /
type verifierPhone struct {
	ctx        context.Context
	riskKind   model.RiskKind
	verifyKind VerifyKind
	code       string
	phone      string
	verified   bool
}

func newVerifierPhone(riskKind model.RiskKind, phone string) IVerifier {
	return &verifierPhone{
		ctx:        gctx.GetInitCtx(),
		riskKind:   riskKind,
		phone:      phone,
		verifyKind: VerifierKind_Phone,
	}
}
func (s *verifierPhone) SendCompletion() error {
	switch s.riskKind {
	case model.RiskKind_Tx:
	case model.RiskKind_BindPhone:
		return service.SmsCode().SendBindingCompletionPhone(s.ctx, s.phone)
	case model.RiskKind_BindMail:
	case model.RiskKind_UpPhone:
		return service.SmsCode().SendUpCompletionPhone(s.ctx, s.phone)
	case model.RiskKind_UpMail:
	}
	return nil
}
func (s *verifierPhone) SendVerificationCode() (string, error) {
	switch s.riskKind {
	case model.RiskKind_Tx:
		return service.SmsCode().SendVerificationCode(s.ctx, s.phone)
	case model.RiskKind_BindPhone:
		return service.SmsCode().SendVerificationCode(s.ctx, s.phone)
		// return service.SmsCode().SendBindingPhoneCode(s.ctx, s.phone)
	case model.RiskKind_BindMail:
		return service.SmsCode().SendVerificationCode(s.ctx, s.phone)
	case model.RiskKind_UpPhone:
		return service.SmsCode().SendUpPhoneCode(s.ctx, s.phone)
	case model.RiskKind_UpMail:
		return service.SmsCode().SendVerificationCode(s.ctx, s.phone)
	}
	return "", nil
}

func (s *verifierPhone) VerifyKind() VerifyKind {
	return VerifierKind_Phone
}
func (s *verifierPhone) RiskKind() model.RiskKind {
	return s.riskKind
}
func (s *verifierPhone) IsDone() bool {
	return s.verified
}

func (s *verifierPhone) SetCode(code string) {
	s.code = code
}
func (s *verifierPhone) Verify(verifierCode *model.VerifyCode) (model.RiskKind, error) {
	if s.code == verifierCode.PhoneCode && verifierCode.PhoneCode != "" {
		s.verified = true
		return "", nil
	} else {
		s.verified = false
		errcode := gerror.WrapCode(mpccode.CodeRiskVerifyPhoneInvalid, mpccode.ErrArg, mpccode.ErrDetails(
			mpccode.ErrDetail("codePhone:", s.code),
			mpccode.ErrDetail("verifierPhoneCode:", verifierCode.PhoneCode),
		))
		return VerifierKind_Phone, errcode
	}
	return "", nil
}

type verifierMail struct {
	ctx        context.Context
	code       string
	riskKind   model.RiskKind
	verifyKind VerifyKind
	mail       string
	verified   bool
}

func newVerifierMail(riskKind model.RiskKind, mail string) IVerifier {
	return &verifierMail{
		ctx:        gctx.GetInitCtx(),
		riskKind:   riskKind,
		verifyKind: VerifierKind_Mail,
		mail:       mail,
	}
}
func (s *verifierMail) SendCompletion() error {
	switch s.riskKind {
	case model.RiskKind_Tx:
	case model.RiskKind_BindPhone:
	case model.RiskKind_BindMail:
		return service.MailCode().SendBindingCompletionMail(s.ctx, s.mail)
	case model.RiskKind_UpPhone:
	case model.RiskKind_UpMail:
		return service.MailCode().SendUpCompletionMail(s.ctx, s.mail)
	}
	return nil
}
func (s *verifierMail) SendVerificationCode() (string, error) {
	switch s.riskKind {
	case model.RiskKind_Tx:
		return service.MailCode().SendVerificationCode(s.ctx, s.mail)
	case model.RiskKind_BindPhone:
		return service.MailCode().SendVerificationCode(s.ctx, s.mail)
	case model.RiskKind_BindMail:
		return service.MailCode().SendBindingMailCode(s.ctx, s.mail)
	case model.RiskKind_UpPhone:
		return service.MailCode().SendVerificationCode(s.ctx, s.mail)
	case model.RiskKind_UpMail:
		return service.MailCode().SendUpMailCode(s.ctx, s.mail)
	}
	return "", nil
}
func (s *verifierMail) Verify(verifierCode *model.VerifyCode) (model.RiskKind, error) {
	if s.code == verifierCode.MailCode && verifierCode.MailCode != "" {
		s.verified = true

		return "", nil
	} else {
		s.verified = false
		errcode := gerror.WrapCode(mpccode.CodeRiskVerifyMailInvalid, mpccode.ErrArg, mpccode.ErrDetails(
			mpccode.ErrDetail("codeMailCode:", s.code),
			mpccode.ErrDetail("verifierMailCodeCode:", verifierCode.MailCode),
		))
		return VerifierKind_Phone, errcode
	}
}
func (s *verifierMail) IsDone() bool {
	return s.verified
}
func (s *verifierMail) VerifyKind() VerifyKind {
	return s.verifyKind
}
func (s *verifierMail) RiskKind() model.RiskKind {
	return s.riskKind
}
func (s *verifierMail) SetCode(code string) {
	s.code = code
}
