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
