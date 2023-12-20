package tfa

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"

	"riskcontral/api/risk/nrpc"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) VerifyCode(ctx context.Context, req *v1.VerifyCodeReq) (res *v1.VerifyCodeRes, err error) {
	g.Log().Notice(ctx, "VerifyCode:", "req:", req)
	// ///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, mpccode.CodeTFANotExist()
	}

	///
	_, err = c.nrpc.RpcVerifyCode(ctx, &nrpc.VerifyCodeReq{
		UserId:     userInfo.UserId,
		RiskSerial: req.RiskSerial,
		MailCode:   req.MailCode,
		PhoneCode:  req.PhoneCode,
	})
	if err != nil {
		return nil, err
	}
	return nil, err
}
