// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLevel is the golang structure of table sys_level for DAO operations like Where/Data.
type SysLevel struct {
	g.Meta     `orm:"table:sys_level, do:true"`
	Id         any         //
	Name       any         //
	Level      any         //
	Icon       any         //
	Experience any         //
	Benefits   any         //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}
