package riskserver

// func ContractRuleEntity2Model(e *entity.Contractrule) *mpcmodel.ContractRule {
// 	return &mpcmodel.ContractRule{
// 		Contract:         e.ContractAddress,
// 		Name:             e.ContractName,
// 		Kind:             e.ContractKind,
// 		MethodName:       e.MethodName,
// 		MethodSig:        e.MethodSignature,
// 		MethodFromField:  e.MethodFromField,
// 		MethodToField:    e.MethodToField,
// 		MethodValueField: e.MethodValueField,

// 		WhiteAddrList: strings.Split(e.WhiteAddrList, ","),
// 	}
// }

// func ContractRuleEntity2Rpc(e *entity.Contractrule) *riskctrl.ContractRuleRes {
// 	return &riskctrl.ContractRuleRes{
// 		Contract:         e.ContractAddress,
// 		Name:             e.ContractName,
// 		Kind:             e.ContractKind,
// 		MethodName:       e.MethodName,
// 		MethodSig:        e.MethodSignature,
// 		MethodFromField:  e.MethodFromField,
// 		MethodToField:    e.MethodToField,
// 		MethodValueField: e.MethodValueField,

// 		EventName:       e.EventName,
// 		EventSig:        e.EventSignature,
// 		EventTopic:      e.EventTopic,
// 		EventFromField:  e.EventFromField,
// 		EventToField:    e.EventToField,
// 		EventValueField: e.EventValueField,
// 		WhiteAddrList: func() []string {
// 			trimStr := strings.TrimSpace(e.WhiteAddrList)
// 			if len(trimStr) > 0 {
// 				return strings.Split(trimStr, ",")
// 			}
// 			return nil
// 		}(),
// 		ThresholdBigintBytes: e.Threshold.BigInt().Bytes(),
// 	}
// }
