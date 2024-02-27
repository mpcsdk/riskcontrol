package model

import (
	"context"
)

const (
	Kind_RiskTx  string = "riskTx"
	Kind_RiskTfa string = "riskTfa"
)
const (
	///
	CodeType_TfaBindPhone   string = "bindPhone"
	CodeType_TfaBindMail    string = "bindMail"
	CodeType_TfaUpdatePhone string = "updatePhone"
	CodeType_TfaUpdateMail  string = "updateMail"
)

func CodeType2RiskKind(codeType string) RiskKind {
	switch codeType {
	case CodeType_TfaBindPhone:
		return RiskKind_BindPhone
	case CodeType_TfaBindMail:
		return RiskKind_BindMail
	case CodeType_TfaUpdatePhone:
		return RiskKind_UpPhone
	case CodeType_TfaUpdateMail:
		return RiskKind_UpMail
	default:
		return RiskKind_Nil
	}
}

type RiskPenndingKey string

func RiskPenddingKey(userId, riskSerial string) RiskPenndingKey {

	return RiskPenndingKey("riskPendding:" + userId + ":" + riskSerial)
}

type RiskKind string
type VerifyKind string

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
	UserId    string   `json:"userId"`
	UserToken string   `json:"token"`
	RiskKind  RiskKind `json:"type"`
	///
	Mail  string `json:"mail"`
	Phone string `json:"phone"`
}
type IVerifier interface {
	Verify(ctx context.Context, verifierCode *VerifyCode) (RiskKind, error)
	SetCode(string)
	RiskKind() RiskKind
	VerifyKind() VerifyKind
	IsDone() bool
	///
	SendVerificationCode() (string, error)
	SendCompletion() error
	//
	Destination() string
}
