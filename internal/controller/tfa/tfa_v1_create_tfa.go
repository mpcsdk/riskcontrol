package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/consts"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// @Summary 验证token，注册用户tfa
func (c *ControllerV1) CreateTFA(ctx context.Context, req *v1.CreateTFAReq) (res *v1.CreateTFARes, err error) {
	///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeAuthFailed)
	}
	///
	_, err = dao.Tfa.Ctx(ctx).Insert(entity.Tfa{
		UserId:         userInfo.UserId,
		CreatedAt:      gtime.Now(),
		Mail:           req.Mail,
		Phone:          req.Phone,
		PhoneUpdatedAt: gtime.Now(),
		MailUpdatedAt:  gtime.Now(),
	})
	return nil, err
}
