package tfa

import (
	"context"
	"riskcontral/api/riskserver"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error) {
	g.Log().Notice(ctx, "SendSmsCode:", "req:", req)

	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil || info == nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	////
	_, err = c.nrpc.RpcSendPhoneCode(ctx, &riskserver.SendPhoneCodeReq{
		Phone:      req.Phone,
		RiskSerial: req.RiskSerial,
		UserId:     info.UserId,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
