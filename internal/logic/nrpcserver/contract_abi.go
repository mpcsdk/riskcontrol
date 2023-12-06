package nats

import (
	"context"
	v1 "riskcontral/api/risk/nrpc/v1"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (*sNrpcServer) RpcContractAbiBriefs(ctx context.Context, req *v1.ContractAbiBriefsReq) (res *v1.ContractAbiBriefsRes, err error) {
	briefs, err := service.DB().GetContractAbiBriefs(ctx, "seceneNo", "address")
	if err != nil {
		return nil, gerror.NewCode(mpccode.CodeInternalError)
	}
	res = &v1.ContractAbiBriefsRes{
		Briefs: map[string]*v1.ContractAbiBrief{},
	}
	for _, b := range briefs {
		res.Briefs[b.ContractAddress] = &v1.ContractAbiBrief{
			SceneNo: b.SceneNo,
			Address: b.ContractAddress,
			Name:    b.ContractName,
			Kind:    b.ContractKind,
		}
	}

	return res, nil
}

func (*sNrpcServer) RpcContractAbi(ctx context.Context, req *v1.ContractAbiReq) (res *v1.ContractAbiRes, err error) {
	abi, err := service.DB().GetContractAbi(ctx, req.SceneNo, req.Address)
	if err != nil {
		return nil, gerror.NewCode(mpccode.CodeInternalError)
	}
	res = &v1.ContractAbiRes{
		Id:              int64(abi.Id),
		ContractName:    abi.ContractName,
		ContractAddress: abi.ContractAddress,
		SceneNo:         abi.SceneNo,
		AbiContent:      abi.AbiContent,
		ContractKind:    abi.ContractKind,
	}
	return res, nil
}
