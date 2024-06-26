package assetdb

import (
	"math/big"
	"riskcontrol/internal/logic/riskengine/intrinsic/slice"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/ethtx/ethutil"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

var ZeroBigInt = &analzyer.BigInt{}

func (s *IntrinsicAssetDB) AssetNftCnt(chainIds *slice.NumberList, from string, asset *AssetAttr, ts int64) int64 {
	g.Log().Debug(s.ctx, "AssetNftCnt:", from, asset, ts)
	///721
	if asset.AssetKind != "erc721" && asset.AssetKind != "erc1155" {
		return 0
	}
	////nft unsupport mulchain asset
	// assets := s.mulChainAssets(chainIds, asset)
	///
	total := int64(0)
	// for _, asset := range assets {
	if asset.AssetKind == "erc721" {
		cnt, err := s.enhanced_riskctrl.GetAggCnt(s.ctx, mpcdao.QueryEnhancedRiskCtrlRes{
			From:     from,
			Contract: asset.Contract,
			ChainId:  asset.ChainId,
			StartTs:  ts,
		})
		g.Log().Debug(s.ctx, "AssetNftCnt :", "from:", from, "contract:", asset.Contract, "ts:", ts, "chainId:", asset.ChainId, "cnt:", cnt)
		if err != nil {
			g.Log().Warning(
				s.ctx, "AssetNftCnt:", from, asset.Contract,
				"err:", err,
			)
			return 0
		}
		total += cnt
	} else {
		///1155 need tokenid
		txs, err := s.enhanced_riskctrl.GetAgg(s.ctx, mpcdao.QueryEnhancedRiskCtrlRes{
			From:     from,
			Contract: asset.Contract,
			ChainId:  asset.ChainId,
			TokenId:  asset.TokenId,
			StartTs:  ts,
		})
		if err != nil {
			g.Log().Warning(
				s.ctx, "AssetNftCnt:", from, asset.Contract,
				"err:", err,
			)
			return 0
		}
		for _, tx := range txs {
			val := big.NewInt(0)
			val.SetString(tx.Value, 10)
			g.Log().Debug(s.ctx, "AssetNftCnt :", "from:", from, "contract:", asset.Contract, "ts:", ts, "chainId:", asset.ChainId, "cnt:", val)
			total += val.Int64()
		}
	}
	// }

	g.Log().Debug(s.ctx, "AssetSendCnt:", from, asset, ts, total)
	///
	return total
}

// ///
func (s *IntrinsicAssetDB) AssetFtSum(chainIds *slice.NumberList, from string, asset *AssetAttr, ts int64) float64 {
	g.Log().Debug(s.ctx, "AssetFtSum:", from, asset, ts)
	if asset.AssetKind != "external" && asset.AssetKind != "erc20" {
		return 0
	}
	////to mulchain asset
	assets := s.mulChainAssets(chainIds, asset)
	///
	data := big.NewInt(0)
	sum := float64(0)
	for _, asset := range assets {
		rst := s.ftSendSum(from, asset, ts)
		f := ethutil.BigDecimal2Float64(rst.Int(), asset.Decimal)
		sum += f
	}

	g.Log().Debug(s.ctx, "AssetSendSum:", from, asset, ts, data)
	///
	// f := big.NewFloat(0)
	// d, _ := f.SetInt(data.Int()).Float64()
	return sum
}

// MUD、MAK、USDT、RPG
func (s *IntrinsicAssetDB) ftSendSum(from string, asset *AssetAttr, ts int64) *analzyer.BigInt {
	g.Log().Debug(s.ctx, "FtSendSum:", "from:", from, "asset:", asset, "ts:", ts)
	data := big.NewInt(0)
	if asset.AssetKind != "external" && asset.AssetKind != "erc20" {
		return (*analzyer.BigInt)(data)
	}
	/////
	////native chain
	if asset.Contract == "" {
		////native
		rst, err := s.enhanced_riskctrl.GetAggSum(s.ctx, mpcdao.QueryEnhancedRiskCtrlRes{
			From:     from,
			Contract: asset.Contract,
			ChainId:  asset.ChainId,
			StartTs:  ts,
		})
		g.Log().Debug(s.ctx, "FtSendSum native:", "from:", from, "contract:", asset.Contract, "ts:", ts, "chainId:", asset.ChainId, "rst:", rst)
		if err != nil {
			g.Log().Warning(
				s.ctx, "FtAssetSum:", from, asset.Contract,
				"err:", err,
			)
			return ZeroBigInt
		}
		///
		data = data.Add(data, rst)
	} else {
		if asset.AssetKind != "erc20" {
			return ZeroBigInt
		}
		///contract
		rst, err := s.enhanced_riskctrl.GetAggSum(s.ctx, mpcdao.QueryEnhancedRiskCtrlRes{
			From:     from,
			Contract: asset.Contract,
			ChainId:  asset.ChainId,
			StartTs:  ts,
		})
		g.Log().Debug(s.ctx, "FtSendSum token:", "from:", from, "contract:", asset.Contract, "ts:", ts, "chainId:", asset.ChainId, "rst:", rst)

		if err != nil {
			g.Log().Warning(
				s.ctx, "FtSendSum:", from, asset.Contract,
				"err:", err,
			)
			return ZeroBigInt
		}
		///
		data = data.Add(data, rst)
	}

	g.Log().Info(s.ctx, "FtSendSum:", "from:", from, "asset:", asset, "ts:", ts, "data:", data)

	return (*analzyer.BigInt)(data)
}

type QueryResult struct {
	Status int `json:"status"`
	Result any
}
type QueryAdvRep struct {
	Status int     `json:"status"`
	Result []*item `json:"result"`
}
type item struct {
	Height    string `json:"height"`
	Blockhash string `json:"blockhash"`
	Ts        string `json:"timestamp"`
	Txhash    string `json:"txhash"`
	Toaddr    string `json:"toaddr"`
	Fromaddr  string `json:"fromaddr"`
	Value     string `json:"value"`
	Contract  string `json:"contract"`
	Gas       string `json:"gas"`
	Gasprice  string `json:"gasprice"`
}
