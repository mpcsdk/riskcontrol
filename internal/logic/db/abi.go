package db

import (
	"context"
	"riskcontral/internal/consts"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sDB) GetAbi(ctx context.Context, addr string) (string, error) {
	var data *entity.ContractAbi
	if addr == "" {
		return "", nil
	}
	rst, err := g.Model(dao.ContractAbi.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     dao.ContractAbi.Table() + addr,
		Force:    false,
		// }).Where("user_id", 1).One()
	}).Where(do.ContractAbi{
		Addr: addr,
	}).One()
	if err != nil {
		g.Log().Warning(ctx, "GetAbi:", addr, err)
		return "", gerror.NewCode(consts.CodeInternalError)
	}
	if rst == nil {
		g.Log().Warning(ctx, "GetAbi not exist:", addr)
		return "", gerror.NewCode(consts.CodeInternalError)
	}
	err = rst.Struct(&data)
	return data.Abi, err
}
