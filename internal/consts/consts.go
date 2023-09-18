package consts

import (
	"time"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

const (
	KEY_RiskUId    string = "riskUserId"
	KEY_RiskCode   string = "riskCode"
	KEY_RiskSerial string = "riskSerial"
)

var SessionDur time.Duration = 0
var TokenDur time.Duration = 0

func init() {
	ctx := gctx.GetInitCtx()
	SessionDur = time.Duration(gcfg.Instance().MustGet(ctx, "cache.sessionDur", 1000).Int())
	SessionDur *= time.Second
	TokenDur = time.Duration(gcfg.Instance().MustGet(ctx, "cache.tokenDur", 0).Int())
	TokenDur *= time.Second
}
