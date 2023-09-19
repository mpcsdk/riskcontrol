package sms

import (
	"testing"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

func Test_foreign(t *testing.T) {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	foreign := &Huawei{
		APIAddress:        cfg.MustGet(ctx, "sms.foreign.APIAddress").String(),
		ApplicationKey:    cfg.MustGet(ctx, "sms.foreign.ApplicationKey").String(),
		ApplicationSecret: cfg.MustGet(ctx, "sms.foreign.ApplicationSecret").String(),
		Sender:            cfg.MustGet(ctx, "sms.foreign.Sender").String(),
		TemplateID:        cfg.MustGet(ctx, "sms.foreign.TemplateID").String(),
		Signature:         cfg.MustGet(ctx, "sms.foreign.Signature").String(),
	}
	resp, stat, err := foreign.SendSms("+8615111226175", "123456")
	if err != nil {
		t.Error(err)
	}
	if stat != "" {
		t.Log(stat)
		t.Error(err)
	}
	t.Log(resp)
}
func Test_domestic(t *testing.T) {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	domestic := &Huawei{
		APIAddress:        cfg.MustGet(ctx, "sms.domestic.APIAddress").String(),
		ApplicationKey:    cfg.MustGet(ctx, "sms.domestic.ApplicationKey").String(),
		ApplicationSecret: cfg.MustGet(ctx, "sms.domestic.ApplicationSecret").String(),
		Sender:            cfg.MustGet(ctx, "sms.domestic.Sender").String(),
		TemplateID:        cfg.MustGet(ctx, "sms.domestic.TemplateID").String(),
		Signature:         cfg.MustGet(ctx, "sms.domestic.Signature").String(),
	}
	resp, stat, err := domestic.SendSms("447862429616", "4567")
	if err != nil {
		t.Error(err)
	}
	if stat != "" {
		t.Log(stat)
		t.Error(err)
	}
	t.Log(resp)
}
