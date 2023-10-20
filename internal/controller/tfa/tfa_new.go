// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package tfa

import (
	"riskcontral/internal/config"
	"riskcontral/api/tfa"
	"time"

	"github.com/gogf/gf/v2/os/gcache"
)

type ControllerV1 struct{
	cache *gcache.Cache
}


var apiInterval = time.Second * 1
var limitSendInterval = time.Second * 60
func init() {
	apiInterval = time.Duration(config.Config.Cache.ApiInterval) *time.Second
	limitSendInterval  = time.Duration(config.Config.UserRisk.VerificationCodeInterval) *time.Second	
}
func NewV1() tfa.ITfaV1 {
	return &ControllerV1{
		cache : gcache.New(),
	}
}

