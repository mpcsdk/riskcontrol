package ethtx

import (
	"encoding/json"
	"riskcontral/common/http"

	"github.com/ethereum/go-ethereum/ethclient"
)

type ChainNode struct {
	url       string
	chainId   string
	clinet    *http.Client
	ethclient *ethclient.Client
}

///

func NewChainNode(url string) *ChainNode {
	s := &ChainNode{url: url}
	c := http.NewClient(url)
	s.clinet = c

	ethclient, err := ethclient.Dial(url)
	if err != nil {
		return nil
	}
	s.ethclient = ethclient

	return s
}
func (s *ChainNode) EthClient() *ethclient.Client {
	return s.ethclient
}

func (s *ChainNode) ChainNodeCall(method string, args ...interface{}) (json.RawMessage, error) {
	msg, _ := NewMessage(method, args...)
	// data, _ := json.Marshal(msg)
	rst, err := s.clinet.Post(msg)
	if err != nil {
		return nil, err
	}
	///
	msg, err = ParseMessage(rst)
	if err != nil {
		return nil, err
	}
	if msg.Error != nil {
		return nil, msg.Error
	}

	return msg.Result, nil

}
