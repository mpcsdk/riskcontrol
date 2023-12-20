package risk

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

var defaultSceneNo = "0"

func (s *sRisk) checkTxs(ctx context.Context, signTxStr string) (int32, error) {
	scencRule := s.getContractRules(ctx, defaultSceneNo)
	if scencRule == nil {
		return 0, nil
	}
	///
	signTx, err := scencRule.analzer.SignTx(signTxStr)
	if err != nil {
		g.Log().Warning(ctx, "checkTxs:", "signTxStr:", signTxStr, "err:", err)
		return mpccode.RiskCodeError, mpccode.CodePerformRiskError()
	}

	for _, tx := range signTx.Txs {
		code, err := s.checkTx(ctx, signTx.Address.String(), tx)
		if err != nil {
			return mpccode.RiskCodeError, err
		}
		if code != mpccode.RiskCodePass {
			return code, nil
		}
	}
	return mpccode.RiskCodePass, nil
}

func (s *sRisk) checkTx(ctx context.Context, from string, riskSignTx *analzyer.SignTxData) (int32, error) {
	scencRule := s.getContractRules(ctx, defaultSceneNo)
	if scencRule == nil {
		return mpccode.RiskCodePass, nil
	}
	///
	// riskSignTx.Target = strings.ToLower(riskSignTx.Target)
	if ftrule, ok := scencRule.ftruleMap[riskSignTx.Target.String()]; ok {
		//ft
		ethtx, err := scencRule.analzer.AnalzyTxDataFT(
			riskSignTx.Target.String(),
			riskSignTx,
			ftrule)
		if err != nil {
			g.Log().Warning(ctx, "checkTx:", "riskSignTx:", riskSignTx, "err:", err)
			return mpccode.RiskCodeError, mpccode.CodePerformRiskError()
		}
		if ethtx == nil {
			g.Log().Warning(ctx, "checkTx ethtx is nil:", "riskSignTx:", riskSignTx)
			return mpccode.RiskCodeNoRiskControl, nil
		}
		if ethtx.MethodName != ftrule.MethodName {
			g.Log().Warning(ctx, "checkTx:", "methodName unmath:", ethtx.MethodName, ftrule.MethodName)
			return mpccode.RiskCodeNoRiskControl, nil
		}
		if ethtx.Value.Cmp(ftrule.Threshold) > 0 {
			g.Log().Notice(ctx, "riskTx.Value > threshold:", ethtx, ftrule.Threshold.String())
			return mpccode.RiskCodeNeedVerification, nil
		}
		///
		cnt, err := rule_ftcnt(ctx, from, ftrule.Contract, ftrule.MethodName)
		if err != nil {
			g.Log().Warning(ctx, "checkTx:", "from:", from,
				"contract:", ftrule.Contract,
				"methodName:", ftrule.MethodName, "err:", err)
			return mpccode.RiskCodeError, mpccode.CodePerformRiskError()
		}

		cnt = cnt.Add(cnt, ethtx.Value)
		if cnt.Cmp(ftrule.Threshold) > 0 {
			g.Log().Notice(ctx, "riskTx.Value > threshold:", "ethtx:", ethtx,
				"txcnt:", cnt.String(),
				"threshold:", ftrule.Threshold.String())
			return mpccode.RiskCodeNeedVerification, nil
		}
		g.Log().Notice(ctx, "riskTx.Value < threshold:", "ethtx:", ethtx,
			"txcnt:", cnt.String(),
			"threshold:", ftrule.Threshold.String())
		return mpccode.RiskCodePass, nil
	} else if nftrule, ok := scencRule.nftruleMap[riskSignTx.Target.String()]; ok {
		ethtx, err := scencRule.analzer.AnalzyTxDataNFT(
			riskSignTx.Target.String(),
			riskSignTx,
			nftrule)
		if err != nil {
			g.Log().Warning(ctx, "checkTx:", "riskSignTx:", riskSignTx, "err:", err)
			return mpccode.RiskCodeError, mpccode.CodePerformRiskError()
		}
		if ethtx == nil {
			g.Log().Warning(ctx, "checkTx ethtx is nil:", "riskSignTx:", riskSignTx)
			return mpccode.RiskCodeNoRiskControl, nil
		}
		if ethtx.MethodName != nftrule.MethodName {
			g.Log().Warning(ctx, "checkTx:", "methodName unmath:", ethtx.MethodName, nftrule.MethodName)
			return mpccode.RiskCodeNoRiskControl, nil
		}
		//nft
		cnt, err := rule_nftcnt(ctx, from, nftrule.Contract, nftrule.MethodName)
		if err != nil {
			g.Log().Warning(ctx, "checkTx:", "from:", from,
				"contract:", ftrule.Contract,
				"methodName:", ftrule.MethodName, "err:", err)
			return mpccode.RiskCodeError, mpccode.CodePerformRiskError()
		}

		cnt += 1
		if cnt > int(nftrule.ThresholdNft) {
			g.Log().Notice(ctx, "riskTx.Value > threshold:", "ethtx:", ethtx,
				"txcnt:", cnt,
				"threshold:", nftrule.Threshold)
			return mpccode.RiskCodeNeedVerification, nil
		}
		g.Log().Notice(ctx, "checTx < threshold:", "ethtx:", ethtx,
			"txcnt:", cnt,
			"threshold:", nftrule.Threshold)
		return mpccode.RiskCodePass, nil

	}
	g.Log().Warning(ctx, "checkTx unkonwo contract:", riskSignTx)
	return mpccode.RiskCodeNoRiskControl, nil
}
