package tfa

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"

	v1 "riskcontrol/api/tfa/v1"
	"riskcontrol/internal/service"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {

	g.Log().Notice(ctx, "SendMailCode:", "req:", req)
	//limit
	if err := c.limiter.ApiLimit(ctx, req.Token, "SendMailCode"); err != nil {
		return nil, err
	}
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	//
	//
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil || userInfo.UserId == "" {
		g.Log().Warning(ctx, "TFAInfo userinfo:", "req:", req, "err:", err)
		return nil, mpccode.CodeTFANotExist()
	}

	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil || info == nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	////tidy mail cfg
	err = service.TFA().TfaTidyMail(ctx, info.UserId, req.RiskSerial, req.Mail)
	if err != nil {
		return nil, err
	}
	////
	err = service.TFA().SendMailCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
