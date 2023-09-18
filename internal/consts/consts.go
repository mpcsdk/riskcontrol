package consts

import (
	"time"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

const (
	STATE_None int = iota
	STATE_Auth
	STATE_HandShake
	STATE_Done
	STATE_Err
)
const (
	KEY_context     string = "context2"
	KEY_privatekey2 string = "privatekey2"
	KEY_hashproof   string = "hashproof"
	KEY_zkproof1    string = "zkproof1"
	KEY_zkproof2    string = "zkproof2"
	KEY_publickey2  string = "public_key_v2"
	KEY_request     string = "request"
	KEY_msg         string = "msg"
	KEY_txs         string = "txs"
	KEY_signature   string = "signature"
	///
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
