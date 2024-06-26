package email

import (
	"context"
	"riskcontrol/internal/conf"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
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
	tencMail *tencMail
	///
}

func (s *sMailCode) SendVerificationCode(ctx context.Context, to string) (string, error) {
	// return "456", nil
	code := rand.RandomDigits(6)
	err := s.tencMail.SendVerificationCode(ctx, to, s.VerificationTemplateId, code)
	if err != nil {
		return "", err
	}
	g.Log().Notice(ctx, "SendVerificationCode:", "to:", to, "code:", code)
	return code, nil
}

func (s *sMailCode) SendBindingMailCode(ctx context.Context, to string) (string, error) {
	code := rand.RandomDigits(6)
	err := s.tencMail.SendVerificationCode(ctx, to, s.BindingVerificationTemplateId, code)
	if err != nil {
		return "", err
	}
	g.Log().Notice(ctx, "SendBindingMailCode:", to)
	return code, nil
}
func (s *sMailCode) SendBindingCompletionMail(ctx context.Context, to string) error {
	err := s.tencMail.SendCompletionMail(ctx, to, s.BindingCompletionTemplateId)
	if err != nil {

		g.Log().Error(ctx, "SendBindingCompletionMail:", "to:", to, "err:", err)
		return err
	}
	g.Log().Notice(ctx, "SendBindingCompletionMail:", to)
	return nil
}

// //
func (s *sMailCode) SendUpMailCode(ctx context.Context, to string) (string, error) {
	code := rand.RandomDigits(6)
	err := s.tencMail.SendVerificationCode(ctx, to, s.UpVerificationTemplateId, code)
	if err != nil {
		return "", err
	}
	g.Log().Notice(ctx, "SendUpMailCode:", "to:", to, "code:", code)
	return code, nil
}

func (s *sMailCode) SendUpCompletionMail(ctx context.Context, to string) error {
	err := s.tencMail.SendCompletionMail(ctx, to, s.UpCompletionTemplateId)
	if err != nil {

		return err
	}
	g.Log().Notice(ctx, "SendUpCompletionMail:", to)
	return nil
}
func New() *sMailCode {

	s := &sMailCode{
		From:                          conf.Config.ExEmail.From,      //cfg.MustGet(ctx, "exemail.From").String(),
		SecretId:                      conf.Config.ExEmail.SecretId,  //cfg.MustGet(ctx, "exemail.SecretId").String(),
		SecretKey:                     conf.Config.ExEmail.SecretKey, //cfg.MustGet(ctx, "exemail.SecretKey").String(),
		VerificationTemplateId:        uint64(conf.Config.ExEmail.VerificationTemplateId),
		BindingVerificationTemplateId: uint64(conf.Config.ExEmail.BindingVerificationTemplateId),
		BindingCompletionTemplateId:   uint64(conf.Config.ExEmail.BindingCompletionTemplateId),
		UpVerificationTemplateId:      uint64(conf.Config.ExEmail.UpVerificationTemplateId),
		UpCompletionTemplateId:        uint64(conf.Config.ExEmail.UpCompletionTemplateId),
		Subject:                       conf.Config.ExEmail.Subject, //cfg.MustGet(ctx, "exemail.Subject").String(),
		tencMail:                      newTencMail(gctx.GetInitCtx()),
	}
	return s
}
