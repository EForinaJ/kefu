// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWorkOrder is the golang structure of table sys_work_order for DAO operations like Where/Data.
type SysWorkOrder struct {
	g.Meta     `orm:"table:sys_work_order, do:true"`
	Id         any         //
	CreatorId  any         //
	AssigneeId any         //
	Code       any         //
	Priority   any         //
	Title      any         //
	Content    any         //
	Status     any         //
	CreateTime *gtime.Time //
}
