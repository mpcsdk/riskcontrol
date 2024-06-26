package riskengine

import (
	"context"
	"riskcontrol/internal/logic/riskengine/intrinsic"
	"riskcontrol/internal/service"

	"github.com/bilibili/gengine/engine"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type ruleEngine struct {
	Id       int
	RuleName string
	Desc     string
	Salience int
	// ChainId  string
	IsEnable bool
	RuleStr  string
	///
	Pool *engine.GenginePool
	///
	tidyRuleStr string
	engineName  string
}

func engineName(ruleName, ChainId string) string {
	return ruleName + "_" + ChainId
}
func (s *ruleEngine) EngineName() string {
	if s.engineName == "" {
		s.engineName = engineName(s.RuleName, s.Desc)
	}
	return s.engineName
}

type sRiskEngine struct {
	ctx           context.Context
	TxEnginePool  []*ruleEngine
	TfaEnginePool []*ruleEngine
	//////
	// riskCtrlRule *mpcdao.RiskCtrlRule
}

// /
func (s *sRiskEngine) AddApi(key string, api interface{}) {
	apis[key] = api
}

var apis = map[string]interface{}{}

// /
func New() *sRiskEngine {
	///
	r := g.Redis("aggRiskCtrl")
	_, err := r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	///
	s := &sRiskEngine{
		ctx:           gctx.GetInitCtx(),
		TxEnginePool:  []*ruleEngine{},
		TfaEnginePool: []*ruleEngine{},
		////
		// riskCtrlRule: mpcdao.NewRiskCtrlRule(r, conf.Config.Cache.Duration),
	}
	//apis
	apis = intrinsic.IntrinsicApis(s.ctx)

	////
	////
	briefs, err := service.DB().RiskCtrl().GetRiskCtrlRuleBriefs(s.ctx)
	if err != nil {
		panic(err)
	}
	for _, brief := range briefs {
		rule, err := service.DB().RiskCtrl().GetRiskCtrlRule(s.ctx, brief.Id, true)
		if err != nil {
			panic(err)
		}
		//
		ruleEngine := &ruleEngine{
			Id:       brief.Id,
			RuleName: rule.RuleName,
			Desc:     rule.Desc,
			Salience: rule.Salience,
			RuleStr:  rule.RuleStr,
			IsEnable: rule.IsEnable == 1,
		}
		_, err = s.upRules(ruleEngine, "add")
		if err != nil {
			g.Log().Error(s.ctx, "NotityRiskRule:", err)
		}

	}
	return s
}
