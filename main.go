package main

import (
	_ "riskcontral/internal/packed"

	_ "riskcontral/internal/logic"
	_ "riskcontral/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"riskcontral/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
