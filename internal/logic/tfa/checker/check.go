package check

import (
	"context"
	v1 "riskcontrol/api/tfa/v1"
	"riskcontrol/internal/logic/tfa/tfaconst"
	"riskcontrol/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type Checker struct {
	forbiddentTime time.Duration
}

func NewChecker(forbiddent time.Duration) *Checker {
	return &Checker{forbiddentTime: forbiddent}
}
func (s *Checker) CheckKind(ctx context.Context, tfaInfo *entity.Tfa, kind tfaconst.RISKKIND, data *v1.RequestData) (int32, error) {
	code := mpccode.RiskCodeError
	var err error = nil
	///
	switch kind {
	case tfaconst.RiskKind_BindPhone:
		if tfaInfo.Phone != "" {
			return code, mpccode.CodeTFAExist()
		}
		///
		code, err = s.CheckTfaBindPhone(ctx, tfaInfo)
	case tfaconst.RiskKind_BindMail:
		if tfaInfo.Mail != "" {
			return code, mpccode.CodeTFAExist()
		}
		////
		code, err = s.CheckTfaBindMail(ctx, tfaInfo)
	case tfaconst.RiskKind_UpPhone:
		if tfaInfo.Phone == "" {
			return code, mpccode.CodeTFANotExist()
		}
		code, err = s.CheckTfaUpPhone(ctx, tfaInfo)
	case tfaconst.RiskKind_UpMail:
		if tfaInfo.Mail == "" {
			return code, mpccode.CodeTFANotExist()
		}
		code, err = s.CheckTfaUpMail(ctx, tfaInfo)
	case tfaconst.RiskKind_TfaRisk:
		if tfaInfo.Mail == "" && tfaInfo.Phone == "" {
			return code, mpccode.CodeTFANotExist()
		}
		code, err = s.CheckTfaRisk(ctx, tfaInfo)

	case tfaconst.RiskKind_TxNeedVerify:
		if tfaInfo.Mail == "" && tfaInfo.Phone == "" {
			return code, mpccode.CodeTFANotExist()
		}
		if tfaInfo.TxNeedVerify == data.Enable {
			return code, mpccode.CodeParamInvalid()
		}
		code, err = s.CheckPersonRisk(ctx, tfaInfo)
	case tfaconst.RiskKind_Tx:
		///
		res, err := service.RiskCtrl().RiskCtrlTx(ctx, data.UserId, tfaInfo, data.SignDataStr, data.ChainId)
		// res, err := service.NrpcClient().RiskTxs(ctx, &riskengine.TxRiskReq{
		// 	UserId:  data.UserId,
		// 	SignTx:  data.SignDataStr,
		// 	ChainId: data.ChainId,
		// })
		if err != nil {
			g.Log().Warning(ctx, "RpcRiskTxs:", "data:", data, "err:", err)
			code = 0
		} else {
			code = res
		}
	///
	default:
		return code, mpccode.CodeParamInvalid()
	}
	////
	return code, err
}
