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
		return nil, err
	}
	//
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		g.Log().Warning(ctx, "TFAInfo:", req, err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	if userInfo.UserId == "" {
		g.Log().Warning(ctx, "TFAInfo:", req, userInfo, err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	info, err := service.DB().FetchTfaInfo(ctx, userInfo.UserId)
	if err != nil || info == nil {
		g.Log().Warning(ctx, "TFAInfo:", req, userInfo, err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	///
	// rst := entity.Tfa{}
	// err = dao.Tfa.Ctx(ctx).Where(dao.Tfa.Columns().UserId, userInfo.UserId).Scan(&rst)
	// if err != nil {
	// 	g.Log().Error(ctx, "TFAinfo no info?:", userInfo, err, req)
	// 	return nil, err
	// }

	res = &v1.TFAInfoRes{
		Phone:       info.Phone,
		UpPhoneTime: info.PhoneUpdatedAt,
		Mail:        info.Mail,
		UpMailTime:  info.MailUpdatedAt,
	}
	return
}
