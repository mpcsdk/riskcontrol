// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ScrapeLogsStat is the golang structure of table scrape_logs_stat for DAO operations like Where/Data.
type ScrapeLogsStat struct {
	g.Meta    `orm:"table:scrape_logs_stat, do:true"`
	ChainId   interface{} //
	LastBlock interface{} //
	UpdatedAt *gtime.Time //
}
