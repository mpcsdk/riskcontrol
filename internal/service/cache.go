// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
)

type (
	ICache interface {
		Get(ctx context.Context, key string) (*gvar.Var, error)
		Set(ctx context.Context, key string, val interface{}, duration time.Duration) error
		Remove(ctx context.Context, key string) (*gvar.Var, error)
	}
)

var (
	localCache ICache
)

func Cache() ICache {
	if localCache == nil {
		panic("implement not found for interface ICache, forgot register?")
	}
	return localCache
}

func RegisterCache(i ICache) {
	localCache = i
}
