// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysWorkOrderDao is the data access object for the table sys_work_order.
type SysWorkOrderDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SysWorkOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SysWorkOrderColumns defines and stores column names for the table sys_work_order.
type SysWorkOrderColumns struct {
	Id         string //
	CreatorId  string //
	AssigneeId string //
	Code       string //
	Priority   string //
	Title      string //
	Content    string //
	Status     string //
	CreateTime string //
}

// sysWorkOrderColumns holds the columns for the table sys_work_order.
var sysWorkOrderColumns = SysWorkOrderColumns{
	Id:         "id",
	CreatorId:  "creator_id",
	AssigneeId: "assignee_id",
	Code:       "code",
	Priority:   "priority",
	Title:      "title",
	Content:    "content",
	Status:     "status",
	CreateTime: "create_time",
}

// NewSysWorkOrderDao creates and returns a new DAO object for table data access.
func NewSysWorkOrderDao(handlers ...gdb.ModelHandler) *SysWorkOrderDao {
	return &SysWorkOrderDao{
		group:    "default",
		table:    "sys_work_order",
		columns:  sysWorkOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysWorkOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysWorkOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysWorkOrderDao) Columns() SysWorkOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysWorkOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysWorkOrderDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysWorkOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
