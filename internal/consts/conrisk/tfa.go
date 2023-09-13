package conrisk

type RiskTfa struct {
	UserId    string `json:"userId"`
	UserToken string `json:"token"`
	Kind      string `json:"kind"`
	///
	Mail  string `json:"mail"`
	Phone string `json:"phone"`
}
