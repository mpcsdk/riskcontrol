package db

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/franklihub/mpcCommon/mpcmodel"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

var NftRules = map[string]*mpcmodel.NftRule{}

var FtRules = map[string]*mpcmodel.FtRule{}

func (s *sDB) GetNftRules(ctx context.Context) (map[string]*mpcmodel.NftRule, error) {
	return NftRules, nil
}

func (s *sDB) GetFtRules(ctx context.Context) (map[string]*mpcmodel.FtRule, error) {
	return FtRules, nil
}
func init() {
	cfg := gcfg.Instance()
	v, err := cfg.Get(context.Background(), "txRisk.nft")
	if err != nil {
		panic(err)
	}
	nftrules := []*mpcmodel.NftRule{}
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
	ftrules := []*mpcmodel.FtRule{}
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
