// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AggFt24HDao is the data access object for table agg_ft_24h.
type AggFt24HDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AggFt24HColumns // columns contains all the column names of Table for convenient usage.
}

// AggFt24HColumns defines and stores column names for table agg_ft_24h.
type AggFt24HColumns struct {
	From       string //
	To         string //
	Value      string //
	Contract   string //
	UpdatedAt  string //
	MethodName string //
	FromBlock  string //
	ToBlock    string //
	MethodSig  string //
	FtName     string //
}

// aggFt24HColumns holds the columns for table agg_ft_24h.
var aggFt24HColumns = AggFt24HColumns{
	From:       "from",
	To:         "to",
	Value:      "value",
	Contract:   "contract",
	UpdatedAt:  "updated_at",
	MethodName: "method_name",
	FromBlock:  "fromBlock",
	ToBlock:    "toBlock",
	MethodSig:  "method_sig",
	FtName:     "ft_name",
}

// NewAggFt24HDao creates and returns a new DAO object for table data access.
func NewAggFt24HDao() *AggFt24HDao {
	return &AggFt24HDao{
		group:   "scrapeLogs",
		table:   "agg_ft_24h",
		columns: aggFt24HColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AggFt24HDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AggFt24HDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AggFt24HDao) Columns() AggFt24HColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AggFt24HDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AggFt24HDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AggFt24HDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
