package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) UpMail(ctx context.Context, req *v1.UpMailReq) (res *v1.UpMailRes, err error) {
	return nil, service.TFA().UpMail(ctx, req.Token, req.Mail)
}
