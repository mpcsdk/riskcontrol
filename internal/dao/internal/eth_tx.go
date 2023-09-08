// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EthTxDao is the data access object for table eth_tx.
type EthTxDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns EthTxColumns // columns contains all the column names of Table for convenient usage.
}

// EthTxColumns defines and stores column names for table eth_tx.
type EthTxColumns struct {
	Id         string //
	CreatedAt  string //
	UpdatedAt  string //
	DeletedAt  string //
	Address    string //
	Target     string //
	MethodId   string //
	MethodName string //
	Sig        string //
	Data       string //
	Args       string //
	From       string //
	To         string //
	Value      string //
}

// ethTxColumns holds the columns for table eth_tx.
var ethTxColumns = EthTxColumns{
	Id:         "id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	Address:    "address",
	Target:     "target",
	MethodId:   "method_id",
	MethodName: "method_name",
	Sig:        "sig",
	Data:       "data",
	Args:       "args",
	From:       "from",
	To:         "to",
	Value:      "value",
}

// NewEthTxDao creates and returns a new DAO object for table data access.
func NewEthTxDao() *EthTxDao {
	return &EthTxDao{
		group:   "default",
		table:   "eth_tx",
		columns: ethTxColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EthTxDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EthTxDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EthTxDao) Columns() EthTxColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EthTxDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EthTxDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EthTxDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
