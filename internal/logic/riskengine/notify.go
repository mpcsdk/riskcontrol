package riskengine

// // /
// func (s *sRiskEngine) NotityRiskRule(ctx context.Context, notice *mq.RiskCtrlRuleMsg) error {
// 	//
// 	rule, err := s.riskCtrlRule.GetRiskCtrlRule(ctx, notice.Data.Id, true)
// 	if err != nil {
// 		g.Log().Warning(ctx, "NotityRiskRule:", err)
// 		return nil
// 	}
// 	///
// 	//
// 	ruleEngine := &ruleEngine{
// 		Id:       notice.Data.Id,
// 		RuleName: rule.RuleName,
// 		Desc:     rule.Desc,
// 		Salience: rule.Salience,
// 		RuleStr:  rule.RuleStr,
// 		IsEnable: notice.Data.IsEnable == 1,
// 	}
// 	_, err = s.upRules(ruleEngine, notice.Opt)
// 	if err != nil {
// 		g.Log().Error(s.ctx, "NotityRiskRule:", err)
// 	}

// 	return err
// }
