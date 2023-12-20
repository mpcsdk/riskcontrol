package nats

import (
	"context"
	"riskcontral/api/risk/nrpc"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (*NrpcServer) RpcContractAbiBriefs(ctx context.Context, req *nrpc.ContractAbiBriefsReq) (res *nrpc.ContractAbiBriefsRes, err error) {
	g.Log().Notice(ctx, "RpcContractAbiBriefs:", "req:", req)
	briefs, err := service.DB().GetContractAbiBriefs(ctx, "seceneNo", "address")
	if err != nil {
		return nil, mpccode.CodeInternalError()
	}
	res = &nrpc.ContractAbiBriefsRes{
		Briefs: map[string]*nrpc.ContractAbiBrief{},
	}
	for _, b := range briefs {
		res.Briefs[b.ContractAddress] = &nrpc.ContractAbiBrief{
			SceneNo: b.SceneNo,
			Address: b.ContractAddress,
			Name:    b.ContractName,
			Kind:    b.ContractKind,
		}
	}

	return res, nil
}

func (*NrpcServer) RpcContractAbi(ctx context.Context, req *nrpc.ContractAbiReq) (res *nrpc.ContractAbiRes, err error) {
	g.Log().Notice(ctx, "RpcContractAbi:", "req:", req)
	abi, err := service.DB().GetContractAbi(ctx, req.SceneNo, req.Address)
	if err != nil {
		return nil, mpccode.CodeInternalError()
	}
	res = &nrpc.ContractAbiRes{
		Id:              int64(abi.Id),
		ContractName:    abi.ContractName,
		ContractAddress: abi.ContractAddress,
		SceneNo:         abi.SceneNo,
		AbiContent:      abi.AbiContent,
		ContractKind:    abi.ContractKind,
	}
	return res, nil
}
