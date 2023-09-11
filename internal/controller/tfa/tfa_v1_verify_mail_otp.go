package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/service"
)

func (c *ControllerV1) VerifyMailOTP(ctx context.Context, req *v1.VerifyMailOTPReq) (res *v1.VerifyMailOTPRes, err error) {
	return nil, service.TFA().VerifyCode(ctx, req.Token, "upMail", req.Code)
}
