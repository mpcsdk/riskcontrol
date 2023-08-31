package rules

import (
	"context"
	"fmt"
	v1 "riskcontral/api/rules/v1"
	"riskcontral/internal/service"
	"strings"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) PerformRisk(ctx context.Context, req *v1.RiskReq) (res *v1.RiskRes, err error) {
	// tolower
	req.To = strings.ToLower(req.To)
	req.From = strings.ToLower(req.From)
	req.Data = strings.ToLower(req.Data)
	//uppack inputs to param
	// param := map[string]interface{}{}
	param, err := service.EthTx().Data2Args(req.To, req.Data)
	if err != nil {
		return nil, err
	}
	// err = json.Unmarshal([]byte(req.Data), &param)
	// if err != nil {
	// 	return nil, err
	// }
	// for k, v := range req.Param {
	// 	fmt.Println(reflect.TypeOf(v))
	// 	fmt.Println(reflect.ValueOf(v).Type())
	// 	n := v.(json.Number)
	// 	param[k], _ = n.Int64()
	// }
	rst, err := service.LEngine().Exec(req.From, param)
	fmt.Println(rst)
	res = &v1.RiskRes{
		Result: rst,
	}
	return res, err
}
