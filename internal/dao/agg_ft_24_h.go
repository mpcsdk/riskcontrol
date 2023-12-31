// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"riskcontral/internal/dao/internal"
)

// internalAggFt24HDao is internal type for wrapping internal DAO implements.
type internalAggFt24HDao = *internal.AggFt24HDao

// aggFt24HDao is the data access object for table agg_ft_24h.
// You can define custom methods on it to extend its functionality as you wish.
type aggFt24HDao struct {
	internalAggFt24HDao
}

var (
	// AggFt24H is globally public accessible object for table agg_ft_24h operations.
	AggFt24H = aggFt24HDao{
		internal.NewAggFt24HDao(),
	}
)

// Fill with you ideas below.
