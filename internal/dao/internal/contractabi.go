// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContractabiDao is the data access object for table contractabi.
type ContractabiDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ContractabiColumns // columns contains all the column names of Table for convenient usage.
}

// ContractabiColumns defines and stores column names for table contractabi.
type ContractabiColumns struct {
	Id              string // ID
	CreateTime      string // 创建时间
	UpdateTime      string // 更新时间
	ContractName    string // 合约名
	ContractAddress string // 合约地址
	SceneNo         string // 场景号
	AbiContent      string // 合约abi
	ContractKind    string // 合约类型
}

// contractabiColumns holds the columns for table contractabi.
var contractabiColumns = ContractabiColumns{
	Id:              "id",
	CreateTime:      "createTime",
	UpdateTime:      "updateTime",
	ContractName:    "contractName",
	ContractAddress: "contractAddress",
	SceneNo:         "sceneNo",
	AbiContent:      "abiContent",
	ContractKind:    "contractKind",
}

// NewContractabiDao creates and returns a new DAO object for table data access.
func NewContractabiDao() *ContractabiDao {
	return &ContractabiDao{
		group:   "default",
		table:   "contractabi",
		columns: contractabiColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ContractabiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ContractabiDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ContractabiDao) Columns() ContractabiColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ContractabiDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ContractabiDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ContractabiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
