package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
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
		g.Log().Error(ctx, "GetContractAbiBriefs:", "sceneNo", SceneNo, "kind", kind, "err", err)
		return nil, mpccode.CodeInternalError()
	}
	///
	rule := []*entity.Contractabi{}
	err = rst.Structs(&rule)
	if err != nil {
		g.Log().Error(ctx, "GetContractAbiBriefs:", "sceneNo", SceneNo, "kind", kind, "err", err)
		return nil, mpccode.CodeInternalError()
	}
	return rule, nil
}

// /
func (s *sDB) GetContractAbi(ctx context.Context, SceneNo string, address string) (*entity.Contractabi, error) {
	rst, err := dao.Contractabi.Ctx(ctx).
		Where(dao.Contractabi.Columns().SceneNo, SceneNo).
		Where(dao.Contractabi.Columns().ContractAddress, address).One()
	if err != nil {
		g.Log().Error(ctx, "GetContractAbi:", "sceneNo", SceneNo, "address", address, "err", err)
		return nil, mpccode.CodeInternalError()
	}
	// /
	rule := &entity.Contractabi{}
	err = rst.Struct(&rule)
	if err != nil {
		g.Log().Error(ctx, "GetContractAbi:", "sceneNo", SceneNo, "address", address, "err", err)
		return nil, mpccode.CodeInternalError()
	}
	return rule, nil
}
