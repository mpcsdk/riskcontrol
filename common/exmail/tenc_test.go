package exmail

import (
	"testing"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

func Test_Tenc_Mail(t *testing.T) {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	From := cfg.MustGet(ctx, "exemail.From").String()
	SecretId := cfg.MustGet(ctx, "exemail.SecretId").String()
	SecretKey := cfg.MustGet(ctx, "exemail.SecretKey").String()
	VerificationTemplateID := cfg.MustGet(ctx, "exemail.VerificationTemplateID").Uint64()
	BindingCompletionTemplateID := cfg.MustGet(ctx, "exemail.BindingCompletionTemplateID").Uint64()
	Subject := cfg.MustGet(ctx, "exemail.Subject").String()
	m := NewTencMailClient(SecretId, SecretKey,
		VerificationTemplateID, BindingCompletionTemplateID,
		From, Subject)
	///
	stat, err := m.SendMail("xinwei.li@mixmarvel.com", "123456")
	if err != nil {
		t.Error(err)
	}

	t.Log(stat)
	////
	stat, err = m.SendBindingMail("xinwei.li@mixmarvel.com")
	if err != nil {
		t.Error(err)
	}

	t.Log(stat)
	////
}
