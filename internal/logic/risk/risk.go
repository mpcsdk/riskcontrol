package risk

import (
	"context"
	"riskcontral/internal/config"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpccode"
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
	////

	riskStatLock sync.Mutex
	riskStat     map[string]*model.RiskStat
	///
}

func (s *sRisk) RiskTxs(ctx context.Context, userId string, signTx string) (string, int32) {
	///
	if !s.txControl {
		return "", mpccode.RiskCodePass
	}
	///
	riskserial := rand.GenNewSid()
	s.riskStat[riskserial] = &model.RiskStat{
		Kind: model.Kind_RiskTx,
		Type: signTx,
	}
	///
	code, err := s.checkTxs(ctx, signTx)
	if err != nil {
		g.Log().Warning(ctx, "PerformRiskTxs:", "userId:", userId)
		g.Log().Errorf(ctx, "%+v", err)
		return riskserial, code
	}
	/////
	switch code {
	case mpccode.RiskCodePass, mpccode.RiskCodeNeedVerification:
		////if pass, chech tfa forbiddent
		info, err := service.TFA().TFAInfo(ctx, userId)
		if err != nil || info == nil {
			g.Log().Warning(ctx, "PerformRiskTxs:", "userId:", userId)
			g.Log().Errorf(ctx, "%+v", err)
			return "", mpccode.RiskCodeError
		}
		///
		if info.MailUpdatedAt != nil {
			befor24h := gtime.Now().Add(BeforH24)
			befor := info.MailUpdatedAt.Before(befor24h)
			g.Log().Notice(ctx, "PerformRiskTxs:", "info.MailUpdatedAt:", info.MailUpdatedAt.String(), "befor24h:", befor24h.String(), "befor:", befor)
			if !befor {
				return "", mpccode.RiskCodeForbidden
			}
		}
		///
		if info.PhoneUpdatedAt != nil {
			befor24h := gtime.Now().Add(BeforH24)
			befor := info.PhoneUpdatedAt.Before(befor24h)
			// befor := info.PhoneUpdatedAt.Before(befor24h.Time())
			g.Log().Notice(ctx, "PerformRiskTxs:", "info.PhoneUpdatedAt:", info.PhoneUpdatedAt.String(), "befor24h:", befor24h.String(), "befor:", befor)
			if !befor {
				return "", mpccode.RiskCodeForbidden
				///, nil
			}
		}
		///
		return riskserial, code
	case mpccode.RiskCodeForbidden:
		return riskserial, mpccode.RiskCodeForbidden
	case mpccode.RiskCodeError:
		return riskserial, mpccode.RiskCodeError
	case mpccode.RiskCodeNoRiskControl:
		return riskserial, mpccode.RiskCodePass
	default:
		g.Log().Error(ctx, "PerformRiskTxs:", "code:", code)
		return riskserial, mpccode.RiskCodeError
	}
}

func (s *sRisk) RiskTFA(ctx context.Context, tfaInfo *entity.Tfa, riskData *model.RiskTfa) (string, int32) {
	if !s.userControl {
		return "", mpccode.RiskCodePass
	}
	//
	riskserial := rand.GenNewSid()
	s.riskStat[riskserial] = &model.RiskStat{
		Kind: model.Kind_RiskTfa,
		Type: riskData.Type,
	}
	///
	code := mpccode.RiskCodePass
	var err error
	///
	switch riskData.Type {
	case model.Type_TfaUpdatePhone:
		code, err = s.checkTfaUpPhone(ctx, tfaInfo)
	case model.Type_TfaUpdateMail:
		code, err = s.checkTfaUpMail(ctx, tfaInfo)
	case model.Type_TfaBindPhone:
		code, err = s.checkTfaBindPhone(ctx, tfaInfo)
	case model.Type_TfaBindMail:
		code, err = s.checkTfaBindMail(ctx, tfaInfo)
	default:
		g.Log().Error(ctx, "RiskTFA:", "kind:", riskData.Type, "not support")
		return riskserial, mpccode.RiskCodeError
	}
	if err != nil {
		g.Log().Warning(ctx, "RiskTFA:", "tfaInfo:", tfaInfo, "riskDAta:", riskData)
		g.Log().Errorf(ctx, "%+v", err)
		return riskserial, mpccode.RiskCodeError
		///, err
	}
	///
	return riskserial, code
}

func (s *sRisk) GetRiskStat(ctx context.Context, riskSerial string) *model.RiskStat {
	s.riskStatLock.Lock()
	defer s.riskStatLock.Unlock()
	if r, ok := s.riskStat[riskSerial]; ok {
		return r
	}
	return nil
}

func (s *sRisk) DelRiskStat(ctx context.Context, riskSerial string) {
	s.riskStatLock.Lock()
	defer s.riskStatLock.Unlock()
	delete(s.riskStat, riskSerial)
}

// ///
var BeforH24 time.Duration

func new() *sRisk {
	///
	s := &sRisk{
		analzer:    analzyer.NewAnalzer(),
		ftruleMap:  map[string]*mpcmodel.FtRule{},
		nftruleMap: map[string]*mpcmodel.NftRule{},
		riskStat:   map[string]*model.RiskStat{},
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

	val := config.Config.UserRisk.ForbiddenTime
	// BeforH24, err = gtime.ParseDuration("-24h")
	BeforH24, err = gtime.ParseDuration(val)
	if err != nil {
		panic(err)
	}
	s.userControl = config.Config.UserRisk.UserControl
	s.txControl = config.Config.UserRisk.TxControl

	return s
}

func init() {
	///
	///
	service.RegisterRisk(new())
}
