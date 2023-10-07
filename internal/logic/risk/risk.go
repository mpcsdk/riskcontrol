package risk

import (
	"context"
	"encoding/json"
	"fmt"
	"riskcontral/common"
	analyzsigndata "riskcontral/common/ethtx/analyzSignData"
	"riskcontral/internal/consts"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gtime"
)

type sRisk struct {
	analzer         *analyzsigndata.Analzyer
	contractRiskMap map[string]*contractRisk
}

func (s *sRisk) PerformRiskTxs(ctx context.Context, userId string, signTx string) (string, int32) {
	//
	g.Log().Debug(ctx, "PerformRiskTxs:", "userId:", userId, "signTx:", signTx)
	///
	///
	riskserial := common.GenNewSid()
	///
	code, err := s.checkTxs(ctx, signTx)
	if err != nil {
		g.Log().Warning(ctx, "PerformRiskTxs:", "checkTxs:", err)
		return riskserial, code
	}
	if code == consts.RiskCodeNoRiskControl {
		return riskserial, consts.RiskCodePass
	}
	////
	info, err := service.TFA().TFAInfo(ctx, userId)
	if err != nil || info == nil {
		g.Log().Warning(ctx, "PerformRiskTxs: tfinfo:", userId, signTx, err)
		// return "", consts.RiskCodeNeedVerification
		//, err
	}
	///
	if info != nil && info.MailUpdatedAt != nil {
		befor24h := gtime.Now().Add(BeforH24)
		g.Log().Debug(ctx, "PerformRiskTxs:", "befor24h:", befor24h.String(), "info.MailUpdatedAt:", info.MailUpdatedAt.String())
		if !s.isBefor(info.MailUpdatedAt, befor24h) {
			return "", consts.RiskCodeForbidden
			///, nil
		}
	}
	///
	if info != nil && info.PhoneUpdatedAt != nil {
		befor24h := gtime.Now().Add(BeforH24)
		g.Log().Debug(ctx, "PerformRiskTxs:", "befor24h:", befor24h.String(), "info.PhoneUpdatedAt:", info.PhoneUpdatedAt.String())
		if !s.isBefor(info.PhoneUpdatedAt, befor24h) {
			return "", consts.RiskCodeForbidden
			///, nil
		}
	}
	///
	//
	g.Log().Debug(ctx, "PerformRiskTxs:",
		"userId:", userId,
		"riskseial:", riskserial, "code:", code, err)
	service.Cache().Set(ctx, riskserial+consts.KEY_RiskUId, userId, 0)
	return riskserial, code
}

func (s *sRisk) PerformRiskTFA(ctx context.Context, userId string, riskData *conrisk.RiskTfa) (string, int32) {
	g.Log().Debug(ctx, "PerformRiskTFA:", "userId:", userId, "riskData:", riskData)
	//
	riskserial := common.GenNewSid()
	///
	code := consts.RiskCodeError
	var err error
	///
	switch riskData.Kind {
	case consts.KEY_TFAKindUpPhone:
		code, err = s.checkTFAUpPhone(ctx, userId)
	case consts.KEY_TFAKindUpMail:
		code, err = s.checkTfaUpMail(ctx, userId)
	case consts.KEY_TFAKindCreate:
		code, err = s.checkTfaCreate(ctx, userId)
	default:
		g.Log().Error(ctx, "PerformRiskTFA:", "kind:", riskData.Kind, "not support")
		return riskserial, consts.RiskCodeError
		//, gerror.NewCode(consts.CodePerformUnKnowRisKind)
	}
	if err != nil {
		g.Log().Error(ctx, "PerformRiskTFA:", err)
		return riskserial, consts.RiskCodeError
		///, err
	}
	///
	service.Cache().Set(ctx, riskserial+consts.KEY_RiskUId, userId, 0)
	return riskserial, code
}

var BeforH24 time.Duration

func new() *sRisk {
	///
	s := &sRisk{
		analzer:         analyzsigndata.NewAnalzer(),
		contractRiskMap: map[string]*contractRisk{},
	}
	////
	ctx := context.Background()
	riskcfg, err := gcfg.Instance().Get(ctx, "contractRisk")
	if err != nil {
		panic(err)
	}
	/////
	for _, val := range riskcfg.Array() {
		if valrisk, ok := val.(map[string]interface{}); !ok {
			panic(fmt.Errorf("contractRisk:%v", val))
		} else {
			Threshold, _ := valrisk["threshold"].(json.Number).Int64()
			r := &contractRisk{
				Contract:   valrisk["contract"].(string),
				Kind:       valrisk["kind"].(string),
				MethodName: valrisk["methodName"].(string),
				Threshold:  int(Threshold),
			}
			s.contractRiskMap[r.Contract] = r
		}
	}
	////todo: get abi
	for _, contract := range s.contractRiskMap {
		abistr, err := service.DB().GetAbi(ctx, contract.Contract)
		if err != nil {
			continue
		}
		///
		s.analzer.AddAbi(contract.Contract, abistr)
	}
	//

	val, err := gcfg.Instance().Get(ctx, "userRisk.forbiddenTime")
	if err != nil {
		panic(err)
	}
	// BeforH24, err = gtime.ParseDuration("-24h")
	BeforH24, err = gtime.ParseDuration(val.String())
	if err != nil {
		panic(err)
	}

	return s
}

func init() {

	///
	///
	service.RegisterRisk(new())
}
