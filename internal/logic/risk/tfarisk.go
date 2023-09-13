package risk

import (
	"context"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
)

func (s *sRisk) checkTfaUpPhone(ctx context.Context, riskName string, riskData interface{}) (bool, error) {
	return false, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *sRisk) checkTfaUpMail(ctx context.Context, riskName string, riskData interface{}) (bool, error) {
	// data := &conrisk.TfaUpMail{}
	// if _, ok := riskData.(*conrisk.TfaUpMail); !ok {
	// 	return false, gerror.NewCode(gcode.CodeInvalidParameter)
	// }
	// data = riskData.(*conrisk.TfaUpMail)
	// /////
	// info, err := service.TFA().TFAInfo(ctx, data.UserId)
	// if err != nil {
	// 	return false, err
	// }

	// befor24h := gtime.Now().Add(BeforH24)
	// if info.MailUpdatedAt.Before(befor24h) {
	// 	return true, nil
	// }
	return false, gerror.NewCode(gcode.CodeNotImplemented)
}
