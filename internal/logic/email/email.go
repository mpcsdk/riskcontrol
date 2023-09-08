package email

import (
	"context"
	"crypto/tls"
	"riskcontral/utility/common"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"gopkg.in/gomail.v2"
)

type sMailCode struct {
	From    string
	Passwd  string
	Host    string
	Port    int
	Subject string
	Body    string
	////

	d *gomail.Dialer
}

func (s *sMailCode) SendMailCode(ctx context.Context, to string, code string) (string, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", s.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "验证码")
	m.SetBody("text/html", s.Body+code)
	return common.GenNewSid(), s.d.DialAndSend(m)

}

func new() *sMailCode {
	cfg := gcfg.Instance()
	ctx := gctx.GetInitCtx()

	s := &sMailCode{
		From:    cfg.MustGet(ctx, "emailOTP.Mail").String(),
		Passwd:  cfg.MustGet(ctx, "emailOTP.Password").String(),
		Host:    cfg.MustGet(ctx, "emailOTP.Host").String(),
		Port:    cfg.MustGet(ctx, "emailOTP.Port").Int(),
		Subject: cfg.MustGet(ctx, "emailOTP.Subject").String(),
		Body:    cfg.MustGet(ctx, "emailOTP.Body").String(),
	}
	d := gomail.NewDialer(s.Host, s.Port, s.From, s.Passwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	s.d = d
	return s
}
