package lengine

import (
	"context"
	"errors"
	"fmt"
	"riskcontral/internal/service"

	"github.com/bilibili/gengine/engine"
	"github.com/gogf/gf/v2/os/gctx"
)

type sLEngine struct {
	RuleEnginePool map[string]*engine.GenginePool
}

func (s *sLEngine) UpRules(ruleId, rules string) error {
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
func (s *sLEngine) List(ctx context.Context, ruleId string) map[string]string {
	r, err := service.RulesDb().Get(ctx, ruleId)
	fmt.Println(err)

	return map[string]string{
		"rules": r,
	}
}

var apis map[string]interface{}

func (s *sLEngine) newPool(ruleId, rules string) error {
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
	rs := service.RulesDb().AllRules(gctx.GetInitCtx())
	for name, rule := range rs {
		e.newPool(name, rule)
	}
	return e
}

func init() {
	service.RegisterLEngine(new())
}
