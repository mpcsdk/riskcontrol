package model

import (
	v1 "riskcontral/api/risk/nrpc/v1"
	"riskcontral/internal/model/entity"
	"strings"

	"github.com/mpcsdk/mpcCommon/mpcmodel"
)

const (
	Kind_RiskTx  string = "riskTx"
	Kind_RiskTfa string = "riskTfa"
)
const (
	///
	Type_TfaBindPhone   string = "bindPhone"
	Type_TfaBindMail    string = "bindMail"
	Type_TfaUpdatePhone string = "updatePhone"
	Type_TfaUpdateMail  string = "updateMail"
)

type RiskKind string

const (
	RiskKind_Nil       = "RiskKind_Nil"
	RiskKind_Tx        = "RiskKind_Tx"
	RiskKind_BindPhone = "RiskKind_BindPhone"
	RiskKind_UpPhone   = "RiskKind_UpPhone"
	RiskKind_BindMail  = "RiskKind_BindMail"
	RiskKind_UpMail    = "RiskKind_UpMail"
)

type RiskStat struct {
	Kind string
	Type string
}
type RiskTfa struct {
	UserId    string `json:"userId"`
	UserToken string `json:"token"`
	Type      string `json:"type"`
	///
	Mail  string `json:"mail"`
	Phone string `json:"phone"`
}

func ContractRuleEntity2Model(e *entity.Contractrule) *mpcmodel.ContractRule {
	return &mpcmodel.ContractRule{
		Contract:         e.ContractAddress,
		Name:             e.ContractName,
		Kind:             e.ContractKind,
		MethodName:       e.MethodName,
		MethodSig:        e.MethodSignature,
		MethodFromField:  e.MethodFromField,
		MethodToField:    e.MethodToField,
		MethodValueField: e.MethodValueField,

		EventName:       e.EventName,
		EventSig:        e.EventSignature,
		EventTopic:      e.EventTopic,
		EventFromField:  e.EventFromField,
		EventToField:    e.EventToField,
		EventValueField: e.EventValueField,
		WhiteAddrList:   strings.Split(e.WhiteAddrList, ","),
		Threshold:       e.Threshold.BigInt(),
		ThresholdNft: func() int64 {
			if e.ContractKind == "nft" {
				return e.Threshold.IntPart()
			}
			return 0
		}(),
	}
}
func ContractRuleEntity2Rpc(e *entity.Contractrule) *v1.ContractRuleRes {
	return &v1.ContractRuleRes{
		Contract:         e.ContractAddress,
		Name:             e.ContractName,
		Kind:             e.ContractKind,
		MethodName:       e.MethodName,
		MethodSig:        e.MethodSignature,
		MethodFromField:  e.MethodFromField,
		MethodToField:    e.MethodToField,
		MethodValueField: e.MethodValueField,

		EventName:       e.EventName,
		EventSig:        e.EventSignature,
		EventTopic:      e.EventTopic,
		EventFromField:  e.EventFromField,
		EventToField:    e.EventToField,
		EventValueField: e.EventValueField,
		WhiteAddrList: func() []string {
			trimStr := strings.TrimSpace(e.WhiteAddrList)
			if len(trimStr) > 0 {
				return strings.Split(trimStr, ",")
			}
			return nil
		}(),
		ThresholdBigintBytes: e.Threshold.BigInt().Bytes(),
	}
}
