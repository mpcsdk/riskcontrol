package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) VerifySmsCode(ctx context.Context, req *v1.VerifySmsCodeReq) (res *v1.VerifySmsCodeRes, err error) {
	return nil, service.TFA().VerifyCode(ctx, req.Token, "upPhone", req.Code)
}
