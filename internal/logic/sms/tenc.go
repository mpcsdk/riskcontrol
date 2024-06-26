package sms

import (
	"context"
	"riskcontrol/internal/conf"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/rand"
	"github.com/mpcsdk/mpcCommon/sms"
)

type tenc struct {
	tenc *sms.TencSms
	cfg  *conf.SmsForeign
}

// /
func (s *tenc) SendVerificationCode(ctx context.Context, to string) (string, error) {
	// return "456", nil
	code := rand.RandomDigits(6)
	_, resp, err := s.tenc.SendSms(to, s.cfg.Tenc.VerificationTemplateId, code)

	if err != nil {
		g.Log().Error(ctx, "SendVerificationCode:", "to:", to, "resp:", resp, "err:", err)
		return "", err
	}
	g.Log().Notice(ctx, "SendVerificationCode:", to, code, resp)
	return code, err
}

func (s *tenc) SendBindingPhoneCode(ctx context.Context, to string) (string, error) {
	code := rand.RandomDigits(6)
	_, resp, err := s.tenc.SendSms(to, s.cfg.Tenc.BindingVerificationTemplateId, code)
	if err != nil {
		g.Log().Error(ctx, "SendBindingPhoneCode:", "to:", to, "resp:", resp, "err:", err)
		return "", err
	}
	g.Log().Notice(ctx, "SendBindingPhoneCode:", to, code, resp)
	return code, err
}
func (s *tenc) SendBindingCompletionPhone(ctx context.Context, to string) error {
	_, resp, err := s.tenc.SendSms(to, s.cfg.Tenc.BindingCompletionTemplateId, "")

	if err != nil {
		g.Log().Error(ctx, "SendBindingCompletionPhone:", "to:", to, "resp:", resp, "err:", err)
		return err
	}
	g.Log().Notice(ctx, "SendBindingCompletionPhone:", to, resp)
	return err
}

// //
func (s *tenc) SendUpPhoneCode(ctx context.Context, to string) (string, error) {
	code := rand.RandomDigits(6)
	_, resp, err := s.tenc.SendSms(to, s.cfg.Tenc.UpVerificationTemplateId, code)
	if err != nil {
		g.Log().Error(ctx, "SendUpPhoneCode:", "to:", to, "resp:", resp, "err:", err)
		return "", err
	}
	g.Log().Notice(ctx, "SendUpPhoneCode:", to, code, resp)
	return code, err
}

func (s *tenc) SendUpCompletionPhone(ctx context.Context, to string) error {
	_, resp, err := s.tenc.SendSms(to, s.cfg.Tenc.UpVerificationTemplateId, "")
	if err != nil {
		g.Log().Error(ctx, "SendUpCompletionPhone:", "to:", to, "resp:", resp, "err:", err)
		return err
	}
	g.Log().Notice(ctx, "SendUpCompletionPhone:", to, resp)
	return err
}

func newTencForeign() *tenc {
	return &tenc{
		tenc: sms.NewTencSms(
			conf.Config.Sms.Foreign.Tenc.SecretId,
			conf.Config.Sms.Foreign.Tenc.SecretKey,
			conf.Config.Sms.Foreign.Tenc.Endpoint,
			conf.Config.Sms.Foreign.Tenc.SignMethod,
			conf.Config.Sms.Foreign.Tenc.Region,
			conf.Config.Sms.Foreign.Tenc.SmsSdkAppId,
			conf.Config.Sms.Foreign.Tenc.SignName,
			conf.Config.Sms.Foreign.Tenc.VerificationTemplateId,
			conf.Config.Sms.Foreign.Tenc.BindingCompletionTemplateId,
		),
		cfg: conf.Config.Sms.Foreign,
	}
}
