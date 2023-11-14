package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sDB) GetAbi(ctx context.Context, addr string) (string, error) {
	var data *entity.ContractAbi
	if addr == "" {
		return "", errArg
	}
	aggdo := &do.ContractAbi{
		Addr: addr,
	}
	rst, err := g.Model(dao.ContractAbi.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dbDuration,
		Name:     dao.ContractAbi.Table() + addr,
		Force:    false,
	}).Where(aggdo).One()
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("aggdo", aggdo),
		))
		return "", err
	}
	if rst == nil {
		err = gerror.Wrap(errEmpty, mpccode.ErrDetails(
			mpccode.ErrDetail("aggdo", aggdo),
		))

		return "", err
	}
	err = rst.Struct(&data)
	return data.Abi, err
}

func (s *sDB) GetAbiAll(ctx context.Context) ([]*entity.ContractAbi, error) {
	var data []*entity.ContractAbi

	rst, err := g.Model(dao.ContractAbi.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: s.dbDuration,
		Name:     dao.ContractAbi.Table() + "all",
		Force:    false,
		// }).Where("user_id", 1).One()
	}).All()
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("getallabi", ""),
		))
		return nil, err
	}
	if rst == nil {
		err = gerror.Wrap(errEmpty, mpccode.ErrDetails(
			mpccode.ErrDetail("getallabi", ""),
		))
		return nil, err
	}
	err = rst.Structs(&data)
	return data, err
}
