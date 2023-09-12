package conrisk

type TfaUpMail struct {
	UserId string `json:"token"`
	Mail   string `json:"mail"`
}
type TfaUpPhone struct {
	UserId string `json:"token"`
	Phone  string `json:"phone"`
}
