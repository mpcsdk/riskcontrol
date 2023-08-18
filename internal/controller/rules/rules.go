package rules

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "riskcontral/api/rules/v1"
	"riskcontral/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) PerformRisk(ctx context.Context, req *v1.RiskReq) (res *v1.RiskRes, err error) {
	param := map[string]interface{}{}
	err = json.Unmarshal([]byte(req.Data), &param)
	if err != nil {
		return nil, err
	}
	// for k, v := range req.Param {
	// 	fmt.Println(reflect.TypeOf(v))
	// 	fmt.Println(reflect.ValueOf(v).Type())
	// 	n := v.(json.Number)
	// 	param[k], _ = n.Int64()
	// }
	rst, err := service.LEngine().Exec("test", param)
	fmt.Println(rst)
	res = &v1.RiskRes{
		Result: rst,
	}
	return res, err
}
