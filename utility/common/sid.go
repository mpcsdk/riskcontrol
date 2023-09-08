package common

import (
	"github.com/gogf/gf/container/gvar"
	"github.com/yitter/idgenerator-go/idgen"
)

func GenNewSid() string {
	var genid gvar.Var
	genid.Set(idgen.NextId())
	sid := genid.String()
	return sid
}
