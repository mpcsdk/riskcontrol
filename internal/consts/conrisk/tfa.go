package conrisk

type TfaUpMail struct {
	Token string `json:"token"`
	Mail  string `json:"mail"`
}
type TfaUpPhone struct {
	Token string `json:"token"`
	Phone string `json:"phone"`
}
