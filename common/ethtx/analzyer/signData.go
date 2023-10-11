package analzyer

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Analzyer struct {
	abis       map[string]string
	abiStructs map[string]*abi.ABI
}

func NewAnalzer() *Analzyer {
	return &Analzyer{
		abis:       map[string]string{},
		abiStructs: map[string]*abi.ABI{},
	}
}

func (s *Analzyer) AddAbi(contract string, abistr string) error {
	contract = strings.ToLower(contract)
	s.abis[contract] = abistr

	abiabi, err := abi.JSON(strings.NewReader(abistr))
	if err != nil {
		return err
	}
	s.abiStructs[contract] = &abiabi

	return nil
}

// //
