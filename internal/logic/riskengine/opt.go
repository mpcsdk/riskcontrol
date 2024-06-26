package riskengine

import (
	"fmt"
	"sort"

	"github.com/bilibili/gengine/engine"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func tidyRuleStr(ruleName, desc, ruleStr string, salience int) string {
	return fmt.Sprintf(`rule "%s" "%s" salience %d
	begin
		%s
	end
	`, ruleName, desc, salience, ruleStr)
}

func (s *sRiskEngine) VerifyRules(ruleName, ruleStr string) error {
	ruleStr = tidyRuleStr(ruleName, "check", ruleStr, 100)
	_, err := s.newPool(ruleName, ruleStr)
	return err
}

func (s *sRiskEngine) UpRules(ruleName, ruleDesc, ruleStr string, salience int) (string, error) {
	return s.upRules(&ruleEngine{
		Id:       0,
		RuleName: ruleName,
		Desc:     ruleDesc,
		Salience: salience,
		RuleStr:  ruleStr,
		IsEnable: true,
	}, "add")
}
func (s *sRiskEngine) upRules(rule *ruleEngine, opt string) (string, error) {
	///
	tidyRuleStr := tidyRuleStr(rule.RuleName, rule.Desc, rule.RuleStr, rule.Salience)
	rule.tidyRuleStr = tidyRuleStr

	//
	if opt == "add" || opt == "update" {
		//
		p, err := s.newPool(rule.RuleName, tidyRuleStr)
		if err != nil {
			return rule.EngineName(), err
		}
		rule.Pool = p
		//
		if opt == "add" {
			s.TxEnginePool = append(s.TxEnginePool, rule)
		} else {
			for i, r := range s.TxEnginePool {
				if r.Id == rule.Id {
					s.TxEnginePool[i] = rule
				}
			}
		}
		sort.Slice(s.TxEnginePool, func(i, j int) bool {
			if s.TxEnginePool[i].Salience == s.TxEnginePool[j].Salience {
				return s.TxEnginePool[i].RuleName < s.TxEnginePool[j].RuleName
			}
			return s.TxEnginePool[i].Salience < s.TxEnginePool[j].Salience
		})
	} else if opt == "delete" {
		for i, r := range s.TxEnginePool {
			if r.Id == rule.Id {
				s.TxEnginePool = append(s.TxEnginePool[:i], s.TxEnginePool[i+1:]...)
			}
		}
	}
	//

	return rule.EngineName(), nil
}

// /
func (s *sRiskEngine) newPool(ruleName, rules string) (*engine.GenginePool, error) {
	if rules == "" {
		return nil, mpccode.CodeParamInvalid()
	}
	p, err := engine.NewGenginePool(10, 100, 2, rules, apis)
	if err != nil {
		return nil, gerror.Wrap(mpccode.CodeParamInvalid(),
			mpccode.ErrDetails(
				mpccode.ErrDetail("engineErr:", err.Error()),
			))
	}
	return p, nil
}
