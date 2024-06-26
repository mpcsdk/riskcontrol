package riskserver

// func (s *sRiskCtrl) NotityContractAbi(ctx context.Context, notice *mq.ContractAbiMsg) error {
// 	abi, err := s.riskCtrlRule.GetContractAbi(ctx, notice.Data.ChainId, notice.Data.ContractAddress, true)
// 	if err != nil {
// 		g.Log().Error(ctx, notice, err)
// 		return err
// 	}
// 	///
// 	analzer := s.chainAnalzer[notice.Data.ChainId]
// 	if analzer == nil {
// 		analzer = &chainAnalzer{
// 			FtruleMap:  map[string]*entity.Contractrule{},
// 			NftruleMap: map[string]*entity.Contractrule{},
// 			Analzer:    analzyer.NewAnalzer(),
// 		}
// 		s.chainAnalzer[notice.Data.ChainId] = analzer
// 	}
// 	err = analzer.Analzer.AddAbi(abi.ContractAddress, abi.AbiContent)
// 	if err != nil {
// 		g.Log().Error(ctx, notice, err)
// 		return err
// 	}
// 	return nil
// }

// // /
// func (s *sRiskCtrl) NotityContractRule(ctx context.Context, notice *mq.ContractRuleMsg) error {
// 	///
// 	s.riskCtrlRule.ClearContractRuleCache(ctx, notice.Data.ChainId, "", notice.Data.ContractAddress)
// 	rule, err := s.riskCtrlRule.GetContractRule(ctx, notice.Data.ChainId, notice.Data.ContractAddress, true)
// 	if err != nil {
// 		g.Log().Error(ctx, notice, err)
// 		return err
// 	}
// 	/////
// 	analzer := s.chainAnalzer[rule.ChainId]
// 	if analzer == nil {
// 		analzer = &chainAnalzer{
// 			FtruleMap:  map[string]*entity.Contractrule{},
// 			NftruleMap: map[string]*entity.Contractrule{},
// 			Analzer:    analzyer.NewAnalzer(),
// 		}
// 		s.chainAnalzer[notice.Data.ChainId] = analzer
// 	}
// 	if analzer != nil {
// 		switch rule.ContractKind {
// 		case "erc20":
// 			analzer.FtruleMap[rule.ContractAddress] = rule
// 		case "erc1155", "erc721":
// 			analzer.NftruleMap[rule.ContractAddress] = rule
// 		default:
// 			g.Log().Error(ctx, notice, err)

// 		}
// 	}
// 	return nil
// }
