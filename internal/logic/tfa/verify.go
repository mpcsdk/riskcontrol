package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) SendPhoneCode(ctx context.Context, userId string, riskSerial string) (string, error) {

	_, err := s.TFAInfo(ctx, userId)
	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", userId, riskSerial, err)
		return "", gerror.NewCode(consts.CodeTFANotExist)
	}
	///
	if f, ok := s.sendpendding[userId+riskSerial]; ok {
		//todo:map conflict
		f()
		delete(s.sendpendding, userId+riskSerial)
		return "", nil
	}

	return "", gerror.NewCode(consts.CodeRiskPerformFailed)
}
func (s *sTFA) sendPhoneCode(ctx context.Context, userId, phone, riskSerial string) (string, error) {
	////
	code, err := service.SmsCode().SendCode(ctx, phone)
	g.Log().Debug(ctx, "SendPhoneCode:", userId, phone, riskSerial, code, err)
	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", userId, riskSerial, err, code)
		return "", gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	s.pendding[userId+riskSerial+code] = func() {
		s.recordPhone(ctx, userId, phone)
	}
	return "", nil
}

func (s *sTFA) SendMailOTP(ctx context.Context, userId string, riskSerial string) (string, error) {
	_, err := s.TFAInfo(ctx, userId)
	if err != nil {
		g.Log().Warning(ctx, "SendMailOTP:", userId, riskSerial, err)
		return "", gerror.NewCode(consts.CodeTFANotExist)
	}
	///
	if f, ok := s.sendpendding[userId+riskSerial]; ok {
		//todo:map conflict
		f()
		delete(s.sendpendding, userId+riskSerial)
		return "", nil
	}
	return "", gerror.NewCode(consts.CodeRiskPerformFailed)
}
func (s *sTFA) sendMailOTP(ctx context.Context, userId, mail, riskSerial string) (string, error) {
	////
	code, err := service.MailCode().SendMailCode(ctx, mail)

	if err != nil {
		g.Log().Warning(ctx, "SendPhoneCode:", userId, riskSerial, err, code)
		return "", gerror.NewCode(consts.CodeRiskPerformFailed)
	}
	s.pendding[userId+riskSerial+code] = func() {
		s.recordMail(ctx, userId, mail)
	}
	return "", nil
}

func (s *sTFA) VerifyCode(ctx context.Context, userId string, riskSerial string, code string) error {
	_, err := s.TFAInfo(ctx, userId)
	if err != nil {
		g.Log().Warning(ctx, "VerifyCode:", userId, riskSerial, err)
		return gerror.NewCode(consts.CodeTFANotExist)
	}
	// 验证验证码
	// err = service.Risk().VerifyCode(ctx, riskSerial, code)
	// if err != nil {
	// 	g.Log().Warning(ctx, "VerifyCode:", token, riskSerial, err)
	// 	return gerror.NewCode(consts.CodeRiskPerformFailed)
	// }
	////
	//todo: concurrent conflict
	if task, ok := s.pendding[userId+riskSerial+code]; ok {
		task()
		delete(s.pendding, userId+riskSerial)
		return nil
	}
	return gerror.NewCode(consts.CodeRiskVerifyInvalid)
}
