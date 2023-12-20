package sms

import (
	"context"
	"riskcontral/internal/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/rand"
	"github.com/mpcsdk/mpcCommon/sms"
)

type huawei struct {
	huawei *sms.Huawei
	cfg    *config.SmsDomestic
}

// /
func (s *huawei) SendVerificationCode(ctx context.Context, to string) (string, error) {
	// return "456", nil
	code := rand.RandomDigits(6)
	ok, resp, err := s.huawei.SendSms(to, s.cfg.Huawei.VerificationTemplateId, code)
	if !ok {
		g.Log().Error(ctx, "SendVerificationCode:", "to:", to, "resp:", resp, "err:", err)
		return "", err
	}
	if err != nil {
		g.Log().Error(ctx, "SendVerificationCode:", "to:", to, "resp:", resp, "err:", err)
		return "", err
	}
	g.Log().Notice(ctx, "SendVerificationCode:", to, code, resp)
	return code, err
}

func (s *huawei) SendBindingPhoneCode(ctx context.Context, to string) (string, error) {
	code := rand.RandomDigits(6)
	_, resp, err := s.huawei.SendSms(to, s.cfg.Huawei.BindingVerificationTemplateId, code)
	if err != nil {
		g.Log().Error(ctx, "SendBindingPhoneCode:", "to:", to, "resp:", resp, "err:", err)
		return "", err
	}
	g.Log().Notice(ctx, "SendBindingPhoneCode:", to, code, resp)
	return code, err
}
func (s *huawei) SendBindingCompletionPhone(ctx context.Context, to string) error {
	_, resp, err := s.huawei.SendSms(to, s.cfg.Huawei.BindingCompletionTemplateId, "")
	if err != nil {
		g.Log().Error(ctx, "SendBindingCompletionPhone:", "to:", to, "resp:", resp, "err:", err)
		return err
	}
	g.Log().Notice(ctx, "SendBindingCompletionPhone:", to, resp)
	return err
}

// //
func (s *huawei) SendUpPhoneCode(ctx context.Context, to string) (string, error) {
	code := rand.RandomDigits(6)
	_, resp, err := s.huawei.SendSms(to, s.cfg.Huawei.UpVerificationTemplateId, code)
	if err != nil {
		g.Log().Error(ctx, "SendUpPhoneCode:", "to:", to, "resp:", resp, "err:", err)
		return "", err
	}
	g.Log().Notice(ctx, "SendUpPhoneCode:", to, code, resp)
	return code, err
}

func (s *huawei) SendUpCompletionPhone(ctx context.Context, to string) error {
	_, resp, err := s.huawei.SendSms(to, s.cfg.Huawei.UpCompletionTemplateId, "")
	if err != nil {
		g.Log().Error(ctx, "SendUpCompletionPhone:", "to:", to, "resp:", resp, "err:", err)
		return err
	}
	g.Log().Notice(ctx, "SendUpCompletionPhone:", to, resp)
	return err
}
func newdomestic() *huawei {
	return &huawei{
		huawei: &sms.Huawei{
			APIAddress:        config.Config.Sms.Domestic.Huawei.APIAddress,
			ApplicationKey:    config.Config.Sms.Domestic.Huawei.ApplicationKey,
			ApplicationSecret: config.Config.Sms.Domestic.Huawei.ApplicationSecret,
			Sender:            config.Config.Sms.Domestic.Huawei.Sender,
			SenderCompletion:  config.Config.Sms.Domestic.Huawei.SenderCompletion,
			TemplateID:        config.Config.Sms.Domestic.Huawei.VerificationTemplateId,
			Signature:         config.Config.Sms.Domestic.Huawei.Signature,
		},
		cfg: config.Config.Sms.Domestic,
	}
}
