package tfa

import (
	"context"
	"riskcontral/internal/service"
)

func (s *sTFA) SendPhoneCode(ctx context.Context, token string, opt string) (string, error) {
	info, err := s.TFAInfo(ctx, token)
	if err != nil {
		return "", err
	}
	//todo: check info
	code, err := service.SmsCode().SendCode(ctx, info.Phone)
	if err == nil {
		s.pendding[token+opt+"mail"] = func() {
		}
		return code, nil
	}
	return code, err
}
func (s *sTFA) SendMailOTP(ctx context.Context, token string, opt string) (string, error) {
	info, err := s.TFAInfo(ctx, token)
	if err != nil {
		return "", err
	}
	//todo: check info
	code, err := service.MailCode().SendMailCode(ctx, info.Mail)
	if err == nil {
		s.pendding[token+opt+"mail"] = func() {
		}
		return code, nil
	}
	return code, err
}

func (s *sTFA) VerifyCode(ctx context.Context, token string, kind, code string) error {
	// 验证验证码
	var err error
	if task, ok := s.pendding[token+kind+code]; ok {
		delete(s.pendding, token+kind+code)
		task()
		return err
	}
	//todo:
	return nil
}
