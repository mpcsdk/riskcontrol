package sms

import (
	"context"
	"riskcontral/internal/config"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/rand"
	"github.com/mpcsdk/mpcCommon/sms"
)

type tenc struct {
	tenc *sms.TencSms
	cfg  *config.SmsForeign
}

// /
func (s *tenc) SendVerificationCode(ctx context.Context, to string) (string, error) {
	// return "456", nil
	code := rand.RandomDigits(6)
	_, resp, err := s.tenc.SendSms(to, s.cfg.Tenc.VerificationTemplateId, code)

	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("sendmailto", to),
			mpccode.ErrDetail("sendmail resp", resp),
		))
		return "", err
	}
	g.Log().Notice(ctx, "SendVerificationCode:", to, code)
	return code, err
}

func (s *tenc) SendBindingPhoneCode(ctx context.Context, to string) (string, error) {
	code := rand.RandomDigits(6)
	_, resp, err := s.tenc.SendSms(to, s.cfg.Tenc.BindingVerificationTemplateId, code)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("SendBindingPhoneCode", to),
			mpccode.ErrDetail("SendBindingPhoneCode resp", resp),
		))
		return "", err
	}
	g.Log().Notice(ctx, "SendBindingPhoneCode:", to, code)
	return code, err
}
func (s *tenc) SendBindingCompletionPhone(ctx context.Context, to string) error {
	_, resp, err := s.tenc.SendSms(to, s.cfg.Tenc.BindingCompletionTemplateId, "")

	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("SendBindingCompletionPhone", to),
			mpccode.ErrDetail("SendBindingCompletionPhone resp", resp),
		))
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
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("SendUpPhoneCode", to),
			mpccode.ErrDetail("SendUpPhoneCode resp", resp),
		))
		return "", err
	}
	g.Log().Notice(ctx, "SendUpPhoneCode:", to, code)
	return code, err
}

func (s *tenc) SendUpCompletionPhone(ctx context.Context, to string) error {
	_, resp, err := s.tenc.SendSms(to, s.cfg.Tenc.UpVerificationTemplateId, "")
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("SendUpCompletionPhone", to),
			mpccode.ErrDetail("SendUpCompletionPhone resp", resp),
		))
		return err
	}
	g.Log().Notice(ctx, "SendUpCompletionPhone:", to, resp)
	return err
}

func newTencForeign() *tenc {
	return &tenc{
		tenc: sms.NewTencSms(
			config.Config.Sms.Foreign.Tenc.SecretId,
			config.Config.Sms.Foreign.Tenc.SecretKey,
			config.Config.Sms.Foreign.Tenc.Endpoint,
			config.Config.Sms.Foreign.Tenc.SignMethod,
			config.Config.Sms.Foreign.Tenc.Region,
			config.Config.Sms.Foreign.Tenc.SmsSdkAppId,
			config.Config.Sms.Foreign.Tenc.SignName,
			config.Config.Sms.Foreign.Tenc.VerificationTemplateId,
			config.Config.Sms.Foreign.Tenc.BindingCompletionTemplateId,
		),
		cfg: config.Config.Sms.Foreign,
	}
}
