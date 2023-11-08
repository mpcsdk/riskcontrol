package model

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
