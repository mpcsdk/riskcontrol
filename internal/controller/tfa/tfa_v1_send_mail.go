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
	riskKind, err := service.TFA().TfaRiskKind(ctx, tfaInfo, req.RiskSerial)
	if err != nil {
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	///
	////
	switch riskKind {
	case model.RiskKind_BindMail, model.RiskKind_UpMail:
		if req.Mail == "" {
			return nil, gerror.NewCode(mpccode.CodeParamInvalid)
		}
		err = service.DB().TfaMailNotExists(ctx, req.Mail)
		if err != nil {
			g.Log().Warning(ctx, "%+v", err)
			return nil, gerror.NewCode(mpccode.CodeTFAMailExists)
		}
		////
		service.TFA().TfaSetMail(ctx, tfaInfo, req.Mail, req.RiskSerial, riskKind)
		///
	case model.RiskKind_BindPhone, model.RiskKind_UpPhone:
	default:
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}

	///
	_, err = service.TFA().SendMailCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTFASendMailFailed)
	}
	return nil, nil
}
