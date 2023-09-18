package tfa

import (
	"context"
	"riskcontral/internal/service"
)

func (s *sTFA) SendPhoneCode(ctx context.Context, token string, riskSerial string) (string, error) {
	_, err := s.TFAInfo(ctx, token)
	if err != nil {
		return "", err
	}
	////
	err = service.Risk().RiskPhoneCode(ctx, riskSerial)
	if err != nil {
	}
	return "", err
}

func (s *sTFA) SendMailOTP(ctx context.Context, token string, riskSerial string) (string, error) {
	_, err := s.TFAInfo(ctx, token)
	if err != nil {
		return "", err
	}
	////
	err = service.Risk().RiskMailCode(ctx, riskSerial)
	if err != nil {
	}
	return "", err
}

func (s *sTFA) VerifyCode(ctx context.Context, token string, riskSerial string, code string) error {
	// 验证验证码
	err := service.Risk().VerifyCode(ctx, riskSerial, code)
	if err != nil {
		//todo:

	}
	////
	if task, ok := s.pendding[token+riskSerial]; ok {
		delete(s.pendding, token+riskSerial)
		task()
		return err
	}
	//todo:
	return nil
}
