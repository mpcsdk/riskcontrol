package tfa

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sTFA) TFAUpPhone(ctx context.Context, tfaInfo *entity.Tfa, phone string, riskSerial string) (string, error) {

	/// need verification
	phoneExists := false
	if tfaInfo.Phone == "" {
		phoneExists = false
	} else {
		phoneExists = true
	}
	//
	event := newRiskEventPhone(phone, func(ctx context.Context) error {
		return s.recordPhone(ctx, tfaInfo.UserId, phone, phoneExists)
	})
	s.riskPenddingContainer.Add(tfaInfo.UserId, riskSerial, event)
	//
	///tfa mailif
	if tfaInfo.Mail != "" {
		event := newRiskEventMail(tfaInfo.Mail, nil)
		s.riskPenddingContainer.Add(tfaInfo.UserId, riskSerial, event)
	}
	///
	return riskSerial, gerror.NewCode(consts.CodePerformRiskNeedVerification)

}
