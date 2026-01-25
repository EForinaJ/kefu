// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysComment is the golang structure of table sys_comment for DAO operations like Where/Data.
type SysComment struct {
	g.Meta     `orm:"table:sys_comment, do:true"`
	Id         any         //
	ManageId   any         //
	UserId     any         //
	ProductId  any         //
	Content    any         //
	Rate       any         //
	Images     any         //
	IsTop      any         //
	Status     any         //
	Ip         any         //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}
