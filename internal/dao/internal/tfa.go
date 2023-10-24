// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TfaDao is the data access object for table tfa.
type TfaDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns TfaColumns // columns contains all the column names of Table for convenient usage.
}

// TfaColumns defines and stores column names for table tfa.
type TfaColumns struct {
	UserId         string //
	CreatedAt      string //
	DeletedAt      string //
	Phone          string //
	Mail           string //
	PhoneUpdatedAt string //
	MailUpdatedAt  string //
}

// tfaColumns holds the columns for table tfa.
var tfaColumns = TfaColumns{
	UserId:         "user_id",
	CreatedAt:      "created_at",
	DeletedAt:      "deleted_at",
	Phone:          "phone",
	Mail:           "mail",
	PhoneUpdatedAt: "phone_updated_at",
	MailUpdatedAt:  "mail_updated_at",
}

// NewTfaDao creates and returns a new DAO object for table data access.
func NewTfaDao() *TfaDao {
	return &TfaDao{
		group:   "tfa",
		table:   "tfa",
		columns: tfaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TfaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TfaDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TfaDao) Columns() TfaColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TfaDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TfaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TfaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
