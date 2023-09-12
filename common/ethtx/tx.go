package ethtx

import (
	"encoding/hex"

	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type AnalzyTxDataResp struct {
	MethodId   string
	MethodName string
	Sig        string
	Data       string
	Args       map[string]interface{}
}

func AnalzyTxData(contractabi, txData string) (*AnalzyTxDataResp, error) {
	txData = strings.TrimPrefix(txData, "0x")
	///
	///
	contract, err := abi.JSON(strings.NewReader(contractabi))
	if err != nil {
		return nil, err
	}
	//data
	dataByte, err := hex.DecodeString(txData)
	if err != nil {
		return nil, err
	}
	////
	method, err := contract.MethodById(dataByte[:4])
	if err != nil {
		return nil, err
	}
	//
	args := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(args, dataByte[4:])
	if err != nil {
		return nil, err
	}
	///
	atx := &AnalzyTxDataResp{
		MethodId:   hex.EncodeToString(method.ID),
		MethodName: method.RawName,
		Sig:        method.Sig,
		Data:       txData,
		Args:       args,
	}
	return atx, nil
}
