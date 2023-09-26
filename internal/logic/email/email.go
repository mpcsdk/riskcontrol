package email

import (
	"context"
	"riskcontral/common"
	"riskcontral/common/exmail"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

type sMailCode struct {

	////
	From       string
	SecretId   string
	SecretKey  string
	TemplateID uint64
	Subject    string
	//
	t *exmail.TencMailClient
}

func (s *sMailCode) SendMailCode(ctx context.Context, to string) (string, error) {
	// return "456", nil
	code := common.RandomDigits(6)
	resp, err := s.t.SendMail(to, code)
	g.Log().Debug(ctx, "SendMailCode:", to, code, resp)
	return code, err
}

func new() *sMailCode {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()

	s := &sMailCode{
		From:       cfg.MustGet(ctx, "exemail.From").String(),
		SecretId:   cfg.MustGet(ctx, "exemail.SecretId").String(),
		SecretKey:  cfg.MustGet(ctx, "exemail.SecretKey").String(),
		TemplateID: cfg.MustGet(ctx, "exemail.TemplateID").Uint64(),
		Subject:    cfg.MustGet(ctx, "exemail.Subject").String(),
	}
	s.t = exmail.NewTencMailClient(s.SecretId, s.SecretKey, s.TemplateID, s.From, s.Subject)
	return s
}

func init() {
	service.RegisterMailCode(new())
}
