package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"
)

func (s *sDB) GetContractAbiBriefs(ctx context.Context, SceneNo string, kind string) ([]*entity.Contractabi, error) {
	model := dao.Contractabi.Ctx(ctx).Fields(
		dao.Contractabi.Columns().SceneNo,
		dao.Contractabi.Columns().ContractAddress,
		dao.Contractabi.Columns().ContractName,
		dao.Contractabi.Columns().ContractKind,
	)
	if SceneNo != "" {
		model = model.Where(dao.Contractabi.Columns().SceneNo, SceneNo)
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
func (s *sDB) GetContractAbi(ctx context.Context, SceneNo string, address string) (*entity.Contractabi, error) {
	rst, err := dao.Contractabi.Ctx(ctx).
		Where(dao.Contractabi.Columns().SceneNo, SceneNo).
		Where(dao.Contractabi.Columns().ContractAddress, address).One()
	if err != nil {
		return nil, err
	}
	// /
	rule := &entity.Contractabi{}
	rst.Struct(&rule)
	return rule, nil
}
