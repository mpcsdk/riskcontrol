// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ScrapeLogsStat is the golang structure for table scrape_logs_stat.
type ScrapeLogsStat struct {
	ChainId   string      `json:"chainId"   ` //
	LastBlock int64       `json:"lastBlock" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
