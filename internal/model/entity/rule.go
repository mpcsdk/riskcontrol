// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Rule is the golang structure for table rule.
type Rule struct {
	Id        int64       `json:"id"        ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
	RuleId    string      `json:"ruleId"    ` //
	RuleDesc  string      `json:"ruleDesc"  ` //
	Rules     string      `json:"rules"     ` //
}
