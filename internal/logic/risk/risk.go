package risk

import (
	"context"
	"fmt"
	"riskcontral/internal/config"
	"riskcontral/internal/model"
	"riskcontral/internal/model/entity"
	"riskcontral/internal/service"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcmodel"
	"github.com/mpcsdk/mpcCommon/mq"
	"github.com/mpcsdk/mpcCommon/rand"
)

type contractRule struct {
	ftruleMap  map[string]*mpcmodel.ContractRule
	nftruleMap map[string]*mpcmodel.ContractRule
	analzer    *analzyer.Analzyer
}

type sRisk struct {
	sceneRules map[string]*contractRule
	////
	userControl bool
	txControl   bool
	////

	///
}

func (s *sRisk) RiskTxs(ctx context.Context, userId string, signTx string) (string, int32) {
	///
	if !s.txControl {
		return "", mpccode.RiskCodePass
	}
	///
	riskserial := rand.GenNewSid()

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
		tfaInfo, err := service.DB().FetchTfaInfo(ctx, userId)
		if err != nil || tfaInfo == nil {
			g.Log().Warning(ctx, "RiskTxs:", "userId:", userId)
			g.Log().Errorf(ctx, "%+v", err)
			return "", mpccode.RiskCodeError
		}
		////if pass, chech tfa forbiddent
		// info, err := service.TFA().TFAInfo(ctx, userId)
		if tfaInfo == nil || (tfaInfo.Mail == "" && tfaInfo.Phone == "") {
			if mpccode.RiskCodePass == code {
				return riskserial, code
			} else {
				g.Log().Warning(ctx, "RiskTxs tfaNotExists:", "code:", code, "userId:", userId)
				return "", mpccode.RiskCodeError
			}
		}
		///
		if tfaInfo.MailUpdatedAt != nil {
			befor24h := gtime.Now().Add(BeforH24)
			befor := tfaInfo.MailUpdatedAt.Before(befor24h)
			g.Log().Notice(ctx, "PerformRiskTxs:", "info.MailUpdatedAt:", tfaInfo.MailUpdatedAt.String(), "befor24h:", befor24h.String(), "befor:", befor)
			if !befor {
				return "", mpccode.RiskCodeForbidden
			}
		}
		///
		if tfaInfo.PhoneUpdatedAt != nil {
			befor24h := gtime.Now().Add(BeforH24)
			befor := tfaInfo.PhoneUpdatedAt.Before(befor24h)
			// befor := info.PhoneUpdatedAt.Before(befor24h.Time())
			g.Log().Notice(ctx, "PerformRiskTxs:", "info.PhoneUpdatedAt:", tfaInfo.PhoneUpdatedAt.String(), "befor24h:", befor24h.String(), "befor:", befor)
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
	// func (s *sRisk) RiskTFA(ctx context.Context, userId string, riskData *model.RiskTfa) (string, int32) {
	if !s.userControl {
		return "", mpccode.RiskCodePass
	}
	//
	riskserial := rand.GenNewSid()

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

// ///
var BeforH24 time.Duration

func (s *sRisk) Notify(ctx context.Context, kind mq.RiskCtrlKind, notice *mq.ContractNotice) {
	if !notice.IsValid() {
		g.Log().Error(ctx, "!IsValid", notice)
		return
	}
	if kind == mq.RiskCtrlKind_ContractRule {
		//RiskCtrlKind_ContractRule
		if _, ok := s.sceneRules[notice.SceneNo]; !ok {
			abi, err := service.DB().GetContractAbi(ctx, notice.SceneNo, notice.ContractAddress)
			if err != nil {
				g.Log().Error(ctx, notice, err)
			}
			///
			s.sceneRules[notice.SceneNo] = &contractRule{
				nftruleMap: map[string]*mpcmodel.ContractRule{},
				ftruleMap:  map[string]*mpcmodel.ContractRule{},
				analzer:    analzyer.NewAnalzer(),
			}
			err = s.sceneRules[notice.SceneNo].analzer.AddAbi(abi.ContractAddress, abi.AbiContent)
			if err != nil {
				g.Log().Error(ctx, notice, err)
				return
			}
		}
		rule, err := service.DB().GetContractRule(ctx, notice.SceneNo, notice.ContractAddress)
		if err != nil {
			g.Log().Error(ctx, notice, err)
		}
		if rule.ContractKind == "ft" {
			s.sceneRules[rule.SceneNo].ftruleMap[rule.ContractAddress] = model.ContractRuleEntity2Model(rule)
		} else if rule.ContractKind == "nft" {
			s.sceneRules[rule.SceneNo].nftruleMap[rule.ContractAddress] = model.ContractRuleEntity2Model(rule)
		} else {
			g.Log().Error(ctx, notice, err)
		}
		//
	} else {
		abi, err := service.DB().GetContractAbi(ctx, notice.SceneNo, notice.ContractAddress)
		if err != nil {
			g.Log().Error(ctx, notice, err)
		}
		//RiskCtrlKind_ContractAbi
		if _, ok := s.sceneRules[notice.SceneNo]; !ok {
			s.sceneRules[notice.SceneNo] = &contractRule{
				nftruleMap: map[string]*mpcmodel.ContractRule{},
				ftruleMap:  map[string]*mpcmodel.ContractRule{},
				analzer:    analzyer.NewAnalzer(),
			}
		}
		err = s.sceneRules[notice.SceneNo].analzer.AddAbi(abi.ContractAddress, abi.AbiContent)
		if err != nil {
			g.Log().Error(ctx, notice, err)
			return
		}
	}
	//
}

// /
func (s *sRisk) getContractRules(ctx context.Context, sceneNo string) *contractRule {
	return s.sceneRules[sceneNo]
}

// /
func new() *sRisk {
	///
	ctx := gctx.GetInitCtx()
	s := &sRisk{
		// analzer:    analzyer.NewAnalzer(),
		sceneRules: map[string]*contractRule{},
		// ftruleMap:  map[string]*mpcmodel.ContractRule{},
		// nftruleMap: map[string]*mpcmodel.ContractRule{},
	}
	///
	briefs, err := service.DB().GetContractRuleBriefs(ctx, "", "")
	if err != nil {
		panic(err)
	}
	//
	for _, b := range briefs {
		rule, err := service.DB().GetContractRule(ctx, b.SceneNo, b.ContractAddress)
		if err != nil {
			panic(err)
		}

		if _, ok := s.sceneRules[rule.SceneNo]; !ok {
			s.sceneRules[rule.SceneNo] = &contractRule{
				ftruleMap:  map[string]*mpcmodel.ContractRule{},
				nftruleMap: map[string]*mpcmodel.ContractRule{},
				analzer:    analzyer.NewAnalzer(),
			}

		}
		if rule.ContractKind == "ft" {
			s.sceneRules[rule.SceneNo].ftruleMap[rule.ContractAddress] = model.ContractRuleEntity2Model(rule)
		} else if rule.ContractKind == "nft" {
			s.sceneRules[rule.SceneNo].nftruleMap[rule.ContractAddress] = model.ContractRuleEntity2Model(rule)
		} else {
			panic(rule)
		}
		//
		abi, err := service.DB().GetContractAbi(ctx, rule.SceneNo, rule.ContractAddress)
		if err != nil {
			panic(err)
		}
		if abi.AbiContent == "" {
			panic(fmt.Sprint("abi.AbiContent is empty:", rule.SceneNo, rule.ContractAddress))
		}
		err = s.sceneRules[rule.SceneNo].analzer.AddAbi(abi.ContractAddress, abi.AbiContent)
		if err != nil {
			panic(err)
		}
	}
	///
	///
	// ////

	val := config.Config.UserRisk.ForbiddenTime
	// BeforH24, err = gtime.ParseDuration("-24h")
	BeforH24, err = gtime.ParseDuration(val)
	if err != nil {
		panic(err)
	}
	s.userControl = config.Config.UserRisk.UserControl
	s.txControl = config.Config.UserRisk.TxControl

	///
	for k, rules := range s.sceneRules {
		g.Log().Info(ctx, "sceneNo:", k)
		j := gjson.New(rules.nftruleMap)
		g.Log().Info(ctx, "nftRule:", j.MustToJsonIndentString())
		j = gjson.New(rules.ftruleMap)
		g.Log().Info(ctx, "ftRule:", j.MustToJsonIndentString())
	}
	return s
}

func init() {
	///
	///
	service.RegisterRisk(new())
}
