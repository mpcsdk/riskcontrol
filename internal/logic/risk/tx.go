package risk

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sRisk) checkTxs(ctx context.Context, signTxStr string) (int32, error) {
	signTx, err := s.analzer.SignTx(signTxStr)
	if err != nil {
		return mpccode.RiskCodeError, err
	}

	for _, tx := range signTx.Txs {
		code, err := s.checkTx(ctx, signTx.Address, tx)
		if err != nil {
			err = gerror.Wrap(err, mpccode.ErrDetails(
				mpccode.ErrDetail("address", signTx.Address),
				mpccode.ErrDetail("tx", tx),
			))
			return mpccode.RiskCodeError, err
		}
		if code != mpccode.RiskCodePass {
			return code, nil
		}
	}
	return mpccode.RiskCodePass, nil
}

func (s *sRisk) checkTx(ctx context.Context, from string, riskSignTx *analzyer.SignTxData) (int32, error) {
	riskSignTx.Target = strings.ToLower(riskSignTx.Target)
	if ftrule, ok := s.ftruleMap[riskSignTx.Target]; ok {
		//ft
		ethtx, err := s.analzer.AnalzyTxDataFT(
			riskSignTx.Target,
			riskSignTx,
			ftrule)
		if err != nil {
			g.Log().Warning(ctx, "checkTx:", "riskSignTx:", riskSignTx)
			g.Log().Warning(ctx, "checkTx:", "ftrule:", ftrule)
			g.Log().Errorf(ctx, "%+v", err)
			return mpccode.RiskCodeNeedVerification, nil
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
			err = gerror.Wrap(err, mpccode.ErrDetails(
				mpccode.ErrDetail("from", from),
				mpccode.ErrDetail("contract", ftrule.Contract),
				mpccode.ErrDetail("methodName", ftrule.MethodName),
			))
			return mpccode.RiskCodeError, err
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
	} else if nftrule, ok := s.nftruleMap[riskSignTx.Target]; ok {
		ethtx, err := s.analzer.AnalzyTxDataNFT(
			riskSignTx.Target,
			riskSignTx,
			nftrule)
		if err != nil {
			g.Log().Warning(ctx, "checkTx:", "riskSignTx:", riskSignTx)
			g.Log().Warning(ctx, "checkTx:", "nftrule:", nftrule)
			g.Log().Errorf(ctx, "%+v", err)
			return mpccode.RiskCodeNeedVerification, nil
		}
		if ethtx.MethodName != nftrule.MethodName {
			g.Log().Warning(ctx, "checkTx:", "methodName unmath:", ethtx.MethodName, nftrule.MethodName)
			return mpccode.RiskCodeNoRiskControl, nil
		}
		//nft
		cnt, err := rule_nftcnt(ctx, from, nftrule.Contract, nftrule.MethodName)
		if err != nil {
			err = gerror.Wrap(err, mpccode.ErrDetails(
				mpccode.ErrDetail("from", from),
				mpccode.ErrDetail("contract", nftrule.Contract),
				mpccode.ErrDetail("methodName", nftrule.MethodName),
			))
			return mpccode.RiskCodeError, err
		}

		cnt += 1
		if cnt > nftrule.Threshold {
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
