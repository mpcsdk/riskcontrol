// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContractruleDao is the data access object for table contractrule.
type ContractruleDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ContractruleColumns // columns contains all the column names of Table for convenient usage.
}

// ContractruleColumns defines and stores column names for table contractrule.
type ContractruleColumns struct {
	Id               string // ID
	CreateTime       string // 创建时间
	UpdateTime       string // 更新时间
	SceneNo          string // 场景号
	ContractAddress  string // 合约地址
	ContractName     string // 合约名
	MethodName       string // 方法名
	MethodSignature  string // 方法签名
	MethodFromField  string // 方法from字段名
	MethodToField    string // 方法to字段名
	MethodValueField string // 方法value字段名
	EventName        string // 事件名
	EventSignature   string // 事件签名
	EventFromField   string // 事件from字段名
	EventToField     string // 事件to字段名
	EventValueField  string // 事件value字段名
	ControlOperation string // 控制操作
	Threshold        string // 阈值
	ContractKind     string // 合约类型
	EventTopic       string // 事件topic
	WhiteAddrList    string // to地址白名单
	ChainId          string // 链id
}

// contractruleColumns holds the columns for table contractrule.
var contractruleColumns = ContractruleColumns{
	Id:               "id",
	CreateTime:       "createTime",
	UpdateTime:       "updateTime",
	SceneNo:          "sceneNo",
	ContractAddress:  "contractAddress",
	ContractName:     "contractName",
	MethodName:       "methodName",
	MethodSignature:  "methodSignature",
	MethodFromField:  "methodFromField",
	MethodToField:    "methodToField",
	MethodValueField: "methodValueField",
	EventName:        "eventName",
	EventSignature:   "eventSignature",
	EventFromField:   "eventFromField",
	EventToField:     "eventToField",
	EventValueField:  "eventValueField",
	ControlOperation: "controlOperation",
	Threshold:        "threshold",
	ContractKind:     "contractKind",
	EventTopic:       "eventTopic",
	WhiteAddrList:    "whiteAddrList",
	ChainId:          "chainId",
}

// NewContractruleDao creates and returns a new DAO object for table data access.
func NewContractruleDao() *ContractruleDao {
	return &ContractruleDao{
		group:   "default",
		table:   "contractrule",
		columns: contractruleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ContractruleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ContractruleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ContractruleDao) Columns() ContractruleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ContractruleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ContractruleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ContractruleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
