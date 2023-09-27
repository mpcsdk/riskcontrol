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
	"strings"
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
	info, err := service.TFA().TFAInfo(ctx, userId)
	if err != nil {
		return "", consts.RiskCodeError
		//, err
	}
	if info == nil {
		return "", consts.RiskCodeError
		///, gerror.NewCode(consts.CodeTFANotExist)
	}
	///
	befor24h := gtime.Now().Add(BeforH24)
	if !s.isBefor(info.MailUpdatedAt, befor24h) {
		return "", consts.RiskCodeError
		///, nil
	}
	///
	code, err := s.checkTxs(ctx, signTx)
	if err != nil {
		g.Log().Warning(ctx, "PerformRiskTxs:", "checkTxs:", err)
		return riskserial, code
	}

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
				Contract:   strings.ToLower(valrisk["contract"].(string)),
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
	return s
}

var BeforH24 time.Duration
var BeforM1 time.Duration

func init() {
	var err error
	// BeforH24, err = gtime.ParseDuration("-24h")
	//todo:
	BeforH24, err = gtime.ParseDuration("-10m")
	if err != nil {
		panic(err)
	}

	BeforM1, err = gtime.ParseDuration("-1m")
	if err != nil {
		panic(err)
	}

	///
	///
	service.RegisterRisk(new())
}
