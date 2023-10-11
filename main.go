package main

import (
	"riskcontral/common"
	_ "riskcontral/internal/packed"
	"riskcontral/internal/service"

	_ "riskcontral/internal/logic"
	_ "riskcontral/internal/service"

	// _ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	"riskcontral/internal/cmd"
)

func main() {
	ctx := gctx.GetInitCtx()
	workId, _ := gcfg.Instance().Get(ctx, "server.workId")
	common.InitIdGen(workId.Int())
	//

	// ///jaeger
	cfg := gcfg.Instance()
	name := cfg.MustGet(ctx, "server.name", "mpc-signer").String()
	jaegerUrl, err := cfg.Get(ctx, "jaegerUrl")
	if err != nil {
		panic(err)
	}
	tp, err := jaeger.Init(name, jaegerUrl.String())
	if err != nil {
		panic(err)
	}
	defer tp.Shutdown(ctx)
	// ///scrapelogs
	v, err := cfg.Get(ctx, "scrapeLogs", false)
	if err != nil && v.Bool() {
		g.Log().Notice(ctx, "Open ScrapeLogs")
		err = service.EthEventGeter().InitByService()
		if err != nil {
			panic(err)
		}
		service.EthEventGeter().RunBySerivce()
	} else {
		g.Log().Notice(ctx, "ShutDown ScrapeLogs")
	}
	///
	///
	cmd.Main.Run(gctx.GetInitCtx())
}
