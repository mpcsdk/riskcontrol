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
	domestic *sms.Huawei
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
	resp := &sms.HuaweiResp{}
	state := ""
	var err error
	if strings.HasPrefix(receiver, "+86") {
		resp, state, err = s.foreign.SendSms(receiver, code)
	} else {
		resp, state, err = s.domestic.SendSms(receiver, code)
	}
	///
	if err != nil {
		g.Log().Warning(ctx, "sendcode:", err)
		return code, err
	}
	if state != "" {
		g.Log().Warning(ctx, "sendcode:", resp, state)
		return code, errors.New(state)
	}
	if resp.Code != "000000" {
		g.Log().Warning(ctx, "sendcode:", resp, state)
		return code, errors.New(resp.Description)
	}
	return code, nil
}

func new() *sSmsCode {

	return &sSmsCode{
		pool:     grpool.New(10),
		foreign:  newforeign(),
		domestic: newdomestic(),
	}
}

func init() {
	service.RegisterSmsCode(new())
}
