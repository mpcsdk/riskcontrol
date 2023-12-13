package tfa

import (
	"context"

	"riskcontral/api/risk/nrpc"
	v1 "riskcontral/api/tfa/v1"
)

func (c *ControllerV1) TfaRequest(ctx context.Context, req *v1.TfaRequestReq) (res *v1.TfaRequestRes, err error) {
	//
	///
	tres, err := c.nrpc.RpcTfaRequest(ctx, &nrpc.TfaRequestReq{
		Token:    req.Token,
		CodeType: req.CodeType,
	})
	if err != nil {
		return nil, err
	}
	///
	res = &v1.TfaRequestRes{
		RiskSerial: tres.RiskSerial,
		VList:      tres.VList,
	}
	return res, nil
	///
}
