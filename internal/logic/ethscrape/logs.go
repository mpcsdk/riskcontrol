package ethscrape

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"riskcontral/common/ethtx"
	"riskcontral/common/ethtx/analzyer"
	"riskcontral/internal/dao"
	"riskcontral/internal/model"
	"riskcontral/internal/model/do"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
	// "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/ethclient"
)

type sEthEventGeter struct {
	ctx     context.Context
	chainId string
	url     string

	////
	lastBlock int64
	abis      []*entity.ContractAbi
	nftRules  map[string]*model.NftRule
	ftRules   map[string]*model.FtRule
	////
	// ethclient *ethclient.Client
	// client    *http.Client
	chainNode *ethtx.ChainNode
	analzer   *analzyer.Analzyer
}

var aggBlock = int64(1000)

func init() {
	ctx := gctx.GetInitCtx()
	cfg := gcfg.Instance()
	v, _ := cfg.Get(ctx, "userRisk.aggBlock", 100)
	aggBlock = v.Int64()

	///
	service.RegisterEthEventGeter(newEthEventGeter())
}
func newEthEventGeter() *sEthEventGeter {
	cfg := gcfg.Instance()
	url, err := cfg.Get(context.Background(), "ethRpc")
	if err != nil {
		panic(err)
	}
	chainId, err := cfg.Get(context.Background(), "chainId")
	if err != nil {
		panic(err)
	}

	s := &sEthEventGeter{
		ctx:     gctx.GetInitCtx(),
		url:     url.String(),
		chainId: chainId.String(),
	}

	s.chainNode = ethtx.NewChainNode(url.String())
	///

	s.analzer = analzyer.NewAnalzer()

	return s
}
func (s *sEthEventGeter) InitByService() error {
	stat, err := service.DB().GetScrapeStat(s.ctx, s.chainId)
	if err != nil {
		return err
	}
	s.lastBlock = stat.LastBlock
	if s.lastBlock == rpc.LatestBlockNumber.Int64() {
		number, err := s.GetBlockNumber(s.ctx)
		if err != nil {
			return err
		}
		s.lastBlock = number - 1
		err = service.DB().SetScrapeStat(s.ctx, s.chainId, rpc.BlockNumber(number))
		if err != nil {
			return err
		}
	}

	////
	abis, err := service.DB().GetAbiAll(s.ctx)
	if err != nil {
		return err
	}
	s.abis = abis
	///
	nftRules, err := service.DB().GetNftRules(s.ctx)
	if err != nil {
		return err
	}
	ftRules, err := service.DB().GetFtRules(s.ctx)
	if err != nil {
		return err
	}
	s.nftRules = nftRules
	s.ftRules = ftRules
	return nil
}
func (s *sEthEventGeter) Stop() {
	s.ctx.Done()
}
func (s *sEthEventGeter) RunBySerivce() {
	s.Run(s.lastBlock, s.abis, s.nftRules, s.ftRules)
}
func (s *sEthEventGeter) Run(
	lastBlock int64,
	abis []*entity.ContractAbi,
	nftRules map[string]*model.NftRule,
	ftRules map[string]*model.FtRule,
) {
	s.lastBlock = lastBlock

	///
	for _, abistr := range abis {
		s.analzer.AddAbi(abistr.Addr, abistr.Abi)
	}
	//
	addresses := []common.Address{}
	topicmap := map[string]common.Hash{}
	for _, rule := range nftRules {
		addresses = append(addresses, common.HexToAddress(rule.Contract))
		topicmap[rule.EventTopic] = common.HexToHash(rule.EventTopic)
	}
	for _, rule := range ftRules {
		addresses = append(addresses, common.HexToAddress(rule.Contract))
		topicmap[rule.EventTopic] = common.HexToHash(rule.EventTopic)
	}
	///
	topics := [][]common.Hash{}
	topic := []common.Hash{}

	for _, v := range topicmap {
		topic = append(topic, v)
	}
	topics = append(topics, topic)
	fmt.Println(topics)
	//
	go func() {
		for {
			///
			time.Sleep(15 * time.Second)
			number, err := s.GetBlockNumber(s.ctx)
			if err != nil {
				g.Log().Error(s.ctx, "GetBlockNumber:", err)
			}
			if number > s.lastBlock {

				logs, err := s.GetLogs(
					s.ctx,
					big.NewInt(s.lastBlock+1),
					big.NewInt(number),
					addresses,
					topics,
				)
				if err != nil {
					g.Log().Error(s.ctx, "GetLogs:", err, s.lastBlock, number, addresses, topic)
					continue
				}
				g.Log().Notice(s.ctx, "GetLogs:", s.lastBlock+1, number, len(logs))
				///set scrape stat
				s.lastBlock = number
				err = service.DB().SetScrapeStat(s.ctx, s.chainId, rpc.
					BlockNumber(number))
				if err != nil {
					g.Log().Error(s.ctx, "SetScrapeStat:", err)
				}
				if len(logs) == 0 {
					continue
				}
				///
				ftxs := s.FilterTx(
					s.ctx,
					logs,
					nftRules,
					ftRules,
				)
				if len(ftxs) == 0 {
					continue
				}
				err = service.DB().InsertTxs(s.ctx, ftxs)
				if err != nil {
					g.Log().Warning(s.ctx, "InsertTxs:", err)
					continue
				}
				///agg , ft
				for _, ft := range ftxs {
					if ft.Kind == "ft" {
						////
						agg, err := s.aggTxFT_24H(s.ctx, ft)
						if err != nil {
							g.Log().Warning(s.ctx, "aggTxFT_24H:", err, ft)
							continue
						}
						err = service.DB().UpAggFT(s.ctx, agg)
						if err != nil {
							g.Log().Warning(s.ctx, "UpAggFT:", err)
							continue
						}
					} else if ft.Kind == "nft" {
						agg := s.aggTxNFT_24H(s.ctx, ft)
						err = service.DB().UpAggNFT(s.ctx, agg)
						if err != nil {
							g.Log().Warning(s.ctx, "UpAggFT:", err)
						}
					}

				}
			}
		}
	}()
}

var errEmpty error = errors.New("empty db")

func (s *sEthEventGeter) aggTxFT_24H(ctx context.Context, ft *entity.EthTx) (*entity.AggFt24H, error) {
	rst, err := dao.EthTx.Ctx(ctx).Where(
		&do.AggFt24H{
			From:       ft.From,
			Contract:   ft.Contract,
			MethodName: ft.MethodName,
		}).
		WhereGTE(dao.EthTx.Columns().BlockNumber, ft.BlockNumber-aggBlock).
		WhereLTE(dao.EthTx.Columns().BlockNumber, ft.BlockNumber).
		OrderAsc(dao.EthTx.Columns().BlockNumber).
		Fields(
			dao.EthTx.Columns().From,
			dao.EthTx.Columns().To,
			dao.EthTx.Columns().Value,
			dao.EthTx.Columns().Contract,
			dao.EthTx.Columns().MethodName,
			dao.EthTx.Columns().MethodSig,
			dao.EthTx.Columns().BlockNumber,
		).
		All()

	if rst.IsEmpty() {
		return nil, errEmpty
	}
	if err != nil {
		return nil, err
	}
	//
	data := []*entity.EthTx{}
	vals := big.NewInt(0)
	///
	err = rst.Structs(&data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errEmpty
	}
	///
	for _, tx := range data {
		i, ok := big.NewInt(0).SetString(tx.Value, 10)
		if ok {
			vals = vals.Add(vals, i)
		} else {

		}
	}
	tx := data[0]
	e := &entity.AggFt24H{
		From:       tx.From,
		To:         tx.To,
		Value:      decimal.NewFromBigInt(vals, 0),
		Contract:   tx.Contract,
		UpdatedAt:  gtime.Now(),
		MethodName: tx.MethodName,
		MethodSig:  tx.MethodSig,
		FromBlock:  data[0].BlockNumber,
		ToBlock:    data[len(data)-1].BlockNumber,
	}
	return e, nil
}
func (s *sEthEventGeter) aggTxNFT_24H(ctx context.Context, nft *entity.EthTx) *entity.AggNft24H {

	return nil
}
func (s *sEthEventGeter) GetBlockNumber(ctx context.Context) (int64, error) {
	number, err := s.chainNode.EthClient().BlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	return int64(number), nil
}

func (s *sEthEventGeter) GetLogs(ctx context.Context,
	from, to *big.Int,
	addresses []common.Address,
	topic [][]common.Hash) ([]*types.Log, error) {
	logs, err := s.chainNode.EthClient().FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: addresses,
		Topics:    topic,
		FromBlock: from,
		ToBlock:   to,
	})

	if err != nil {
		return nil, err
	}
	rstLogs := []*types.Log{}
	for i, _ := range logs {
		rstLogs = append(rstLogs, &logs[i])
	}

	return rstLogs, err
}

func (s *sEthEventGeter) FilterTx(
	ctx context.Context,
	logs []*types.Log,
	nftrules map[string]*model.NftRule,
	ftrules map[string]*model.FtRule,
) []*entity.EthTx {

	filter := []*entity.EthTx{}
	for _, log := range logs {
		if r, ok := nftrules[strings.ToLower(log.Address.Hex())]; ok {
			tx, err := s.analzer.AnalzyLogNFT(log.Address.String(), log, r)
			if tx == nil {
				continue
			}
			g.Log().Debug(ctx, "FilterTx nft:", log, "tx:", tx)
			if err != nil {
				g.Log().Warning(ctx, "FilterTx err:", err, "log:", log)
				continue
			}

			filter = append(filter, tx)
		} else if r, ok := ftrules[strings.ToLower(log.Address.Hex())]; ok {
			tx, err := s.analzer.AnalzyLogFT(log.Address.String(), log, r)
			g.Log().Debug(ctx, "FilterTx ft:", log, "tx:", tx)
			if tx == nil {
				continue
			}
			if err != nil {
				g.Log().Warning(ctx, "FilterTx err:", err, "log:", log)
				continue
			}

			filter = append(filter, tx)
		} else {
			g.Log().Notice(ctx, "not match:", log)
		}
	}
	return filter
	//
}
