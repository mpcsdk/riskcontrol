package risk

import (
	"context"
	analzyer "riskcontral/common/ethtx/analzyer"
	"riskcontral/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *sRisk) checkTxs(ctx context.Context, signTxStr string) (int32, error) {
	signTx, err := s.analzer.SignTx(signTxStr)
	if err != nil {
		return consts.RiskCodeError, err
	}

	for _, tx := range signTx.Txs {

		code, err := s.checkTx(ctx, signTx.Address, tx)
		if err != nil {
			return consts.RiskCodeError, err
		}
		if code != consts.RiskCodePass {
			return code, nil
		}
	}
	return consts.RiskCodePass, nil
}

func (s *sRisk) checkTx(ctx context.Context, from string, riskSignTx *analzyer.SignTxData) (int32, error) {

	if ftrule, ok := s.ftruleMap[riskSignTx.Target]; ok {
		//ft
		ethtx, err := s.analzer.AnalzyTxDataFT(
			riskSignTx.Target,
			riskSignTx,
			ftrule)
		if err != nil {
			g.Log().Warning(ctx, "AnalzyTxDataFT:", riskSignTx, err)
			return consts.RiskCodeNeedVerification, nil
		}
		if ethtx.Value.Cmp(ftrule.Threshold) > 0 {
			g.Log().Notice(ctx, "riskTx.Value > threshold:", ethtx, ftrule.Threshold.String())
			return consts.RiskCodeNeedVerification, nil
		}
		///
		cnt, err := rule_ftcnt(ctx, from, ftrule.Contract, ftrule.MethodName)
		if err != nil {
			g.Log().Warning(ctx, "checTx rule_ftcnt:", riskSignTx, err)
			return consts.RiskCodeError, err
		}

		cnt = cnt.Add(cnt, ethtx.Value)
		if cnt.Cmp(ftrule.Threshold) > 0 {
			g.Log().Warning(ctx, "checTx > threshold:", riskSignTx, cnt.String())
			return consts.RiskCodeNeedVerification, nil
		}
		g.Log().Debug(ctx, "checTx < threshold:", cnt.String(), ethtx, ftrule, riskSignTx)
		return consts.RiskCodePass, nil
	} else if nftrule, ok := s.nftruleMap[riskSignTx.Target]; ok {
		//nft
		cnt, err := rule_nftcnt(ctx, from, nftrule.Contract, nftrule.MethodName)
		if err != nil {
			g.Log().Warning(ctx, "checTx rule_Token:", riskSignTx, err)
			return consts.RiskCodeError, err
		}
		cnt += 1
		if cnt > nftrule.Threshold {
			g.Log().Debug(ctx, "checTx > threshold:", riskSignTx, cnt, nftrule.Threshold)
			return consts.RiskCodeNeedVerification, nil
		}
		g.Log().Debug(ctx, "checTx < threshold:", cnt, nftrule.Threshold, riskSignTx)
		return consts.RiskCodePass, nil

	}
	g.Log().Warning(ctx, "checkTx unkonwo contract:", riskSignTx)
	// return consts.RiskCodeError, gerror.NewCode(gcode.CodeInvalidParameter)
	return consts.RiskCodePass, nil
}
