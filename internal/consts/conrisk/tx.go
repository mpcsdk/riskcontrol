package conrisk

type RiskTx struct {
	Token    string `json:"token"`
	Address  string `json:"address"`
	Contract string `json:"contract"`

	//////
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
	//风控阈值
	Threshold int `json:"threshold"`
}
