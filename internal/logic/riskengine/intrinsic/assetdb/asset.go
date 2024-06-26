package assetdb

import (
	"context"
	"riskcontrol/internal/conf"
	"riskcontrol/internal/logic/riskengine/intrinsic/slice"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/ethtx/ethutil"
	"github.com/mpcsdk/mpcCommon/mpcdao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type AssetAttr struct {
	ChainId   int64
	Symbol    string
	Contract  string
	AssetKind string
	Decimal   int
	TokenId   string
	////
	Value float64
}

func (s *IntrinsicAssetDB) AssetAttr(chainId int64, tx *analzyer.AnalzyedSignTx) *AssetAttr {
	g.Log().Debug(s.ctx, "AssetAttr:", "chainId:", chainId, "tx:", tx)
	if tx == nil {
		return nil
	}
	////
	target := tx.Target
	if target == "" {
		///is native
		chain := s.allchaincfg[chainId]
		// chain, err := s.chaincfg.GetCfg(s.ctx, chainId)
		if chain == nil {
			g.Log().Warning(s.ctx, "Asset not exists:", chainId, tx.Target)
			return nil
		}
		////
		f := ethutil.BigDecimal2Float64(tx.Value.Int(), 18)
		rst := &AssetAttr{
			ChainId:   chainId,
			Symbol:    chain.Coin,
			Decimal:   18,
			AssetKind: "external",
			Value:     f,
		}
		g.Log().Debug(s.ctx, "AssetAttr native:", rst)
		return rst
	}
	/////contract
	abi, err := s.riskCtrlRule.GetContractAbi(s.ctx, chainId, target, false)
	if err != nil {
		g.Log().Warning(s.ctx, "Asset:", chainId, target, err)
		return nil
	}
	rst := &AssetAttr{
		ChainId:   chainId,
		Symbol:    abi.ContractName,
		Decimal:   abi.Decimal,
		Contract:  abi.ContractAddress,
		AssetKind: abi.ContractKind,
	}
	if abi.ContractKind == "erc20" {
		if tx.Value == nil || tx.Value.CmpInt(0) == 0 {
			rst.Value = 0
		} else {
			f := ethutil.BigDecimal2Float64(tx.Value.Int(), abi.Decimal)
			rst.Value = f
		}
	} else if abi.ContractKind == "erc721" {

	} else if abi.ContractKind == "erc1155" {
		if tx.Value == nil || tx.Value.CmpInt(0) == 0 {
			rst.Value = 0
		} else {
			rst.Value = float64(tx.Value.Int().Int64())
		}
		rst.TokenId = abi.TokenId
	} else {
		g.Log().Warning(s.ctx, "AssetAttr unknow kind:", tx)
	}

	g.Log().Debug(s.ctx, "AssetAttr:", rst)
	return rst
}

type IntrinsicAssetDB struct {
	ctx context.Context
	//
	enhanced_riskctrl *mpcdao.EnhancedRiskCtrl
	chaincfg          *mpcdao.ChainCfg
	riskCtrlRule      *mpcdao.RiskCtrlRule
	////
	chainsAsset map[int64]map[string]*AssetAttr
	allChainIds *slice.NumberList
	allchaincfg map[int64]*entity.Chaincfg
}

func NewIntrinsicAssetDB(ctx context.Context) *IntrinsicAssetDB {
	///
	r := g.Redis("aggRiskCtrl")
	_, err := r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	///
	s := &IntrinsicAssetDB{
		ctx:               ctx,
		enhanced_riskctrl: mpcdao.NewEnhancedRiskCtrl(r, conf.Config.Cache.Duration),
		chaincfg:          mpcdao.NewChainCfg(r, conf.Config.Cache.Duration),
		riskCtrlRule:      mpcdao.NewRiskCtrlRule(r, conf.Config.Cache.Duration),
		chainsAsset:       map[int64]map[string]*AssetAttr{},
		allchaincfg:       map[int64]*entity.Chaincfg{},
		allChainIds:       &slice.NumberList{},
	}
	////allchain
	chains, err := s.chaincfg.AllCfg(ctx)
	if err != nil {
		panic(err)
	}
	for _, chain := range chains {
		s.allchaincfg[chain.ChainId] = chain
		s.allChainIds.Add(chain.ChainId)
	}
	////contractabi
	abibriefs, err := s.riskCtrlRule.GetContractAbiBriefs(ctx, 0, "")
	if err != nil {
		panic(err)
	}
	for _, abi := range abibriefs {
		if _, ok := s.chainsAsset[abi.ChainId]; ok {
			s.chainsAsset[abi.ChainId][abi.ContractName] = &AssetAttr{
				ChainId:   abi.ChainId,
				Symbol:    abi.ContractName,
				Decimal:   abi.Decimal,
				AssetKind: abi.ContractKind,
				Contract:  abi.ContractAddress,
			}
		} else {
			m := map[string]*AssetAttr{
				abi.ContractName: &AssetAttr{
					ChainId:   abi.ChainId,
					Symbol:    abi.ContractName,
					Decimal:   abi.Decimal,
					AssetKind: abi.ContractKind,
					Contract:  abi.ContractAddress,
				},
			}
			s.chainsAsset[abi.ChainId] = m
		}
	}
	return s
}

//
