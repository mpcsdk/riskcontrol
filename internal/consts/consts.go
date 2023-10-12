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
	///
	KEY_TFAKindUpMail  string = "upMail"
	KEY_TFAKindUpPhone string = "upPhone"
	KEY_TFAKindCreate  string = "createTFA"
	KEY_TFAInfoCache   string = "tfaInfoCache"
)

var SessionDur time.Duration = 0

func init() {
	ctx := gctx.GetInitCtx()
	SessionDur = time.Duration(gcfg.Instance().MustGet(ctx, "cache.sessionDuration", 1000).Int())
	SessionDur *= time.Second

}
