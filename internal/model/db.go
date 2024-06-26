package model

import (
	"riskcontrol/api/riskctrl"
	"strings"

	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

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

//			WhiteAddrList: strings.Split(e.WhiteAddrList, ","),
//		}
//	}
func ContractRuleEntity2Rpc(e *entity.Contractrule) *riskctrl.ContractRuleRes {
	return &riskctrl.ContractRuleRes{
		Contract:         e.ContractAddress,
		Name:             e.ContractName,
		Kind:             e.ContractKind,
		MethodName:       e.MethodName,
		MethodSig:        e.MethodSignature,
		MethodFromField:  e.MethodFromField,
		MethodToField:    e.MethodToField,
		MethodValueField: e.MethodValueField,

		WhiteAddrList: func() []string {
			trimStr := strings.TrimSpace(e.WhiteAddrList)
			if len(trimStr) > 0 {
				return strings.Split(trimStr, ",")
			}
			return nil
		}(),
	}
}
