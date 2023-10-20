package tfa

import (
	"context"
	"riskcontral/internal/consts"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) exit(ctx context.Context) {

}
func (c *ControllerV1) counter(ctx context.Context, tokenId string, method string) error {
	key := tokenId + method + "counter"
	if v, err := c.cache.Get(ctx, key); err != nil || !v.IsEmpty() {
		return gerror.NewCode(consts.ErrApiLimit)
	} else {
		c.cache.Set(ctx, key, 1, apiInterval)
		return nil
	}
}
func (c *ControllerV1) limitSendVerification(ctx context.Context, tokenId string, method string) error {
	key := tokenId + method + "limitSendVerification"
	if v, err := c.cache.Get(ctx, key); err != nil || !v.IsEmpty() {
		return gerror.NewCode(consts.ErrApiLimit)
	} else {
		c.cache.Set(ctx, key, 1, limitSendInterval)
		return nil
	}
}
