package sms

import (
	"context"
	"riskcontrol/internal/conf"
	"strings"
)

type sSmsCode struct {
	foreign  *tenc
	domestic *huawei

	foreignCfg  *conf.SmsForeign
	domesticCfg *conf.SmsDomestic
}

// }

// /
func (s *sSmsCode) SendVerificationCode(ctx context.Context, to string) (string, error) {
	var err error
	code := ""
	if strings.HasPrefix(to, "+86") {
		code, err = s.domestic.SendVerificationCode(ctx, to)
	} else {
		// resp, state, err = s.domestic.SendSms(receiver, code)
		code, err = s.foreign.SendVerificationCode(ctx, to)
	}
	return code, err
}

func (s *sSmsCode) SendBindingPhoneCode(ctx context.Context, to string) (string, error) {
	var err error
	code := ""
	if strings.HasPrefix(to, "+86") {
		code, err = s.domestic.SendBindingPhoneCode(ctx, to)
	} else {
		// resp, state, err = s.domestic.SendSms(receiver, code)
		code, err = s.foreign.SendBindingPhoneCode(ctx, to)
	}
	return code, err
}
func (s *sSmsCode) SendBindingCompletionPhone(ctx context.Context, to string) error {
	var err error
	if strings.HasPrefix(to, "+86") {
		err = s.domestic.SendBindingCompletionPhone(ctx, to)
	} else {
		err = s.foreign.SendBindingCompletionPhone(ctx, to)
	}

	return err
}

// //
func (s *sSmsCode) SendUpPhoneCode(ctx context.Context, to string) (string, error) {
	var err error
	code := ""
	if strings.HasPrefix(to, "+86") {
		code, err = s.domestic.SendUpPhoneCode(ctx, to)
	} else {
		// resp, state, err = s.domestic.SendSms(receiver, code)
		code, err = s.foreign.SendUpPhoneCode(ctx, to)
	}
	return code, err
}

func (s *sSmsCode) SendUpCompletionPhone(ctx context.Context, receiver string) error {
	var err error
	if strings.HasPrefix(receiver, "+86") {
		err = s.domestic.SendUpCompletionPhone(ctx, receiver)
	} else {
		// resp, state, err = s.domestic.SendSms(receiver, code)
		err = s.foreign.SendUpCompletionPhone(ctx, receiver)
	}

	return err
}
func New() *sSmsCode {
	return &sSmsCode{
		foreign:     newTencForeign(),
		domestic:    newdomestic(),
		foreignCfg:  conf.Config.Sms.Foreign,
		domesticCfg: conf.Config.Sms.Domestic,
	}
}
