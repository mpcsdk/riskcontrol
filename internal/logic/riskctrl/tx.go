package riskserver

import (
	"context"
	"riskcontrol/internal/model"
	"riskcontrol/internal/service"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

func (s *sRiskCtrl) isSwapOrMaket(signtx *analzyer.SignTx, chainId int64) bool {
	///
	if market, ok := s.nftMarketMap[chainId]; ok {
		if market == signtx.Target.String() {
			return true
		}
	}
	////
	if swap, ok := s.swapMap[chainId]; ok {
		if swap == signtx.Target.String() {
			return true
		}
	}
	return false
}

func (s *sRiskCtrl) RiskCtrlTx(ctx context.Context, userId string, tfaInfo *entity.Tfa, signDataStr string, chainId string) (int32, error) {
	g.Log().Notice(ctx, "RiskCtrlTx:", "userId:", userId, "signDataStr:", signDataStr, "ChainId:", chainId)
	//trace
	if tfaInfo == nil {
		return mpccode.RiskCodeError, mpccode.CodeInternalError("tfaInfo")
	}
	////is domain
	if strings.Index(signDataStr, "domain") != -1 {
		if tfaInfo.TxNeedVerify {
			return mpccode.RiskCodeNeedVerification, nil
		}
		return mpccode.RiskCodePass, nil
	}
	///analzy txdata
	signData, err := analzyer.DeSignData(signDataStr)
	g.Log().Debug(ctx, "RiskCtrl:DeSignData:", signData)
	if err != nil {
		g.Log().Warning(ctx, "RiskCtrlTx:", "signData:", signData, "err:", err)
		return mpccode.RiskCodeError, mpccode.CodeTxsInvalid()
	}
	/////
	chainAnalzer := s.chainAnalzer[int64(signData.ChainId)]
	if chainAnalzer == nil {
		g.Log().Warning(ctx, "RiskCtrlTx:", "signData:", signData, "err:", err)
		return mpccode.RiskCodeError, mpccode.CodeTxsInvalid()
	}
	/////
	txs := []*analzyer.AnalzyedSignTx{}
	for _, signtx := range signData.Txs {
		if contractRule, ok := chainAnalzer.FtruleMap[signtx.Target.String()]; ok {
			tx, err := chainAnalzer.Analzer.AnalzySignTx(signtx, contractRule)
			if err != nil {
				g.Log().Warning(ctx, "RiskCtrlTx AnalzySignTx err:", "chain:", signData.ChainId, "contract:", contractRule.ContractAddress, "err:", err)
				return mpccode.RiskCodeError, mpccode.CodeTxContractAbiInvalid(signtx.Target.String())
			}
			if tx.From == "" {
				tx.From = signData.Address.String()
			}
			txs = append(txs, tx)
		} else if contractRule, ok := chainAnalzer.NftruleMap[signtx.Target.String()]; ok {
			tx, err := chainAnalzer.Analzer.AnalzySignTx(signtx, contractRule)
			if err != nil {
				return mpccode.RiskCodeError, mpccode.CodeTxContractAbiInvalid(signtx.Target.String())
			}
			if tx.From == "" {
				tx.From = signData.Address.String()
			}
			txs = append(txs, tx)
		} else if signtx.Data == "0x" {
			tx := &analzyer.AnalzyedSignTx{
				///notice: native coin has no target
				From:  signData.Address.String(),
				To:    signtx.Target.String(),
				Value: (*analzyer.BigInt)(signtx.Value.BigInt()),
			}
			txs = append(txs, tx)
		} else {
			tx := &analzyer.AnalzyedSignTx{
				Target: signtx.Target.String(),
				// From:   signData.Address.String(),
				// To:     signtx.Target.String(),
				// Value:  (*analzyer.BigInt)(signtx.Value.BigInt()),
			}
			txs = append(txs, tx)
			/// is swap or market or other tx
			g.Log().Warning(ctx, "RiskCtrlTx unassorted tx", signtx)
		}
	}

	g.Log().Debug(ctx, "RiskCtrl:ExecTx:", txs)
	////perform riskscript
	ok, err := service.RiskEngine().ExecTx(ctx, &model.RiskExecData{
		SignTxs: txs,
		Context: &model.RiskContext{
			ChainId: signData.ChainId,
			CurTfa: func() *entity.Tfa {
				if tfaInfo == nil {
					return &entity.Tfa{}
				} else {
					return tfaInfo
				}
			}(),
		},
	})
	if err != nil {
		return mpccode.RiskCodeError, err
	}
	if ok != mpccode.RiskCodePass {
		return ok, nil
	}
	// }
	return mpccode.RiskCodePass, nil
}
