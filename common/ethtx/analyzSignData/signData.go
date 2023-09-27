package analyzsigndata

import (
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type SignTx struct {
	ChainId uint64        `json:"chainId,omitempty"`
	Address string        `json:"address,omitempty"`
	Number  uint64        `json:"number,omitempty"`
	Txs     []*SignTxData `json:"txs,omitempty"`
	TxHash  string        `json:"txHash,omitempty"`
}
type SignTxData struct {
	Target string `json:"target,omitempty"`
	Data   string `json:"data,omitempty"`
}
type analzyer struct {
	abis map[string]string
}

func NewAnalzer() *analzyer {
	return &analzyer{
		abis: map[string]string{},
	}
}
func (s *analzyer) AnalzySignTxData(signData string) error {
	signtx := &SignTx{}
	err := json.Unmarshal([]byte(signData), signtx)
	if err != nil {
		return err
	}
	///
	atx := &AnalzyedTx{}
	///
	for _, tx := range signtx.Txs {
		adata, err := s.analzyTx(tx)
		if err != nil {
			return err
		}
		atx.Txs = append(atx.Txs, adata)
	}
	return nil
}
func (s *analzyer) AddAbi(addr string, abi string) {
	s.abis[addr] = abi
}

// //
func (s *analzyer) analzyTx(tx *SignTxData) (*AnalzyedTxData, error) {

	if abistr, ok := s.abis[tx.Target]; !ok {
		return nil, nil
	} else {

		///
		contract, err := abi.JSON(strings.NewReader(abistr))
		if err != nil {
			return nil, err
		}
		//data
		dataByte, err := hex.DecodeString(tx.Data)
		if err != nil {
			return nil, err
		}
		////
		method, err := contract.MethodById(dataByte[:4])
		if err != nil {
			return nil, err
		}
		args := make(map[string]interface{})
		err = method.Inputs.UnpackIntoMap(args, dataByte[4:])
		if err != nil {
			return nil, err
		}
		///
		atx := &AnalzyedTxData{
			Target:     tx.Target,
			MethodId:   hex.EncodeToString(method.ID),
			MethodName: method.RawName,
			Sig:        method.Sig,
			Data:       tx.Data,
			Args:       args,
		}
		return atx, nil
	}
}
