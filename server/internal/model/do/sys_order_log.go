// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrderLog is the golang structure of table sys_order_log for DAO operations like Where/Data.
type SysOrderLog struct {
	g.Meta     `orm:"table:sys_order_log, do:true"`
	Id         any         //
	OrderId    any         //
	ManageId   any         //
	Type       any         //
	Content    any         //
	CreateTime *gtime.Time //
}
