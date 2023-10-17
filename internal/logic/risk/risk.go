package risk

import (
	"context"
	"riskcontral/internal/config"
	"riskcontral/internal/consts"
	"riskcontral/internal/consts/conrisk"
	"riskcontral/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpcmodel"
	"github.com/mpcsdk/mpcCommon/rand"
)

type sRisk struct {
	analzer    *analzyer.Analzyer
	ftruleMap  map[string]*mpcmodel.FtRule
	nftruleMap map[string]*mpcmodel.NftRule
	////
	userControl bool
	txControl   bool
}

func (s *sRisk) PerformRiskTxs(ctx context.Context, userId string, signTx string) (string, int32) {
	g.Log().Debug(ctx, "PerformRiskTxs:", "userId:", userId, "signTx:", signTx, s.txControl)
	///
	if !s.txControl {
		return "", consts.RiskCodePass
	}
	///
	riskserial := rand.GenNewSid()
	///
	code, err := s.checkTxs(ctx, signTx)
	if err != nil {
		g.Log().Warning(ctx, "PerformRiskTxs:", "checkTxs:", err)
		return riskserial, code
	}
	switch code {
	case consts.RiskCodePass, consts.RiskCodeNeedVerification:
		////if pass, chech tfa forbiddent
		info, err := service.TFA().TFAInfo(ctx, userId)
		if err != nil {
			g.Log().Warning(ctx, "PerformRiskTxs err tfinfo:", userId, signTx, err)
			return "", consts.RiskCodeError
		}
		if info == nil {
			g.Log().Warning(ctx, "PerformRiskTxs err tfinfo:", userId, signTx, err)
			return "", consts.RiskCodeError
		}
		///
		if info != nil && info.MailUpdatedAt != nil {
			befor24h := gtime.Now().Add(BeforH24)
			g.Log().Debug(ctx, "PerformRiskTxs:", "befor24h:", befor24h.String(), "info.MailUpdatedAt:", info.MailUpdatedAt.String())
			if !s.isBefor(info.MailUpdatedAt, befor24h) {
				return "", consts.RiskCodeForbidden
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
		return riskserial, code
	case consts.RiskCodeForbidden:
		return riskserial, consts.RiskCodeForbidden
	case consts.RiskCodeError:
		return riskserial, consts.RiskCodeError
	case consts.RiskCodeNoRiskControl:
		return riskserial, consts.RiskCodePass
	default:
		g.Log().Error(ctx, "PerformRiskTxs:", "code:", code)
		return riskserial, consts.RiskCodeError
	}
}

func (s *sRisk) PerformRiskTFA(ctx context.Context, userId string, riskData *conrisk.RiskTfa) (string, int32) {
	g.Log().Debug(ctx, "PerformRiskTFA:", "userId:", userId, "riskData:", riskData, s.userControl)
	if !s.userControl {
		return "", consts.RiskCodePass
	}
	//
	riskserial := rand.GenNewSid()
	///
	code := consts.RiskCodePass
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
	}
	if err != nil {
		g.Log().Error(ctx, "PerformRiskTFA:", err)
		return riskserial, consts.RiskCodeError
		///, err
	}
	///
	// service.Cache().Set(ctx, riskserial+consts.KEY_RiskUId, userId, 0)
	return riskserial, code
}

var BeforH24 time.Duration

func new() *sRisk {
	///
	s := &sRisk{
		analzer:    analzyer.NewAnalzer(),
		ftruleMap:  map[string]*mpcmodel.FtRule{},
		nftruleMap: map[string]*mpcmodel.NftRule{},
	}
	///
	s.ftruleMap, _ = service.DB().GetFtRules(context.TODO())
	s.nftruleMap, _ = service.DB().GetNftRules(context.TODO())
	///
	abis, err := service.DB().GetAbiAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, a := range abis {
		s.analzer.AddAbi(a.Addr, a.Abi)
	}
	////
	// val, err := gcfg.Instance().Get(context.TODO(), "userRisk.forbiddenTime")
	// if err != nil {
	// 	panic(err)
	// }
	val := config.Config.UserRisk.ForbiddenTime
	// BeforH24, err = gtime.ParseDuration("-24h")
	BeforH24, err = gtime.ParseDuration(val)
	if err != nil {
		panic(err)
	}
	// v, _ := gcfg.Instance().Get(context.Background(), "userRisk.userControl", false)
	s.userControl = config.Config.UserRisk.UserControl
	// v, _ = gcfg.Instance().Get(context.Background(), "userRisk.txControl", false)
	s.txControl = config.Config.UserRisk.TxControl

	return s
}

func init() {

	///
	///
	service.RegisterRisk(new())
}
