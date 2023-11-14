package nats

import (
	"context"
	v1 "riskcontral/api/risk/nrpc/v1"
	"riskcontral/internal/model"
	"riskcontral/internal/service"
)

func (s *sNrpcServer) RpcRiskTFA(ctx context.Context, req *v1.TFARiskReq) (res *v1.TFARiskRes, err error) {
	tfaInfo, err := service.NrpcClient().RpcTfaInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	riskData := &model.RiskTfa{
		UserId: req.UserId,
		Type:   req.Type,
		Mail:   req.Mail,
		Phone:  req.Phone,
	}
	riskSerial, code := service.Risk().RiskTFA(ctx, tfaInfo, riskData)
	return &v1.TFARiskRes{
		Ok:         code,
		RiskSerial: riskSerial,
	}, nil
}
