package db

import (
	"context"
	"riskcontral/internal/model"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

var NftRules = map[string]*model.NftRule{}

var FtRules = map[string]*model.FtRule{}

func (s *sDB) GetNftRules(ctx context.Context) (map[string]*model.NftRule, error) {
	return NftRules, nil
}

func (s *sDB) GetFtRules(ctx context.Context) (map[string]*model.FtRule, error) {
	return FtRules, nil
}
func init() {
	cfg := gcfg.Instance()
	v, err := cfg.Get(context.Background(), "txRisk.nft")
	if err != nil {
		panic(err)
	}
	nftrules := []*model.NftRule{}
	err = v.Structs(&nftrules)
	if err != nil {
		panic(err)
	}
	///
	///
	v, err = cfg.Get(context.Background(), "txRisk.ft")
	if err != nil {
		panic(err)
	}
	ftrules := []*model.FtRule{}
	err = v.Structs(&ftrules)
	if err != nil {
		panic(err)
	}
	///tidy
	///
	for _, v := range nftrules {
		v.EventTopic = hexutil.Encode(crypto.Keccak256([]byte(v.EventSig)))
	}
	for _, v := range ftrules {
		v.EventTopic = hexutil.Encode(crypto.Keccak256([]byte(v.EventSig)))
	}
	///
	for _, v := range nftrules {
		NftRules[v.Contract] = v
	}
	for _, v := range ftrules {
		FtRules[v.Contract] = v
	}
	///
	g.Log().Notice(context.Background(),
		"nftRules:", NftRules,
		"ftRules:", FtRules,
	)
}
