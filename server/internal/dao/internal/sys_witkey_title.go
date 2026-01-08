// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysWitkeyTitleDao is the data access object for the table sys_witkey_title.
type SysWitkeyTitleDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns SysWitkeyTitleColumns // columns contains all the column names of Table for convenient usage.
}

// SysWitkeyTitleColumns defines and stores column names for the table sys_witkey_title.
type SysWitkeyTitleColumns struct {
	Id          string //
	ManageId    string //
	WitkeyId    string //
	OldGameId   string //
	OldTitleId  string //
	NewGameId   string //
	NewTitleId  string //
	Type        string //
	Description string //
	Status      string //
	CreateTime  string //
}

// sysWitkeyTitleColumns holds the columns for the table sys_witkey_title.
var sysWitkeyTitleColumns = SysWitkeyTitleColumns{
	Id:          "id",
	ManageId:    "manage_id",
	WitkeyId:    "witkey_id",
	OldGameId:   "old_game_id",
	OldTitleId:  "old_title_id",
	NewGameId:   "new_game_id",
	NewTitleId:  "new_title_id",
	Type:        "type",
	Description: "description",
	Status:      "status",
	CreateTime:  "create_time",
}

// NewSysWitkeyTitleDao creates and returns a new DAO object for table data access.
func NewSysWitkeyTitleDao() *SysWitkeyTitleDao {
	return &SysWitkeyTitleDao{
		group:   "default",
		table:   "sys_witkey_title",
		columns: sysWitkeyTitleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysWitkeyTitleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysWitkeyTitleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysWitkeyTitleDao) Columns() SysWitkeyTitleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysWitkeyTitleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysWitkeyTitleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysWitkeyTitleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
