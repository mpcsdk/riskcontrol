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

func (c *ControllerV1) TFAInfo(ctx context.Context, req *v1.TFAInfoReq) (res *v1.TFAInfoRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "TFAInfo")
	defer span.End()
	if err := c.counter(ctx, req.Token, "TFAInfo"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	//
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		g.Log().Warning(ctx, "TFAInfo userinfo:", req)
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	if userInfo.UserId == "" {
		g.Log().Warning(ctx, "TFAInfo no userId:", "req:", req, "userInfo:", userInfo)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	info, err := service.DB().FetchTfaInfo(ctx, userInfo.UserId)
	if err != nil || info == nil {
		g.Log().Warning(ctx, "TFAInfo no info:", "req:", req, "userInfo:", userInfo)
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	///

	res = &v1.TFAInfoRes{
		Phone:       info.Phone,
		UpPhoneTime: info.PhoneUpdatedAt,
		Mail:        info.Mail,
		UpMailTime:  info.MailUpdatedAt,
	}
	return
}
