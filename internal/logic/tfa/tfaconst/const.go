package tfaconst

import (
	"context"
	"errors"
	"riskcontral/internal/model"
)

var ErrRiskNotDone error = errors.New("risk not done")

const (
	VerifierKind_Nil   = "nil"
	VerifierKind_Phone = "Phone"
	VerifierKind_Mail  = "Mail"
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
	CodeType_TxNeedVerify   string = "upTxNeedVerify"
	CodeType_TfaRisk        string = "tfaRisk"
	// CodeType_TxRisk         string = "txRisk"
)

func CodeType2RiskKind(codeType string) RISKKIND {
	switch codeType {
	case CodeType_TfaBindPhone:
		return RiskKind_BindPhone
	case CodeType_TfaBindMail:
		return RiskKind_BindMail
	case CodeType_TfaUpdatePhone:
		return RiskKind_UpPhone
	case CodeType_TfaUpdateMail:
		return RiskKind_UpMail
	case CodeType_TfaRisk:
		return RiskKind_TfaRisk
	case CodeType_TxNeedVerify:
		return RiskKind_TxNeedVerify
	default:
		return RiskKind_Nil
	}
}

type RISKKIND string
type VERIFYKIND string

const (
	RiskKind_Nil          = "RiskKind_Nil"
	RiskKind_Tx           = "RiskKind_Tx"
	RiskKind_BindPhone    = "RiskKind_BindPhone"
	RiskKind_UpPhone      = "RiskKind_UpPhone"
	RiskKind_BindMail     = "RiskKind_BindMail"
	RiskKind_UpMail       = "RiskKind_UpMail"
	RiskKind_TxNeedVerify = "RiskKind_TxNeedVerify"
	RiskKind_TfaRisk      = "RiskKind_TfaRisk"
)

type IVerifier interface {
	Verify(ctx context.Context, verifierCode *model.VerifyCode) (RISKKIND, error)
	SetCode(string)
	RiskKind() RISKKIND
	VerifyKind() VERIFYKIND
	IsDone() bool
	///
	SendVerificationCode() (string, error)
	SendCompletion() error
	//
	Destination() string
}
