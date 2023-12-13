package tfa

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"

	"riskcontral/api/risk/nrpc"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) TfaInfo(ctx context.Context, req *v1.TfaInfoReq) (res *v1.TfaInfoRes, err error) {
	//
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		g.Log().Warning(ctx, "TFAInfo userinfo:", req)
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	if userInfo.UserId == "" {
		g.Log().Warning(ctx, "TFAInfo no userId:", "req:", req, "userInfo:", userInfo)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	///

	tfaInfo, err := c.nrpc.RpcTfaInfo(ctx, &nrpc.TfaInfoReq{
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
