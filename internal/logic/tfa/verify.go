package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) SendPhoneCode(ctx context.Context, token string, riskSerial string) (string, error) {
	_, err := s.TFAInfo(ctx, token)
	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", token, riskSerial, err)
		return "", gerror.NewCode(consts.CodeTFANotExist)
	}
	////
	err = service.Risk().RiskPhoneCode(ctx, riskSerial)
	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", token, riskSerial, err)
		return "", gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	return "", err
}

func (s *sTFA) SendMailOTP(ctx context.Context, token string, riskSerial string) (string, error) {
	_, err := s.TFAInfo(ctx, token)
	if err != nil {
		g.Log().Warning(ctx, "SendMailOTP:", token, riskSerial, err)
		return "", gerror.NewCode(consts.CodeTFANotExist)
	}
	////
	err = service.Risk().RiskMailCode(ctx, riskSerial)
	if err != nil {
		g.Log().Warning(ctx, "SendMailOTP:", token, riskSerial, err)
		return "", gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	return "", err
}

func (s *sTFA) VerifyCode(ctx context.Context, token string, riskSerial string, code string) error {
	_, err := s.TFAInfo(ctx, token)
	if err != nil {
		g.Log().Warning(ctx, "VerifyCode:", token, riskSerial, err)
		return gerror.NewCode(consts.CodeTFANotExist)
	}
	// 验证验证码
	err = service.Risk().VerifyCode(ctx, riskSerial, code)
	if err != nil {
		g.Log().Warning(ctx, "VerifyCode:", token, riskSerial, err)
		return gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	////
	//todo: concurrent conflict
	if task, ok := s.pendding[token+riskSerial]; ok {
		delete(s.pendding, token+riskSerial)
		task()
	}
	return nil
}
