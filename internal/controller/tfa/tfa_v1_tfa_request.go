package tfa

import (
	"context"
	v1 "riskcontral/api/tfa/v1"
	"riskcontral/internal/model"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) TfaRequest(ctx context.Context, req *v1.TfaRequestReq) (res *v1.TfaRequestRes, err error) {
	//
	g.Log().Notice(ctx, "TfaRequest:", "req:", req)
	///
	// limit
	if err := c.limiter.ApiLimit(ctx, req.Token, "TfaRequest"); err != nil {
		return nil, err
	}
	//trace
	ctx, span := gtrace.NewSpan(ctx, "TfaRequest")
	defer span.End()
	/////
	//
	info, _ := service.UserInfo().GetUserInfo(ctx, req.Token)
	if info == nil {
		return nil, mpccode.CodeTokenInvalid()
	}
	///
	riskKind := model.CodeType2RiskKind(req.CodeType)
	//
	return service.TFA().TfaRequest(ctx, info, riskKind)
	///
	///
}
