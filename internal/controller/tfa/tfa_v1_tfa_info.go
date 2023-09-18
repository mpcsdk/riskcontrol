package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) TFAInfo(ctx context.Context, req *v1.TFAInfoReq) (res *v1.TFAInfoRes, err error) {
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		g.Log().Warning(ctx, "TFAInfo:", req, err)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	///
	rst := entity.Tfa{}
	err = dao.Tfa.Ctx(ctx).Where(dao.Tfa.Columns().UserId, userInfo.UserId).Scan(&rst)
	if err != nil {
		g.Log().Error(ctx, "TFAinfo no info?:", err, req)
		return nil, gerror.NewCode(consts.CodeTFANotExist)
	}
	res = &v1.TFAInfoRes{
		Phone:       rst.Phone,
		UpPhoneTime: rst.PhoneUpdatedAt,
		Mail:        rst.Mail,
		UpMailTime:  rst.MailUpdatedAt,
	}
	return
}
