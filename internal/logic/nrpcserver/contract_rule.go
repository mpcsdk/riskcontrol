package nats

import (
	"context"
	"encoding/json"
	v1 "riskcontral/api/risk/nrpc/v1"
	"riskcontral/internal/model"
	"riskcontral/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mq"
	"github.com/nats-io/nats.go"
)

func (s *sNrpcServer) NatsPub() {

	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanSubscribe(mq.RiskCtrlSubsject, ch)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			case msg := <-ch:
				data := &mq.RiskCtrlMQ{}
				err := json.Unmarshal(msg.Data, data)
				rule := &mq.ContractNotice{}
				err = gconv.Struct(data.Data, &rule)
				if err != nil {
					g.Log().Error(s.ctx, err)
				}
				///
				service.Risk().Notify(s.ctx, data.GetKind(), rule)
			case <-s.ctx.Done():
				sub.Unsubscribe()
				close(ch)
				sub.Drain()
			}
		}
	}()
}

// /
func (*sNrpcServer) RpcContractRuleBriefs(ctx context.Context, req *v1.ContractRuleBriefsReq) (res *v1.ContractRuleBriefsRes, err error) {
	briefs, err := service.DB().GetContractRuleBriefs(ctx, req.SceneNo, req.Address)
	if err != nil {
		return nil, gerror.NewCode(mpccode.CodeInternalError)
	}
	res = &v1.ContractRuleBriefsRes{
		Briefs: map[string]*v1.ContractRuleBriefs{},
	}
	for _, b := range briefs {
		res.Briefs[b.ContractAddress] = &v1.ContractRuleBriefs{
			SceneNo: b.SceneNo,
			Address: b.ContractAddress,
			Name:    b.ContractName,
			Kind:    b.ContractKind,
		}
	}
	return res, nil
}

func (*sNrpcServer) RpcContractRule(ctx context.Context, req *v1.ContractRuleReq) (res *v1.ContractRuleRes, err error) {
	rule, err := service.DB().GetContractRule(ctx, req.SceneNo, req.Address)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeInternalError)
	}
	res = model.ContractRuleEntity2Rpc(rule)
	return res, nil

}
