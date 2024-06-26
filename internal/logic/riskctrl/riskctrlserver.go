package riskserver

import (
	"riskcontrol/internal/conf"
	"riskcontrol/internal/service"
	"slices"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type riskCtrlTfa struct {
	IsEnable    bool
	RiskEngines *ruleEngine
}
type riskCtrlTx struct {
	IsEnable    bool
	RiskEngines *ruleEngine
}
type ruleEngine struct {
	Id         int
	Salience   int
	RuleName   string
	ChainId    string
	engineName string
}

func (s *ruleEngine) EngineName() string {
	if s.engineName == "" {
		s.engineName = s.RuleName + "_" + s.ChainId
	}
	return s.engineName
}

type chainAnalzer struct {
	chainId    int64
	FtruleMap  map[string]*entity.Contractrule
	NftruleMap map[string]*entity.Contractrule
	Analzer    *analzyer.Analzyer
}

type RiskRuleId string
type sRiskCtrl struct {
	chainAnalzer map[int64]*chainAnalzer
	///chain cfg
	allChain     map[int64]*entity.Chaincfg
	allChainIds  []int64
	swapMap      map[int64]string
	nftMarketMap map[int64]string
}

// //

func New() *sRiskCtrl {
	///
	s := &sRiskCtrl{
		chainAnalzer: map[int64]*chainAnalzer{},
		///
		allChain:    map[int64]*entity.Chaincfg{},
		allChainIds: []int64{},
	}
	ctx := gctx.GetInitCtx()
	////
	chains, err := service.DB().ChainCfg().AllCfg(ctx)
	if err != nil {
		panic(err)
	}
	for _, chain := range chains {
		s.allChain[chain.ChainId] = chain
		s.allChainIds = append(s.allChainIds, chain.ChainId)
	}
	slices.Sort(s.allChainIds)
	//// abi
	abibriefs, err := service.DB().RiskCtrl().GetContractAbiBriefs(ctx, 0, "")
	if err != nil {
		panic(err)
	}
	for _, abis := range abibriefs {
		abi, err := service.DB().RiskCtrl().GetContractAbi(ctx, abis.ChainId, abis.ContractAddress, true)
		if err != nil {
			panic(err)
		}
		///
		analzer := s.chainAnalzer[abi.ChainId]
		if analzer == nil {
			analzer = &chainAnalzer{
				FtruleMap:  map[string]*entity.Contractrule{},
				NftruleMap: map[string]*entity.Contractrule{},
				Analzer:    analzyer.NewAnalzer(),
			}
			s.chainAnalzer[abi.ChainId] = analzer
		}
		err = analzer.Analzer.AddAbi(abi.ContractAddress, abi.AbiContent)
		if err != nil {
			panic(err)
		}
	}
	/// cnotract rule
	rulebriefs, err := service.DB().RiskCtrl().GetContractRuleBriefs(ctx, 0, "")
	if err != nil {
		panic(err)
	}
	for _, brief := range rulebriefs {
		rule, err := service.DB().RiskCtrl().GetContractRule(ctx, brief.ChainId, brief.ContractAddress, true)
		if err != nil {
			panic(err)
		}
		// ///
		analzer := s.chainAnalzer[rule.ChainId]
		if analzer == nil {
			analzer = &chainAnalzer{
				FtruleMap:  map[string]*entity.Contractrule{},
				NftruleMap: map[string]*entity.Contractrule{},
				Analzer:    analzyer.NewAnalzer(),
			}
			s.chainAnalzer[rule.ChainId] = analzer
		}
		if analzer != nil {
			switch rule.ContractKind {
			case "erc20":
				analzer.FtruleMap[rule.ContractAddress] = rule
			case "erc1155", "erc721":
				analzer.NftruleMap[rule.ContractAddress] = rule
			default:
				panic("unknown contract kind")
			}
		}
	}
	////
	s.swapMap = make(map[int64]string)
	for _, v := range conf.Config.Swap {
		s.swapMap[v.ChainId] = v.Contract
	}
	s.nftMarketMap = make(map[int64]string)
	for _, v := range conf.Config.NftMarket {
		s.nftMarketMap[v.ChainId] = v.Contract
	}
	////

	return s
}
