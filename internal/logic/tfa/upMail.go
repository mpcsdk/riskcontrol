package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTFA) TFAUpMail(ctx context.Context, tfaInfo *entity.Tfa, mail string, riskSerial string) (string, error) {
	g.Log().Debug(ctx, "TFAUpMail:", tfaInfo, mail)
	//
	////
	mailExists := false
	if tfaInfo.Mail == "" {
		mailExists = false
	} else {
		mailExists = true
	}
	event := newRiskEventMail(mail, func(ctx context.Context) error {
		err := s.recordMail(ctx, tfaInfo.UserId, mail, mailExists)
		if err != nil {
			return err
		}
		return service.MailCode().SendBindingMail(ctx, mail)
	})
	s.riskPenddingContainer.Add(tfaInfo.UserId, riskSerial, event)
	///tfa phone if
	if tfaInfo.Phone != "" {
		event := newRiskEventPhone(tfaInfo.Phone, nil)
		// s.addRiskEvent(ctx, userId, riskSerial, event)
		s.riskPenddingContainer.Add(tfaInfo.UserId, riskSerial, event)
	}
	//
	return riskSerial, gerror.NewCode(consts.CodePerformRiskNeedVerification)
	//
}
