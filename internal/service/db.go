// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"riskcontral/internal/model"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"

	"github.com/ethereum/go-ethereum/rpc"
)

type (
	IDB interface {
		GetAbi(ctx context.Context, addr string) (string, error)
		GetAbiAll(ctx context.Context) ([]*entity.ContractAbi, error)
		// /
		UpAggFT(ctx context.Context, ft *entity.AggFt24H) error
		UpAggNFT(ctx context.Context, nft *entity.AggNft24H) error
		GetAggFT(ctx context.Context, from, contract, methodName string) (*entity.AggFt24H, error)
		GetRules(ctx context.Context, ruleId string) (string, error)
		AllRules(ctx context.Context) map[string]string
		GetNftRules(ctx context.Context) (map[string]*model.NftRule, error)
		GetFtRules(ctx context.Context) (map[string]*model.FtRule, error)
		GetScrapeStat(ctx context.Context, chainId string) (*entity.ScrapeLogsStat, error)
		SetScrapeStat(ctx context.Context, chainId string, nr rpc.BlockNumber) error
		InsertTfaInfo(ctx context.Context, userId string, data *do.Tfa) error
		// //
		UpdateTfaInfo(ctx context.Context, userId string, data *do.Tfa) error
		FetchTfaInfo(ctx context.Context, userId string) (*entity.Tfa, error)
		// /
		InsertTx(ctx context.Context, d *entity.EthTx) error
		// /
		InsertTxs(ctx context.Context, txs []*entity.EthTx) error
	}
)

var (
	localDB IDB
)

func DB() IDB {
	if localDB == nil {
		panic("implement not found for interface IDB, forgot register?")
	}
	return localDB
}

func RegisterDB(i IDB) {
	localDB = i
}
