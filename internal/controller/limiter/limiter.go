package limiter

import (
	"context"
	"riskcontral/internal/config"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

type Limiter struct {
	cache *gcache.Cache
}

var apiInterval = time.Second * 1

func (s *Limiter) ApiLimit(ctx context.Context, tokenId string, method string) error {
	key := tokenId + method + "counter"
	if v, err := s.cache.Get(ctx, key); err != nil {
		g.Log().Warning(ctx, "counter:", "tokenId:", tokenId, "method", method, "err", err)
		return mpccode.CodeApiLimit()
	} else if !v.IsEmpty() {
		g.Log().Info(ctx, "counter:", "tokenId:", tokenId, "method", method)
		return mpccode.CodeApiLimit()
	} else {
		s.cache.Set(ctx, key, 1, apiInterval)
		return nil
	}
}
func new() *Limiter {
	apiInterval = time.Duration(config.Config.Cache.ApiInterval) * time.Second
	redisCache := gcache.NewAdapterRedis(g.Redis())
	s := &Limiter{
		cache: gcache.New(),
	}
	s.cache.SetAdapter(redisCache)
	return s
}

var once sync.Once
var limiter *Limiter

func Instance() *Limiter {
	once.Do(func() {
		limiter = new()
	})
	return limiter
}
