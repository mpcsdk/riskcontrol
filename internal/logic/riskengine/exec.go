package riskengine

import (
	"context"
	"riskcontrol/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sRiskEngine) Exec(ctx context.Context, ruleName string, data *model.RiskExecData) (int32, error) {
	for _, rule := range s.TxEnginePool {
		if rule.RuleName != ruleName {
			continue
		}
		select {
		case <-ctx.Done():
			return mpccode.RiskCodeError, mpccode.CodePerformRiskTimeOut("ctx done")
		default:
		}
		g.Log().Debug(ctx, "ExecTx:", rule.RuleName, "IsEnable:", rule.IsEnable)
		if !rule.IsEnable {
			continue
		}
		err, rst := rule.Pool.Execute(map[string]interface{}{
			"SignTxs": data.SignTxs,
			"CurTfa":  data.Context.CurTfa,
			"Context": data.Context,
		}, false)
		//
		if err != nil {
			return mpccode.RiskCodeError, mpccode.CodePerformRiskTimeOut(err.Error())
		}
		///
		if v, ok := rst[rule.RuleName]; ok {
			rst := gconv.Int32(v)
			if rst != mpccode.RiskCodePass {
				return rst, nil
			}
		}
	}
	return mpccode.RiskCodePass, nil
}

func (s *sRiskEngine) ExecTx(ctx context.Context, data *model.RiskExecData) (int32, error) {
	g.Log().Debug(ctx, "ExecTx:", data.SignTxs, "Context:", data.Context, "riskrule:", len(s.TxEnginePool))

	for _, rule := range s.TxEnginePool {
		select {
		case <-ctx.Done():
			return mpccode.RiskCodeError, mpccode.CodePerformRiskTimeOut("ctx done")
		default:
		}
		g.Log().Debug(ctx, "ExecTx:", rule.RuleName, "IsEnable:", rule.IsEnable)
		if !rule.IsEnable {
			continue
		}
		err, rst := rule.Pool.Execute(map[string]interface{}{
			"SignTxs": data.SignTxs,
			"CurTfa":  data.Context.CurTfa,
			"Context": data.Context,
		}, false)
		//
		if err != nil {
			g.Log().Warning(ctx, "ExecTx:", rule.RuleName, "err:", err)
			return mpccode.RiskCodeError, mpccode.CodePerformRiskTimeOut(err.Error())
		}
		///
		g.Log().Debug(ctx, "ExecTx", rule.RuleName, "rst:", rst)
		if v, ok := rst[rule.RuleName]; ok {
			rst := gconv.Int32(v)
			if rst != mpccode.RiskCodePass {
				return rst, nil
			}
		}
	}
	return mpccode.RiskCodePass, nil
}

func (s *sRiskEngine) ExecRule(ctx context.Context, ruleName string, data *model.RiskExecData) (int32, error) {
	g.Log().Debug(ctx, "ExecTx:", data.SignTxs, "Context:", data.Context)

	for _, rule := range s.TxEnginePool {
		if rule.engineName != ruleName {
			continue
		}
		g.Log().Debug(ctx, "ExecTx:", rule.RuleName, "IsEnable:", rule.IsEnable)
		if !rule.IsEnable {
			continue
		}
		////todo: check done
		select {
		case <-ctx.Done():
			return mpccode.RiskCodeError, mpccode.CodePerformRiskTimeOut("ctx done")
		default:
		}
		///

		err, rst := rule.Pool.Execute(map[string]interface{}{
			"SignTxs": data.SignTxs,
			"CurTfa":  data.Context.CurTfa,
			"Context": data.Context,
		}, false)
		//
		if err != nil {
			g.Log().Warning(ctx, "ExecTx:", rule.RuleName, "err:", err)
			return mpccode.RiskCodeError, mpccode.CodePerformRiskTimeOut(err.Error())
		}
		///
		if v, ok := rst[rule.RuleName]; ok {
			rst := gconv.Int32(v)
			if rst != mpccode.RiskCodePass {
				return rst, nil
			}
		}
		return 0, nil
	}
	return mpccode.RiskCodePass, nil
}
