package ethtx

import (
	"riskcontral/common/chttp"
	"riskcontral/common/ethtx"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func Test_Logs(t *testing.T) {
	///
	query := &ethereum.FilterQuery{
		Addresses: []common.Address{
			common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
		},
	}
	call := ethtx.JsonrpcCall("eth_newFilter", []interface{}{query})

	newfilerresult := ""
	res := &ethtx.JsonrpcResp{
		Result: &newfilerresult,
	}

	c := chttp.NewClient("https://mainnet-rpc.dexon.org")

	_, err := c.Post(call, res)
	if err != nil {
		t.Error(err)
	}
	////

	for {
		call = ethtx.JsonrpcCall("eth_getFilterChanges", []interface{}{newfilerresult})

		logs := []*types.Log{}
		res = &ethtx.JsonrpcResp{
			Result: logs,
		}

		_, err = c.Post(call, res)
		if err != nil {
			t.Error(err)
		}
		if res.Error != nil {
			t.Error(res.Error)
		}
		if len(logs) > 0 {
			break
		}
		time.Sleep(15 * time.Second)
	}

	// req.Header.Add("accept", "application/json")
	// req.Header.Add("content-type", "application/json")
	// fmt.Println(logs)
}
