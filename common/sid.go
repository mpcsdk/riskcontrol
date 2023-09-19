package common

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/yitter/idgenerator-go/idgen"
)

func GenNewSid() string {
	var genid gvar.Var
	genid.Set(idgen.NextId())
	sid := genid.String()
	return sid
}
func InitIdGen(workId int) {
	option := idgen.NewIdGeneratorOptions(uint16(workId))
	idgen.SetIdGenerator(option)
}
