package email

import (
	"context"
	"riskcontral/internal/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/exmail"
)

type tencMail struct {
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
}

func newTencMail(ctx context.Context) *tencMail {
	s := &tencMail{
		From:                          config.Config.ExEmail.From,      //cfg.MustGet(ctx, "exemail.From").String(),
		SecretId:                      config.Config.ExEmail.SecretId,  //cfg.MustGet(ctx, "exemail.SecretId").String(),
		SecretKey:                     config.Config.ExEmail.SecretKey, //cfg.MustGet(ctx, "exemail.SecretKey").String(),
		VerificationTemplateId:        uint64(config.Config.ExEmail.VerificationTemplateId),
		BindingVerificationTemplateId: uint64(config.Config.ExEmail.BindingVerificationTemplateId),
		BindingCompletionTemplateId:   uint64(config.Config.ExEmail.BindingCompletionTemplateId),
		UpVerificationTemplateId:      uint64(config.Config.ExEmail.UpVerificationTemplateId),
		UpCompletionTemplateId:        uint64(config.Config.ExEmail.UpCompletionTemplateId),
		Subject:                       config.Config.ExEmail.Subject, //cfg.MustGet(ctx, "exemail.Subject").String(),
	}
	s.t = exmail.NewTencMailClient(s.SecretId, s.SecretKey,
		s.From, s.Subject)
	return s
}
func (s *tencMail) SendVerificationCode(ctx context.Context, to string, templatedId uint64, code string) error {
	resp, err := s.t.SendVerificationCode(to, templatedId, code)
	if err != nil {
		g.Log().Error(ctx, "SendVerificationCode:", "to:", to, "resp:", resp, "err:", err)
		return err
	}
	g.Log().Notice(ctx, "SendVerificationCode:", to, resp)
	return nil
}
func (s *tencMail) SendCompletionMail(ctx context.Context, to string, templatedId uint64) error {
	resp, err := s.t.SendCompletion(to, templatedId)
	if err != nil {
		g.Log().Error(ctx, "SendBindingCompletionMail:", "to:", to, "resp:", resp, "err:", err)
		return err
	}
	g.Log().Notice(ctx, "SendBindingCompletionMail:", to, resp)
	return nil
}
