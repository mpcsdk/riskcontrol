package main

import (
	"riskcontral/internal/conf"
	_ "riskcontral/internal/packed"

	_ "riskcontral/internal/logic"
	_ "riskcontral/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"

	"riskcontral/internal/cmd"

	_ "riskcontral/internal/conf"

	"github.com/mpcsdk/mpcCommon/rand"
)

func main() {
	///
	g.Log().SetAsync(true)

	///
	ctx := gctx.GetInitCtx()
	workId := conf.Config.Server.WorkId
	rand.InitIdGen(workId)
	//
	gtime.SetTimeZone("Asia/Shanghai")
	// ///jaeger
	name := conf.Config.Server.Name
	jaegerUrl := conf.Config.JaegerUrl
	tp, err := jaeger.Init(name, jaegerUrl)
	if err != nil {
		panic(err)
	}
	defer tp.Shutdown(ctx)

	///
	///
	cmd.Main.Run(gctx.GetInitCtx())
}
