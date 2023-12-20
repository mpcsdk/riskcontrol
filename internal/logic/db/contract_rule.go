package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// var NftRules = map[string]*mpcmodel.NftRule{}

// var FtRules = map[string]*mpcmodel.FtRule{}

func (s *sDB) GetContractRuleBriefs(ctx context.Context, SceneNo string, kind string) ([]*entity.Contractrule, error) {
	model := dao.Contractrule.Ctx(ctx).Fields(
		dao.Contractrule.Columns().SceneNo,
		dao.Contractrule.Columns().ContractAddress,
		dao.Contractrule.Columns().ContractName,
		dao.Contractrule.Columns().ContractKind,
	)
	if SceneNo != "" {
		model = model.Where(dao.Contractrule.Columns().SceneNo, SceneNo)
	}
	if kind != "" {
		model = model.Where(dao.Contractrule.Columns().ContractKind, kind)
	}
	rst, err := model.All()
	if err != nil {
		g.Log().Error(ctx, "GetContractRuleBriefs:", "sceneNo", SceneNo, "kind", kind, "err", err)
		return nil, mpccode.CodeInternalError()
	}
	///
	rule := []*entity.Contractrule{}
	rst.Structs(&rule)
	return rule, nil
}

// /
func (s *sDB) GetContractRule(ctx context.Context, SceneNo string, address string) (*entity.Contractrule, error) {
	rst, err := dao.Contractrule.Ctx(ctx).
		Where(dao.Contractrule.Columns().SceneNo, SceneNo).
		Where(dao.Contractrule.Columns().ContractAddress, address).One()
	if err != nil {
		g.Log().Error(ctx, "GetContractRule:", "sceneNo", SceneNo, "address", address, "err", err)
		return nil, mpccode.CodeInternalError()
	}
	// /
	rule := &entity.Contractrule{}
	rst.Struct(&rule)
	return rule, nil
}
