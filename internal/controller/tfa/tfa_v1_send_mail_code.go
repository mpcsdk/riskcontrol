package tfa

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"

	"riskcontral/api/riskserver"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {

	g.Log().Notice(ctx, "SendMailCode:", "req:", req)
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil || info == nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	////
	_, err = c.nrpc.RpcSendMailCode(ctx, &riskserver.SendMailCodeReq{
		Mail:       req.Mail,
		RiskSerial: req.RiskSerial,
		UserId:     info.UserId,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
