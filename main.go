package main

import (
	"riskcontral/common"
	_ "riskcontral/internal/packed"

	_ "riskcontral/internal/logic"
	_ "riskcontral/internal/service"

	// _ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	"riskcontral/internal/cmd"
)

func main() {
	workId, _ := gcfg.Instance().Get(gctx.GetInitCtx(), "server.workId")
	common.InitIdGen(workId.Int())
	cmd.Main.Run(gctx.GetInitCtx())
}
