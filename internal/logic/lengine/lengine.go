package lengine

import (
	"errors"
	"fmt"
	"riskcontral/internal/service"

	"github.com/bilibili/gengine/engine"
)

type sLEngine struct {
	RuleEnginePool map[string]*engine.GenginePool
}

func (s *sLEngine) UpRules(name, rules string) error {
	fmt.Println("uprules:", name, rules)
	err := s.newPool(name, rules)
	///
	if err != nil {
		return err
	}

	service.RulesDb().Set(name, rules)

	return nil
}

func (s *sLEngine) Exec(name string, param map[string]interface{}) (bool, error) {
	fmt.Println("exec:", name, param)
	if p, ok := s.RuleEnginePool[name]; !ok {
		return false, errors.New("no rules:" + name)
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
func (s *sLEngine) List(name string) map[string]string {
	r, err := service.RulesDb().Get(name)
	fmt.Println(err)

	return map[string]string{
		name: r,
	}
}

var apis map[string]interface{}

func (s *sLEngine) newPool(name, rules string) error {
	p, err := engine.NewGenginePool(10, 100, 2, rules, apis)
	if err != nil {
		panic(err)
	}
	s.RuleEnginePool[name] = p
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
