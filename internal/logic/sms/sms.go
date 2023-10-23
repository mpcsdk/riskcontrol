package sms

import (
	"context"
	"riskcontral/internal/config"
	"riskcontral/internal/service"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/rand"
	"github.com/mpcsdk/mpcCommon/sms"
)

type sSmsCode struct {
	// domestic *sms.Huawei
	foreign  *sms.TencSms
	domestic *sms.Huawei
	pool     *grpool.Pool
}

// func newforeign() *sms.Huawei {
// 	// cfg := gcfg.Instance()
// 	// ctx := gctx.GetInitCtx()
// 	// return &sms.Huawei{
// 	// 	APIAddress:        cfg.MustGet(ctx, "sms.foreign.huawei.APIAddress").String(),
// 	// 	ApplicationKey:    cfg.MustGet(ctx, "sms.foreign.huawei.ApplicationKey").String(),
// 	// 	ApplicationSecret: cfg.MustGet(ctx, "sms.foreign.huawei.ApplicationSecret").String(),
// 	// 	Sender:            cfg.MustGet(ctx, "sms.foreign.huawei.Sender").String(),
// 	// 	TemplateID:        cfg.MustGet(ctx, "sms.foreign.huawei.TemplateID").String(),
// 	// 	Signature:         cfg.MustGet(ctx, "sms.foreign.huawei.Signature").String(),
// 	// }

// }

func newdomestic() *sms.Huawei {
	// cfg := gcfg.Instance()
	// ctx := gctx.GetInitCtx()
	// return &sms.Huawei{
	// 	APIAddress:        cfg.MustGet(ctx, "sms.domestic.huawei.APIAddress").String(),
	// 	ApplicationKey:    cfg.MustGet(ctx, "sms.domestic.huawei.ApplicationKey").String(),
	// 	ApplicationSecret: cfg.MustGet(ctx, "sms.domestic.huawei.ApplicationSecret").String(),
	// 	Sender:            cfg.MustGet(ctx, "sms.domestic.huawei.Sender").String(),
	// 	TemplateID:        cfg.MustGet(ctx, "sms.domestic.huawei.TemplateID").String(),
	// 	Signature:         cfg.MustGet(ctx, "sms.domestic.huawei.Signature").String(),
	// }
	return &sms.Huawei{
		APIAddress:        config.Config.Sms.Domestic.Huawei.APIAddress,
		ApplicationKey:    config.Config.Sms.Domestic.Huawei.ApplicationKey,
		ApplicationSecret: config.Config.Sms.Domestic.Huawei.ApplicationSecret,
		Sender:            config.Config.Sms.Domestic.Huawei.Sender,
		TemplateID:        config.Config.Sms.Domestic.Huawei.TemplateID,
		Signature:         config.Config.Sms.Domestic.Huawei.Signature,
	}
}
func newTencForeign() *sms.TencSms {
	// cfg := gcfg.Instance()
	// ctx := gctx.GetInitCtx()
	// return sms.NewTencSms(
	// 	cfg.MustGet(ctx, "sms.foreign.tenc.SecretId").String(),
	// 	cfg.MustGet(ctx, "sms.foreign.tenc.SecretKey").String(),
	// 	cfg.MustGet(ctx, "sms.foreign.tenc.Endpoint").String(),
	// 	cfg.MustGet(ctx, "sms.foreign.tenc.SignMethod").String(),
	// 	cfg.MustGet(ctx, "sms.foreign.tenc.Region").String(),
	// 	cfg.MustGet(ctx, "sms.foreign.tenc.SmsSdkAppId").String(),
	// 	cfg.MustGet(ctx, "sms.foreign.tenc.SignName").String(),
	// 	cfg.MustGet(ctx, "sms.foreign.tenc.VerificationTemplateId").String(),
	// 	cfg.MustGet(ctx, "sms.foreign.tenc.BindingCompletionTemplateId").String(),
	// )
	return sms.NewTencSms(
		config.Config.Sms.Foreign.Tenc.SecretId,
		config.Config.Sms.Foreign.Tenc.SecretKey,
		config.Config.Sms.Foreign.Tenc.Endpoint,
		config.Config.Sms.Foreign.Tenc.SignMethod,
		config.Config.Sms.Foreign.Tenc.Region,
		config.Config.Sms.Foreign.Tenc.SmsSdkAppId,
		config.Config.Sms.Foreign.Tenc.SignName,
		config.Config.Sms.Foreign.Tenc.VerificationTemplateId,
		config.Config.Sms.Foreign.Tenc.BindingCompletionTemplateId,
	)
}

// //
// //
func (s *sSmsCode) sendCode(ctx context.Context, receiver, code string) error {
	//todo: dstphone
	resp, status, err := s.foreign.SendSms(receiver, code)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("resp", resp),
			mpccode.ErrDetail("status", status),
		))
		return err
	}
	///
	return nil
}

func (s *sSmsCode) SendCode(ctx context.Context, receiver string) (string, error) {
	// return "123", nil
	code := rand.RandomDigits(6)
	ok := false
	state := ""
	var err error
	if strings.HasPrefix(receiver, "+86") {
		ok, state, err = s.domestic.SendSms(receiver, code)
	} else {
		// resp, state, err = s.domestic.SendSms(receiver, code)
		ok, state, err = s.foreign.SendSms(receiver, code)
	}
	///
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("stat", state),
		))
		return code, err
	}
	if ok != true {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("stat", state),
		))
		return code, err
	}

	return code, nil
}

func new() *sSmsCode {
	return &sSmsCode{
		pool:     grpool.New(10),
		foreign:  newTencForeign(),
		domestic: newdomestic(),
	}
}

func init() {
	service.RegisterSmsCode(new())
}
