package lengine

import (
	"errors"
	"fmt"
	"riskcontral/internal/service"
	"strings"

	"github.com/bilibili/gengine/engine"
)

type sLEngine struct {
	RuleEnginePool map[string]*engine.GenginePool
}

func (s *sLEngine) UpRules(ruleId, rules string) error {
	ruleId = strings.ToLower(ruleId)
	fmt.Println("uprules:", ruleId, rules)
	err := s.newPool(ruleId, rules)
	//ruleId/
	if err != nil {
		return err
	}

	// service.RulesDb().Set(ruleId, rules)

	return nil
}

func (s *sLEngine) Exec(ruleId string, param map[string]interface{}) (bool, error) {
	ruleId = strings.ToLower(ruleId)
	fmt.Println("exec:", ruleId, param)
	if p, ok := s.RuleEnginePool[ruleId]; !ok {
		return false, errors.New("no rules:" + ruleId)
	} else {

		// param := map[string]interface{}{}
		// param["User"] = uer
		err, rst := p.ExecuteConcurrent(param)
		fmt.Println(rst)
		if err != nil {
			return false, err
		}
		for _, v := range rst {
			if v == false {
				return false, nil
			}
		}
		return true, nil
	}
}
func (s *sLEngine) List(ruleId string) map[string]string {
	ruleId = strings.ToLower(ruleId)
	r, err := service.RulesDb().Get(ruleId)
	fmt.Println(err)

	return map[string]string{
		"rules": r,
	}
}

var apis map[string]interface{}

func (s *sLEngine) newPool(ruleId, rules string) error {
	ruleId = strings.ToLower(ruleId)
	if rules == "" {
		//todo: mutex
		delete(s.RuleEnginePool, ruleId)
		return nil
	}
	p, err := engine.NewGenginePool(10, 100, 2, rules, apis)
	if err != nil {
		panic(err)
	}
	s.RuleEnginePool[ruleId] = p
	return err
}
func new() *sLEngine {
	e := &sLEngine{
		RuleEnginePool: make(map[string]*engine.GenginePool),
	}
	rs := service.RulesDb().AllRules()
	for name, rule := range rs {
		e.newPool(name, rule)
	}
	return e
}

func init() {
	service.RegisterLEngine(new())
}
