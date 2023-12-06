package db

import (
	"context"
	"riskcontral/internal/dao"
	"riskcontral/internal/model/entity"
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
		return nil, err
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
		return nil, err
	}
	// /
	rule := &entity.Contractrule{}
	rst.Struct(&rule)
	return rule, nil
}

// func (s *sDB) GetNftRules(ctx context.Context, SceneNo string, address string) (*entity.Contractrule, error) {
// 	rst, err := dao.Contractrule.Ctx(ctx).
// 		Where(dao.Contractrule.Columns().SceneNo, SceneNo).
// 		Where(dao.Contractrule.Columns().ContractAddress, address).One()
// 	if err != nil {
// 		return nil, err
// 	}
// 	///
// 	rule := &entity.Contractrule{}
// 	rst.Struct(&rule)
// 	return rule, nil
// }

//	func (s *sDB) GetFtRules(ctx context.Context, SceneNo string, address string) (*entity.Contractrule, error) {
//		rst, err := dao.Contractrule.Ctx(ctx).
//			Where(dao.Contractrule.Columns().SceneNo, SceneNo).
//			Where(dao.Contractrule.Columns().ContractAddress, address).One()
//		if err != nil {
//			return nil, err
//		}
//		///
//		rule := &entity.Contractrule{}
//		rst.Struct(&rule)
//		return rule, nil
//	}
// func init() {
// 	cfg := gcfg.Instance()
// 	v, err := cfg.Get(context.Background(), "txRisk.nft")
// 	if err != nil {
// 		panic(err)
// 	}
// 	nftrules := []*mpcmodel.NftRule{}
// 	err = v.Structs(&nftrules)
// 	if err != nil {
// 		panic(err)
// 	}
// 	///
// 	///
// 	v, err = cfg.Get(context.Background(), "txRisk.ft")
// 	if err != nil {
// 		panic(err)
// 	}
// 	ftrules := []*mpcmodel.FtRule{}
// 	err = v.Structs(&ftrules)
// 	if err != nil {
// 		panic(err)
// 	}
// 	///tidy
// 	///
// 	for _, v := range nftrules {
// 		v.EventTopic = hexutil.Encode(crypto.Keccak256([]byte(v.EventSig)))
// 	}
// 	for _, v := range ftrules {
// 		v.EventTopic = hexutil.Encode(crypto.Keccak256([]byte(v.EventSig)))
// 	}
// 	///
// 	for _, v := range nftrules {
// 		NftRules[v.Contract] = v
// 	}
// 	for _, v := range ftrules {
// 		FtRules[v.Contract] = v
// 	}
// 	///
// 	g.Log().Notice(context.Background(),
// 		"nftRules:", NftRules,
// 		"ftRules:", FtRules,
// 	)
// }
