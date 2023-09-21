package sms

import (
	"context"
	"errors"
	"riskcontral/common"
	"riskcontral/common/sms"
	"riskcontral/internal/service"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
)

type sSmsCode struct {
	// domestic *sms.Huawei
	domestic *sms.TencSms
	foreign  *sms.Huawei
	pool     *grpool.Pool
}

func newforeign() *sms.Huawei {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	return &sms.Huawei{
		APIAddress:        cfg.MustGet(ctx, "sms.foreign.APIAddress").String(),
		ApplicationKey:    cfg.MustGet(ctx, "sms.foreign.ApplicationKey").String(),
		ApplicationSecret: cfg.MustGet(ctx, "sms.foreign.ApplicationSecret").String(),
		Sender:            cfg.MustGet(ctx, "sms.foreign.Sender").String(),
		TemplateID:        cfg.MustGet(ctx, "sms.foreign.TemplateID").String(),
		Signature:         cfg.MustGet(ctx, "sms.foreign.Signature").String(),
	}
}
func newdomestic() *sms.Huawei {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	return &sms.Huawei{
		APIAddress:        cfg.MustGet(ctx, "sms.domestic.APIAddress").String(),
		ApplicationKey:    cfg.MustGet(ctx, "sms.domestic.ApplicationKey").String(),
		ApplicationSecret: cfg.MustGet(ctx, "sms.domestic.ApplicationSecret").String(),
		Sender:            cfg.MustGet(ctx, "sms.domestic.Sender").String(),
		TemplateID:        cfg.MustGet(ctx, "sms.domestic.TemplateID").String(),
		Signature:         cfg.MustGet(ctx, "sms.domestic.Signature").String(),
	}
}
func newTencDomestic() *sms.TencSms {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()
	return sms.NewTencSms(
		cfg.MustGet(ctx, "sms.tenc.domestic.SecretId").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SecretKey").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.Endpoint").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SignMethod").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.Region").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SmsSdkAppId").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.SignName").String(),
		cfg.MustGet(ctx, "sms.tenc.domestic.TemplateId").String(),
	)
}

// //
// //
func (s *sSmsCode) sendCode(ctx context.Context, receiver, code string) error {
	//todo: dstphone
	resp, status, err := s.foreign.SendSms(receiver, code)
	g.Log().Info(ctx, "sendcode:", resp, status, err)
	///
	return err
}

func (s *sSmsCode) SendCode(ctx context.Context, receiver string) (string, error) {
	code := common.RandomDigits(6)
	ok := false
	state := ""
	var err error
	if strings.HasPrefix(receiver, "+86") {
		ok, state, err = s.foreign.SendSms(receiver, code)
	} else {
		// resp, state, err = s.domestic.SendSms(receiver, code)
		ok, state, err = s.domestic.SendSms(receiver, code)
	}
	///
	if err != nil {
		g.Log().Warning(ctx, "sendcode:", err)
		return code, err
	}
	if ok != true {
		g.Log().Warning(ctx, "sendcode:", ok, state)
		return code, errors.New(state)
	}

	return code, nil
}

func new() *sSmsCode {
	return &sSmsCode{
		pool:     grpool.New(10),
		foreign:  newforeign(),
		domestic: newTencDomestic(),
	}
}

func init() {
	service.RegisterSmsCode(new())
}
