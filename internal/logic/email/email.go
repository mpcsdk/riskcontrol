package email

import (
	"context"
	"riskcontral/internal/config"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/exmail"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/rand"
)

type sMailCode struct {

	////
	From      string
	SecretId  string
	SecretKey string
	Subject   string
	///
	VerificationTemplateId        uint64
	BindingVerificationTemplateId uint64
	BindingCompletionTemplateId   uint64
	UpVerificationTemplateId      uint64
	UpCompletionTemplateId        uint64
	//
	t *exmail.TencMailClient
	///
}

func (s *sMailCode) SendVerificationCode(ctx context.Context, to string) (string, error) {
	// return "456", nil
	code := rand.RandomDigits(6)
	resp, err := s.t.SendVerificationCode(to, s.VerificationTemplateId, code)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("sendmailto", to),
			mpccode.ErrDetail("sendmail resp", resp),
		))
		return "", err
	}
	g.Log().Notice(ctx, "SendVerificationCode:", to, resp)
	return code, err
}

func (s *sMailCode) SendBindingMailCode(ctx context.Context, to string) (string, error) {
	code := rand.RandomDigits(6)
	resp, err := s.t.SendVerificationCode(to, s.BindingVerificationTemplateId, code)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("SendBindingMailCode", to),
			mpccode.ErrDetail("SendBindingMailCode resp", resp),
		))
		return "", err
	}
	g.Log().Notice(ctx, "SendBindingMailCode:", to, resp)
	return code, err
}
func (s *sMailCode) SendBindingCompletionMail(ctx context.Context, to string) error {
	resp, err := s.t.SendCompletion(to, s.BindingCompletionTemplateId)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("SendBindingCompletionMail", to),
			mpccode.ErrDetail("SendBindingCompletionMail resp", resp),
		))
		return err
	}
	g.Log().Notice(ctx, "SendBindingCompletionMail:", to, resp)
	return err
}

// //
func (s *sMailCode) SendUpMailCode(ctx context.Context, to string) (string, error) {
	code := rand.RandomDigits(6)
	resp, err := s.t.SendVerificationCode(to, s.UpVerificationTemplateId, code)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("SendUpMailCode", to),
			mpccode.ErrDetail("SendUpMailCode resp", resp),
		))
		return "", err
	}
	g.Log().Notice(ctx, "SendUpMailCode:", to, resp)
	return code, err
}

func (s *sMailCode) SendUpCompletionMail(ctx context.Context, to string) error {
	resp, err := s.t.SendCompletion(to, s.UpCompletionTemplateId)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("SendUpCompletionMail", to),
			mpccode.ErrDetail("SendUpCompletionMail resp", resp),
		))
		return err
	}
	g.Log().Notice(ctx, "SendUpCompletionMail:", to, resp)
	return err
}
func new() *sMailCode {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()

	s := &sMailCode{
		From:                          cfg.MustGet(ctx, "exemail.From").String(),
		SecretId:                      cfg.MustGet(ctx, "exemail.SecretId").String(),
		SecretKey:                     cfg.MustGet(ctx, "exemail.SecretKey").String(),
		VerificationTemplateId:        uint64(config.Config.ExEmail.VerificationTemplateId),
		BindingVerificationTemplateId: uint64(config.Config.ExEmail.BindingVerificationTemplateId),
		BindingCompletionTemplateId:   uint64(config.Config.ExEmail.BindingCompletionTemplateId),
		UpVerificationTemplateId:      uint64(config.Config.ExEmail.UpVerificationTemplateId),
		UpCompletionTemplateId:        uint64(config.Config.ExEmail.UpCompletionTemplateId),
		Subject:                       cfg.MustGet(ctx, "exemail.Subject").String(),
	}
	s.t = exmail.NewTencMailClient(s.SecretId, s.SecretKey,
		s.From, s.Subject)
	return s
}

func init() {
	service.RegisterMailCode(new())
}
