package nats

import (
	"context"
	"riskcontral/api/riskserver"
)

// func (s *NrpcServer) NatsPub() {

// 	ch := make(chan *nats.Msg, 64)
// 	sub, err := s.nc.ChanSubscribe(mq.RiskCtrlSubsject, ch)
// 	if err != nil {
// 		panic(err)
// 	}
// 	go func() {
// 		for {
// 			select {
// 			case msg := <-ch:
// 				err := msg.Respond([]byte("reple"))
// 				fmt.Println(err)
// 				data := &mq.RiskCtrlMQ{}
// 				err = json.Unmarshal(msg.Data, data)
// 				rule := &mq.ContractNotice{}
// 				err = gconv.Struct(data.Data, &rule)
// 				if err != nil {
// 					g.Log().Error(s.ctx, err)
// 				}
// 				///
// 				service.Risk().Notify(s.ctx, data.GetKind(), rule)
// 			case <-s.ctx.Done():
// 				sub.Unsubscribe()
// 				close(ch)
// 				sub.Drain()
// 			}
// 		}
// 	}()
// }

// /
func (*NrpcServer) RpcContractRuleBriefs(ctx context.Context, req *riskserver.ContractRuleBriefsReq) (res *riskserver.ContractRuleBriefsRes, err error) {
	// g.Log().Notice(ctx, "RpcContractRuleBriefs:", "req:", req)

	// briefs, err := service.DB().GetContractRuleBriefs(ctx, req.SceneNo, req.Address)
	// if err != nil {
	// 	return nil, mpccode.CodeInternalError()
	// }
	// res = &riskserver.ContractRuleBriefsRes{
	// 	Briefs: map[string]*riskserver.ContractRuleBriefs{},
	// }
	// for _, b := range briefs {
	// 	res.Briefs[b.ContractAddress] = &riskserver.ContractRuleBriefs{
	// 		SceneNo: b.SceneNo,
	// 		Address: b.ContractAddress,
	// 		Name:    b.ContractName,
	// 		Kind:    b.ContractKind,
	// 	}
	// }
	// return res, nil
	return nil, nil
}

func (*NrpcServer) RpcContractRule(ctx context.Context, req *riskserver.ContractRuleReq) (res *riskserver.ContractRuleRes, err error) {

	// g.Log().Notice(ctx, "RpcContractRule:", "req:", req)
	// rule, err := service.DB().GetContractRule(ctx, req.SceneNo, req.Address)
	// if err != nil {
	// 	return nil, mpccode.CodeInternalError()
	// }
	// res = model.ContractRuleEntity2Rpc(rule)
	return res, nil

}
