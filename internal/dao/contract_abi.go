// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"riskcontral/internal/dao/internal"
)

// internalContractAbiDao is internal type for wrapping internal DAO implements.
type internalContractAbiDao = *internal.ContractAbiDao

// contractAbiDao is the data access object for table contract_abi.
// You can define custom methods on it to extend its functionality as you wish.
type contractAbiDao struct {
	internalContractAbiDao
}

var (
	// ContractAbi is globally public accessible object for table contract_abi operations.
	ContractAbi = contractAbiDao{
		internal.NewContractAbiDao(),
	}
)

// Fill with you ideas below.