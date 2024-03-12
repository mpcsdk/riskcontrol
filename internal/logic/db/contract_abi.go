package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
)

func (s *sDB) GetContractAbiBriefs(ctx context.Context, chainId string, kind string) ([]*entity.Contractabi, error) {
	model := dao.Contractabi.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.Contractabi.Table() + "GetContractAbiBriefs" + chainId + kind,
		Force:    true,
	}).Fields(
		dao.Contractabi.Columns().ChainId,
		dao.Contractabi.Columns().ContractAddress,
		dao.Contractabi.Columns().ContractName,
		dao.Contractabi.Columns().ContractKind,
	)
	if chainId != "" {
		model = model.Where(dao.Contractabi.Columns().ChainId, chainId)
	}
	if kind != "" {
		model = model.Where(dao.Contractabi.Columns().ContractKind, kind)
	}
	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	///
	rule := []*entity.Contractabi{}
	rst.Structs(&rule)
	return rule, nil
}

// /
func (s *sDB) GetContractAbi(ctx context.Context, chainId string, address string) (*entity.Contractabi, error) {
	rst, err := dao.Contractabi.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.Contractabi.Table() + "GetContractAbi" + chainId + address,
		Force:    true,
	}).
		Where(dao.Contractabi.Columns().ChainId, chainId).
		Where(dao.Contractabi.Columns().ContractAddress, address).One()
	if err != nil {
		return nil, err
	}
	// /
	rule := &entity.Contractabi{}
	rst.Struct(&rule)
	return rule, nil
}
