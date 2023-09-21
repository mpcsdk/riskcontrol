package sms

import (
	"testing"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

func Test_Tenc_domestic(t *testing.T) {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	domestic := NewTencSms(
		cfg.MustGet(ctx, "sms.tenc.domestic.SecretId").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SecretKey").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.Endpoint").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SignMethod").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.Region").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SmsSdkAppId").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SignName").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.TemplateId").String(),
	)
	resp, stat, err := domestic.SendSms("+447862429616", "4567")
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
		cfg.MustGet(ctx, "sms.tenc.domestic.SecretId").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SecretKey").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.Endpoint").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SignMethod").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.Region").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SmsSdkAppId").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SignName").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.TemplateId").String(),
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
		cfg.MustGet(ctx, "sms.tenc.domestic.SecretId").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SecretKey").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.Endpoint").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SignMethod").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.Region").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SmsSdkAppId").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SignName").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.TemplateId").String(),
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

func Test_Tenc_domestic_xinjiapo2(t *testing.T) {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	domestic := NewTencSms(
		cfg.MustGet(ctx, "sms.tenc.domestic.SecretId").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SecretKey").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.Endpoint").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SignMethod").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.Region").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SmsSdkAppId").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SignName").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.TemplateId").String(),
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
