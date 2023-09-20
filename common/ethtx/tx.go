package ethtx

import (
	"encoding/hex"

	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type jsonrpcCall struct {
	ID      int64
	Jsonrpc string
	Method  string
	Params  []interface{}
}

func JsonrpcCall(method string, params []interface{}) *jsonrpcCall {
	return &jsonrpcCall{
		ID:      1,
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
	}
}

type JsonrpcResp struct {
	Jsonrpc string
	ID      int64
	Method  string
	Args    []interface{}
	// The result is unmarshaled into this field. Result must be set to a
	// non-nil pointer value of the desired type, otherwise the response will be
	// discarded.
	Result interface{}
	// Error is set if the server returns an error for this request, or if
	// unmarshaling into Result fails. It is not set for I/O errors.
	Error error
}

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
