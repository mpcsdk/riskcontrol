package tfa

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"

	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) TfaInfo(ctx context.Context, req *v1.TfaInfoReq) (res *v1.TfaInfoRes, err error) {
	g.Log().Notice(ctx, "TfaInfo:", "req:", req)
	//
	//limit
	if err := c.limiter.ApiLimit(ctx, req.Token, "TfaInfo"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	//
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, mpccode.CodeTFANotExist()
	}
	if userInfo.UserId == "" {
		return nil, mpccode.CodeTFANotExist()
	}
	///
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, userInfo.UserId)
	if err != nil {
		return nil, mpccode.CodeTFANotExist()
	}
	if tfaInfo == nil {
		return nil, nil
	}

	///
	res = &v1.TfaInfoRes{
		Phone: tfaInfo.Phone,
		UpPhoneTime: func() string {
			if tfaInfo.PhoneUpdatedAt == nil {
				return ""
			}

			return tfaInfo.PhoneUpdatedAt.String()
		}(),
		Mail: tfaInfo.Mail,
		UpMailTime: func() string {
			if tfaInfo.MailUpdatedAt == nil {
				return ""
			}
			return tfaInfo.MailUpdatedAt.String()
		}(),
	}
	return
}
