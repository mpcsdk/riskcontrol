package ethmodel

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type StringList []string
type Log struct {
	BlockNumber      int64      `bson:"blockNumber"`
	BlockHash        string     `bson:"blockHash"`
	TransactionHash  string     `bson:"transactionHash"`
	TransactionIndex int        `bson:"transactionIndex"`
	Address          string     `bson:"address"`
	Topics           StringList `bson:"topics"`
	Data             string     `bson:"data"`
	LogIndex         int        `bson:"logIndex"`
	Removed          bool       `bson:"removed"`
}

// type EthLog struct {
// 	Address string   `json:"address"`
// 	Topics  []string `json:"topics"`
// 	Data    string   `json:"data"`

// 	BlockNumber hexutil.Uint `json:"blockNumber"`
// 	TxHash      string       `json:"transactionHash"`
// 	TxIndex     hexutil.Uint `json:"transactionIndex"`
// 	BlockHash   string       `json:"blockHash"`
// 	Index       hexutil.Uint `json:"logIndex"`
// 	Removed     bool         `json:"removed"`
// }

// /
type EthFilterQuery struct {
	BlockHash *common.Hash     // used by eth_getLogs, return logs only from block with this hash
	FromBlock *rpc.BlockNumber // beginning of the queried range, nil means genesis block
	ToBlock   *rpc.BlockNumber // end of the range, nil means latest block
	Address   []common.Address // restricts matches to events created by specific contracts
	Topics    []common.Hash
}
