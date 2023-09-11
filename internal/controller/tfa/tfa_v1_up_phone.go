package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) UpPhone(ctx context.Context, req *v1.UpPhoneReq) (res *v1.UpPhoneRes, err error) {
	//

	return nil, service.TFA().UpPhone(ctx, req.Token, req.Phone)
}
