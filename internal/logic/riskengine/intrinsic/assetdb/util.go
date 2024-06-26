package assetdb

import (
	"riskcontrol/internal/logic/riskengine/intrinsic/slice"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *IntrinsicAssetDB) mulChainAssets(chainIds *slice.NumberList, asset *AssetAttr) []*AssetAttr {
	assets := []*AssetAttr{asset}
	if asset.AssetKind == "erc721" || asset.AssetKind == "erc1155" || asset.AssetKind == "external" {
		return assets
	}
	/////
	var chains *slice.NumberList = chainIds
	if chainIds.Len() == 1 {
		if chainIds.Exist(0) {
			chains = s.allChainIds
		}
	}

	////erc20
	for _, id := range *chains {
		if id == asset.ChainId {

		} else {
			if _, ok := s.chainsAsset[id]; ok {
				if _, ok := s.chainsAsset[id][asset.Symbol]; ok {
					assets = append(assets, s.chainsAsset[id][asset.Symbol])
				} else {
					g.Log().Warning(s.ctx, "FtSendSum mulchain have no asset:", "chainId:", id, "asset:", asset.Symbol)
				}
			} else {
				g.Log().Warning(s.ctx, "FtSendSum mulchain have no asset:", "chainId:", id, "asset:", asset.Symbol)
			}
		}
	}
	return assets
}
