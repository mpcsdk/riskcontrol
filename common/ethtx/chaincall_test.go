package ethtx

import (
	"context"
	"encoding/json"
	"math/big"
	"riskcontral/common/ethtx/ethmodel"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

func Test_chaincall(t *testing.T) {
	node := NewChainNode("https://robin.rangersprotocol.com/pubhub/api/jsonrpc")
	////

	msg, err := node.ChainNodeCall("eth_blockNumber")
	if err != nil {
		t.Error(err)
	}
	///

	rst := rpc.BlockNumber(0)
	err = json.Unmarshal(msg, &rst)
	if err != nil {
		t.Error(err)
	}
	t.Log(rst)
}

func Test_chaincall_log(t *testing.T) {
	node := NewChainNode("https://robin.rangersprotocol.com/pubhub/api/jsonrpc")
	////

	from := rpc.BlockNumber(56228576)
	to := rpc.BlockNumber(56228576)
	query := ethmodel.EthFilterQuery{
		Address: []common.Address{
			common.HexToAddress("0x9e4ac58cfbdf5cfe0685ad034bb5c6e26363a72a"),
		},
		FromBlock: &from,
		ToBlock:   &to,
	}
	msg, err := node.ChainNodeCall("eth_getLogs", query)
	if err != nil {
		t.Error(err)
	}
	///

	rst := []*ethmodel.Log{}
	err = json.Unmarshal(msg, &rst)
	if err != nil {
		t.Error(err)
	}
	t.Log(rst)
}

func Test_ethcall(t *testing.T) {
	node := NewChainNode("https://robin.rangersprotocol.com/pubhub/api/jsonrpc")
	////
	from := big.NewInt(56228576)
	to := big.NewInt(56228576)
	logs, err := node.EthClient().FilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{
			common.HexToAddress("0x9e4ac58cfbdf5cfe0685ad034bb5c6e26363a72a"),
		},
		FromBlock: from,
		ToBlock:   to,
	})
	if err != nil {
		t.Error(err)
	}
	///
	t.Log(logs)
}
func Test_eth_number(t *testing.T) {
	node := NewChainNode("https://robin.rangersprotocol.com/pubhub/api/jsonrpc")
	////
	number, err := node.EthClient().BlockNumber(context.Background())
	if err != nil {
		t.Error(err)
	}
	///
	t.Log(number)
}
