// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserExperience is the golang structure of table sys_user_experience for DAO operations like Where/Data.
type SysUserExperience struct {
	g.Meta     `orm:"table:sys_user_experience, do:true"`
	Id         any         //
	UserId     any         //
	Experience any         //
	Type       any         //
	CreateTime *gtime.Time //
}
