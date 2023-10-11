package analzyer

import (
	"math/big"
	"riskcontral/internal/model"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var signData string = `{
	"chainId": 9527,
	"address": "0x77990137A0032b8f31d4C3AE696f60d6AFa0ba99",
	"number": 22,
	"txs": [
		{
			"_isUnipassWalletTransaction": true,
			"callType": 0,
			"revertOnError": true,
			"gasLimit": {
				"type": "BigNumber",
				"hex": "0x00"
			},
			"target": "0x71d9cfd1b7adb1e8eb4c193ce6ffbe19b4aee0db",
			"value": {
				"type": "BigNumber",
				"hex": "0x00"
			},
			"data": "0xa9059cbb000000000000000000000000752ab37a4471bf059602863f6c8225816975730e000000000000000000000000000000000000000000000000016345785d8a0000"
		}
	],
	"txHash": "0xc664d6ea9ffb7fcd1028c6c484ba72c4354789f4d89258c34e14440017174536"
}`

var contractABI string = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"guy","type":"address"},{"name":"wad","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"src","type":"address"},{"name":"dst","type":"address"},{"name":"wad","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"wad","type":"uint256"}],"name":"withdraw","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"dst","type":"address"},{"name":"wad","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"deposit","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"anonymous":false,"inputs":[{"indexed":true,"name":"src","type":"address"},{"indexed":true,"name":"guy","type":"address"},{"indexed":false,"name":"wad","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"src","type":"address"},{"indexed":true,"name":"dst","type":"address"},{"indexed":false,"name":"wad","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"dst","type":"address"},{"indexed":false,"name":"wad","type":"uint256"}],"name":"Deposit","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"src","type":"address"},{"indexed":false,"name":"wad","type":"uint256"}],"name":"Withdrawal","type":"event"}]
`

// var data string = `0xa9059cbb000000000000000000000000aa5c1d42f766c98089a233ce1496bce18cfac5840000000000000000000000000000000000000000000000000000000000989680`
// var contractABI string = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"guy","type":"address"},{"name":"wad","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"src","type":"address"},{"name":"dst","type":"address"},{"name":"wad","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"wad","type":"uint256"}],"name":"withdraw","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"dst","type":"address"},{"name":"wad","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"deposit","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"payable":true,"stateMutability":"payable","type":"fallback"},{"anonymous":false,"inputs":[{"indexed":true,"name":"src","type":"address"},{"indexed":true,"name":"guy","type":"address"},{"indexed":false,"name":"wad","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"src","type":"address"},{"indexed":true,"name":"dst","type":"address"},{"indexed":false,"name":"wad","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"dst","type":"address"},{"indexed":false,"name":"wad","type":"uint256"}],"name":"Deposit","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"src","type":"address"},{"indexed":false,"name":"wad","type":"uint256"}],"name":"Withdrawal","type":"event"}]`
// var data string = `0xa9059cbb000000000000000000000000752ab37a4471bf059602863f6c8225816975730e000000000000000000000000000000000000000000000000016345785d8a0000`
var ftRules = map[string]*model.FtRule{
	"0x71d9cfd1b7adb1e8eb4c193ce6ffbe19b4aee0db": &model.FtRule{
		Contract: "0x71d9cfd1b7adb1e8eb4c193ce6ffbe19b4aee0db",
		Name:     "RPG",

		MethodName:       "transfer",
		MethodSig:        "transfer(address,uint256)",
		MethodFromField:  "",
		MethodToField:    "dst",
		MethodValueField: "wad",

		EventName:       "Transfer",
		EventSig:        "Transfer(address,address,uint256)",
		EventTopic:      hexutil.Encode(crypto.Keccak256([]byte("Transfer(address,address,uint256)"))),
		EventFromField:  "from",
		EventToField:    "to",
		EventValueField: "value",

		SkipToAddr: []string{"0x71d9cfd1b7adb1e8eb4c193ce6ffbe19b4aee0db"},
		Threshold:  big.NewInt(0).SetUint64(1000000000000000000),
	},
}

func Test_AnalzyTx(t *testing.T) {
	analzer := NewAnalzer()
	analzer.AddAbi("0x71d9cfd1b7adb1e8eb4c193ce6ffbe19b4aee0db", contractABI)
	signTx, err := analzer.SignTx(signData)
	if err != nil {
		t.Error(err)
	}

	ethtx, err := analzer.AnalzyTxDataFT(
		"0x71d9cfd1b7adb1e8eb4c193ce6ffbe19b4aee0db",
		signTx.Txs[0],
		ftRules["0x71d9cfd1b7adb1e8eb4c193ce6ffbe19b4aee0db"])
	if err != nil {
		t.Error(err)
	}
	if ethtx.From == "" {
		ethtx.From = signTx.Address
	}
	if ethtx.Value.String() != "100000000000000000" {
		t.Error(ethtx)
	}
}
