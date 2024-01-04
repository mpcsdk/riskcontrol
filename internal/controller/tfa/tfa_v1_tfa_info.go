package tfa

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"

	"riskcontral/api/riskserver"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) TfaInfo(ctx context.Context, req *v1.TfaInfoReq) (res *v1.TfaInfoRes, err error) {
	//
	g.Log().Notice(ctx, "TfaInfo:", "req:", req)

	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, mpccode.CodeTFANotExist()
	}
	if userInfo.UserId == "" {
		return nil, mpccode.CodeTFANotExist()
	}
	///

	tfaInfo, err := c.nrpc.RpcTfaInfo(ctx, &riskserver.TfaInfoReq{
		UserId: userInfo.UserId,
	})
	if err != nil {
		return nil, err
	}

	///
	res = &v1.TfaInfoRes{
		Phone:       tfaInfo.Phone,
		UpPhoneTime: tfaInfo.UpPhoneTime,
		Mail:        tfaInfo.Mail,
		UpMailTime:  tfaInfo.UpMailTime,
	}
	return
}
