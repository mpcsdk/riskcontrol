package ethtx

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"riskcontral/internal/service"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	url := "https://base-mainnet.g.alchemy.com/v2/docs-demo"

	payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"eth_newFilter\",\"params\":[{\"address\":[\"0xb59f67a8bff5d8cd03f6ac17265c550ed8f33907\"],\"fromBlock\":\"0x429d3b\",\"toBlock\":\"latest\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x00000000000000000000000000b46c2526e227482e2ebb8f4c69e4674d262e75\",\"0x00000000000000000000000054a2d42a40f51259dedd1978f6c118a0f0eff078\"]}]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}

type sEthTx struct {
	sctx context.Context
}

func (s *sEthTx) getFilterLogs(ctx context.Context) error {
	for {
		select {
		case <-s.sctx.Done():
		default:
		}
		/////

	}
	return nil
}

// /
func (s *sEthTx) Data2Args(ctx context.Context, target string, data string) (map[string]interface{}, error) {
	//
	// rawBytes, err := hex.DecodeString(req.Data)
	// rtx := &types.Transaction{}
	// rlp.Decode(r, val)
	// d := rtx.Hash()
	// fmt.Println(common.Bytes2Hex(d))
	///
	data = strings.TrimPrefix(data, "0x")
	contractABI, err := service.RulesDb().GetAbi(ctx, target)
	if err != nil {
		return nil, err
	}
	c, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, err
	}
	bdata := common.Hex2Bytes(data)
	method, err := c.MethodById(bdata[:4])
	if err != nil {
		return nil, errors.New("Risk check failed")
	}
	args := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(args, bdata[4:])
	if err != nil {
		return nil, err
	}
	//todo: tx argx
	fmt.Println(args)
	// res, err := s.client.PerformRisk(s.ctx, &v1.RiskReq{
	// 	To:   tx.To,
	// 	From: tx.From,
	// 	//todo: data,
	// 	Data: tx.Data,
	// })
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return args, nil
}

func new() *sEthTx {
	return &sEthTx{}
}
func init() {
	service.RegisterEthTx(new())
}
