package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/model"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {
	//limit
	if err := c.counter(ctx, req.Token, "SendMailCode"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	if err := c.limitSendVerification(ctx, req.Token, "SendMailCode"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeLimitSendMailCode)
	}
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	///
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil || info == nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTokenInvalid)
	}
	tfaInfo, err := service.TFA().TFAInfo(ctx, info.UserId)
	if err != nil || tfaInfo == nil {
		g.Log().Warning(ctx, "SendMailCode:", req, err)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	////
	riskStat := service.Risk().GetRiskStat(ctx, req.RiskSerial)
	if riskStat == nil {
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	///
	switch riskStat.Type {
	case model.Type_TfaBindMail, model.Type_TfaUpdateMail:
		if req.Mail == "" {
			return nil, gerror.NewCode(mpccode.CodeParamInvalid)
		}
		err = service.DB().TfaMailNotExists(ctx, req.Mail)
		if err != nil {
			g.Log().Warning(ctx, "%+v", err)
			return nil, gerror.NewCode(mpccode.CodeTFAMailExists)
		}
		////
		service.TFA().TfaSetMail(ctx, tfaInfo, req.Mail, req.RiskSerial, riskStat.Type)
		///
	case model.Type_TfaBindPhone, model.Type_TfaUpdatePhone:
		service.TFA().TfaSetPhone_mail(ctx, tfaInfo, req.RiskSerial, riskStat.Type)
	default:
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}

	///
	_, err = service.TFA().SendMailCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	return nil, nil
}
