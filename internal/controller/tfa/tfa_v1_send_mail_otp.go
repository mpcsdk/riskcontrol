package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) SendMailOTP(ctx context.Context, req *v1.SendMailOTPReq) (res *v1.SendMailOTPRes, err error) {
	// service.
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
