package packed

import (
	"riskcontrol/internal/logic/db"
	"riskcontrol/internal/logic/email"
	riskserver "riskcontrol/internal/logic/riskctrl"
	"riskcontrol/internal/logic/riskengine"
	"riskcontrol/internal/logic/sms"
	"riskcontrol/internal/logic/tfa"
	"riskcontrol/internal/logic/userInfo"
	"riskcontrol/internal/service"
)

func init() {
	service.RegisterDB(db.New())
	service.RegisterRiskEngine(riskengine.New())
	service.RegisterRiskCtrl(riskserver.New())
	service.RegisterMailCode(email.New())
	service.RegisterSmsCode(sms.New())

	////
	service.RegisterTFA(tfa.New())
	service.RegisterUserInfo(userInfo.New())
}
