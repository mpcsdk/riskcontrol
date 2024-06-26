package email

import (
	"context"
	"riskcontrol/internal/service"
	"testing"
)

func Test_SendMail(t *testing.T) {
	service.RegisterMailCode(New())
	_, err := service.MailCode().SendVerificationCode(context.Background(), "xinwei.li@mixmarvel.com")
	if err != nil {
		t.Error(err)
	}
}
