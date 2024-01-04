package nats

import (
	"context"
	"riskcontral/api/riskserver"
)

func (*NrpcServer) RpcContractAbiBriefs(ctx context.Context, req *riskserver.ContractAbiBriefsReq) (res *riskserver.ContractAbiBriefsRes, err error) {
	return nil, nil
	// g.Log().Notice(ctx, "RpcContractAbiBriefs:", "req:", req)
	// briefs, err := service.DB().GetContractAbiBriefs(ctx, "seceneNo", "address")
	// if err != nil {
	// 	return nil, mpccode.CodeInternalError()
	// }
	// res = &riskserver.ContractAbiBriefsRes{
	// 	Briefs: map[string]*riskserver.ContractAbiBrief{},
	// }
	// for _, b := range briefs {
	// 	res.Briefs[b.ContractAddress] = &riskserver.ContractAbiBrief{
	// 		SceneNo: b.SceneNo,
	// 		Address: b.ContractAddress,
	// 		Name:    b.ContractName,
	// 		Kind:    b.ContractKind,
	// 	}
	// }

	// return res, nil
}

func (*NrpcServer) RpcContractAbi(ctx context.Context, req *riskserver.ContractAbiReq) (res *riskserver.ContractAbiRes, err error) {
	// g.Log().Notice(ctx, "RpcContractAbi:", "req:", req)
	// abi, err := service.DB().GetContractAbi(ctx, req.SceneNo, req.Address)
	// if err != nil {
	// 	return nil, mpccode.CodeInternalError()
	// }
	// res = &riskserver.ContractAbiRes{
	// 	Id:              int64(abi.Id),
	// 	ContractName:    abi.ContractName,
	// 	ContractAddress: abi.ContractAddress,
	// 	SceneNo:         abi.SceneNo,
	// 	AbiContent:      abi.AbiContent,
	// 	ContractKind:    abi.ContractKind,
	// }
	// return res, nil
	return nil, nil
}
