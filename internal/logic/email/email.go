package email

import (
	"context"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/exmail"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/rand"
)

type sMailCode struct {

	////
	From                        string
	SecretId                    string
	SecretKey                   string
	VerificationTemplateID      uint64
	BindingCompletionTemplateID uint64
	Subject                     string
	//
	t *exmail.TencMailClient
}

func (s *sMailCode) SendMailCode(ctx context.Context, to string) (string, error) {
	// return "456", nil
	code := rand.RandomDigits(6)
	resp, err := s.t.SendMail(to, code)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("sendmailto", to),
			mpccode.ErrDetail("sendmail resp", resp),
		))
		return "", err
	}
	g.Log().Notice(ctx, "SendMailCode:", to, code)
	return code, err
}
func (s *sMailCode) SendBindingMail(ctx context.Context, to string) error {
	resp, err := s.t.SendBindingMail(to)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("SendBindingMail", to),
			mpccode.ErrDetail("SendBindingMail resp", resp),
		))
		return err
	}
	g.Log().Notice(ctx, "SendBindingMail:", to, resp)
	return err
}

func new() *sMailCode {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()

	s := &sMailCode{
		From:                        cfg.MustGet(ctx, "exemail.From").String(),
		SecretId:                    cfg.MustGet(ctx, "exemail.SecretId").String(),
		SecretKey:                   cfg.MustGet(ctx, "exemail.SecretKey").String(),
		VerificationTemplateID:      cfg.MustGet(ctx, "exemail.VerificationTemplateID").Uint64(),
		BindingCompletionTemplateID: cfg.MustGet(ctx, "exemail.BindingCompletionTemplateID").Uint64(),
		Subject:                     cfg.MustGet(ctx, "exemail.Subject").String(),
	}
	s.t = exmail.NewTencMailClient(s.SecretId, s.SecretKey, s.VerificationTemplateID, s.BindingCompletionTemplateID, s.From, s.Subject)
	return s
}

func init() {
	service.RegisterMailCode(new())
}
