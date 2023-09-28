package analyzsigndata

import (
	"encoding/hex"
	"encoding/json"
	"errors"
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
type Analzyer struct {
	abis map[string]string
}

func NewAnalzer() *Analzyer {
	return &Analzyer{
		abis: map[string]string{},
	}
}
func (s *Analzyer) AnalzySignTxData(signData string) (*AnalzyedTx, error) {
	signtx := &SignTx{}
	err := json.Unmarshal([]byte(signData), signtx)
	if err != nil {
		return nil, err
	}
	///
	atx := &AnalzyedTx{}
	atx.Address = signtx.Address
	///
	for _, tx := range signtx.Txs {
		adata, err := s.analzyTx(tx)
		if err != nil {
			return nil, err
		}
		atx.Txs = append(atx.Txs, adata)
	}
	return atx, nil
}
func (s *Analzyer) AddAbi(addr string, abi string) {
	s.abis[addr] = abi
}

// //
func (s *Analzyer) analzyTx(tx *SignTxData) (*AnalzyedTxData, error) {

	if abistr, ok := s.abis[tx.Target]; !ok {
		return nil, errors.New("abi not found:"+tx.Target)
	} else {

		///
		contract, err := abi.JSON(strings.NewReader(abistr))
		if err != nil {
			return nil, err
		}
		//data
		dataByte, err := hex.DecodeString(strings.TrimPrefix(tx.Data, "0x"))
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
