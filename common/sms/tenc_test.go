package sms

import (
	"testing"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

func Test_Tenc_foreign(t *testing.T) {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	domestic := NewTencSms(
		cfg.MustGet(ctx, "sms.foreign.tenc.SecretId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SecretKey").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.Endpoint").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SignMethod").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.Region").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SmsSdkAppId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SignName").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.VerificationTemplateId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.BindingCompletionTemplateId").String(),
	)
	resp, stat, err := domestic.SendSms("+447862429616", "456712")
	if err != nil {
		t.Error(err)
	}
	if resp == false {
		t.Error(stat)
	}
	t.Log(resp, stat)
	////

}

func Test_Tenc_foreign_binding(t *testing.T) {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	domestic := NewTencSms(
		cfg.MustGet(ctx, "sms.foreign.tenc.SecretId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SecretKey").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.Endpoint").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SignMethod").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.Region").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SmsSdkAppId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SignName").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.VerificationTemplateId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.BindingCompletionTemplateId").String(),
	)

	///
	resp, stat, err := domestic.SendBinding("+447862429616")
	if err != nil {
		t.Error(err)
	}
	if resp == false {
		t.Error(stat)
	}
	t.Log(resp, stat)
	///

}

func Test_Tenc_domestic_incorrect(t *testing.T) {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	domestic := NewTencSms(
		cfg.MustGet(ctx, "sms.foreign.tenc.SecretId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SecretKey").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.Endpoint").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SignMethod").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.Region").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SmsSdkAppId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SignName").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.VerificationTemplateId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.BindingCompletionTemplateId").String(),
	)
	resp, stat, err := domestic.SendSms("+4478624296161", "4567")
	if err != nil {
		t.Error(err)
	}
	if resp == false {
		t.Error(stat)
	}

	t.Log(resp, stat)
	///

}
func Test_Tenc_domestic_xinjiapo(t *testing.T) {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	domestic := NewTencSms(
		cfg.MustGet(ctx, "sms.foreign.tenc.SecretId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SecretKey").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.Endpoint").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SignMethod").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.Region").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SmsSdkAppId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SignName").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.VerificationTemplateId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.BindingCompletionTemplateId").String(),
	)

	resp, stat, err := domestic.SendSms("+659035559", "4567")
	if err != nil {
		t.Error(err)
	}
	if resp == false {
		t.Error(stat)
	}

	t.Log(resp, stat)
}

func Test_Tenc_foreign_xinjiapo2(t *testing.T) {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	domestic := NewTencSms(
		cfg.MustGet(ctx, "sms.foreign.tenc.SecretId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SecretKey").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.Endpoint").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SignMethod").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.Region").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SmsSdkAppId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.SignName").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.VerificationTemplateId").String(),
		cfg.MustGet(ctx, "sms.foreign.tenc.BindingCompletionTemplateId").String(),
	)

	resp, stat, err := domestic.SendSms("+6588606326", "4567")
	if err != nil {
		t.Error(err)
	}
	if resp == false {
		t.Error(stat)
	}

	t.Log(resp, stat)
}
