// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPunish is the golang structure of table sys_punish for DAO operations like Where/Data.
type SysPunish struct {
	g.Meta     `orm:"table:sys_punish, do:true"`
	Id         any         //
	Code       any         //
	WitkeyId   any         //
	ManageId   any         //
	Money      any         //
	Content    any         //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}
