package tfa

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/service"
)

func (c *ControllerV1) VerifyCode(ctx context.Context, req *v1.VerifyCodeReq) (res *v1.VerifyCodeRes, err error) {

	//trace
	ctx, span := gtrace.NewSpan(ctx, "VerifyCode")
	defer span.End()
	// ///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}

	for _, v := range req.VerifyReq {
		err = service.TFA().VerifyCode(ctx, userInfo.UserId, v.RiskSerial, v.Code)
		if err != nil {
			g.Log().Warning(ctx, "VerifyCode", req, err)
			return nil, err
		}
	}
	return nil, nil
	//
	// ///
	// return nil, service.TFA().VerifyCode(ctx, userInfo.UserId, req.RiskSerial, req.Code)

	// //trace
	// ctx, span := gtrace.NewSpan(ctx, "VerifyMailCode")
	// defer span.End()
	// //
	// ///
	// userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	// if err != nil {
	// 	g.Log().Warning(ctx, "VerifyMailCode", req, err)
	// 	return nil, gerror.NewCode(consts.CodeTFANotExist)
	// }
	// ///
	// err = service.TFA().VerifyCode(ctx, userInfo.UserId, req.RiskSerial, req.Code)
	// if err != nil {
	// 	g.Log().Warning(ctx, "VerifyMailCode", req, err)
	// 	return nil, err
	// }
	// return nil, err
}
