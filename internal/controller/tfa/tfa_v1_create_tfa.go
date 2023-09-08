package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

func (c *ControllerV1) CreateTFA(ctx context.Context, req *v1.CreateTFAReq) (res *v1.CreateTFARes, err error) {
	//todo:
	_, err = dao.Tfa.Ctx(ctx).Insert(entity.Tfa{
		Token:          req.Token,
		CreatedAt:      gtime.Now(),
		Mail:           req.Mail,
		Phone:          req.Phone,
		PhoneUpdatedAt: gtime.Now(),
		MailUpdatedAt:  gtime.Now(),
	})
	return nil, err
}
