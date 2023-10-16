package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/net/gtrace"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	if err := c.counter(ctx, req.Token); err != nil {
		return nil, err
	}
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	// err = service.Risk().RiskMailCode(ctx, req.RiskSerial)
	_, err = service.TFA().SendMailCode(ctx, info.UserId, req.RiskSerial)
	return nil, err
}
