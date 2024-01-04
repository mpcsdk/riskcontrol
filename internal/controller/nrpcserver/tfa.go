package nats

import (
	"context"
	"riskcontral/api/riskserver"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (*NrpcServer) RpcTfaInfo(ctx context.Context, req *riskserver.TfaInfoReq) (res *riskserver.TfaInfoRes, err error) {
	g.Log().Notice(ctx, "RpcTfaInfo:", "req:", req)

	//trace
	ctx, span := gtrace.NewSpan(ctx, "RpcSendSmsCode")
	defer span.End()

	if req.UserId == "" {
		g.Log().Warning(ctx, "TFAInfo no userId:", "req:", req, "userInfo:", req)
		return nil, mpccode.CodeTFANotExist()
	}
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
	if err != nil {
		return nil, mpccode.CodeTFANotExist()
	}
	if tfaInfo == nil {
		return nil, nil
	}
	res = &riskserver.TfaInfoRes{
		UserId: tfaInfo.UserId,
		Phone:  tfaInfo.Phone,
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
	g.Log().Info(ctx, "RpcTfaInfo:", res)
	return res, nil
}
