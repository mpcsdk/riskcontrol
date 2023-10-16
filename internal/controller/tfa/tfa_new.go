// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package tfa

import (
	"context"
	"riskcontral/internal/config"
	"riskcontral/api/tfa"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"riskcontral/internal/consts"
)

type ControllerV1 struct{
	cache *gcache.Cache
}
func (c *ControllerV1)exit(ctx context.Context) {

}
func (c *ControllerV1) counter(ctx context.Context, tokenId string, method string) error {
	key := tokenId + method
	if v, err := c.cache.Get(ctx, key); err != nil ||  !v.IsEmpty(){
		return gerror.NewCode(consts.ErrApiLimit)
	}else{
		c.cache.Set(ctx, key, 1, apiInterval)
		return nil
	}
}

var apiInterval = time.Second * 1
func init() {
	apiInterval = time.Duration(config.Config.Cache.ApiInterval) *time.Second

}
func NewV1() tfa.ITfaV1 {
	return &ControllerV1{
		cache : gcache.New(),
	}
}

