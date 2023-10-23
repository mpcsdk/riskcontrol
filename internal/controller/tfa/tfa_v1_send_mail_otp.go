package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	//limit
	if err := c.counter(ctx, req.Token, "SendMailCode"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	if err := c.limitSendVerification(ctx, req.Token, "SendMailCode"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(consts.ErrLimitSendMailCode)
	}
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	// err = service.Risk().RiskMailCode(ctx, req.RiskSerial)
	_, err = service.TFA().SendMailCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {

		g.Log().Errorf(ctx, "%+v", err)
	}
	return nil, nil
}
