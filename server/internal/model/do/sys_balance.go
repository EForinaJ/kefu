// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysBalance is the golang structure of table sys_balance for DAO operations like Where/Data.
type SysBalance struct {
	g.Meta     `orm:"table:sys_balance, do:true"`
	Id         any         //
	UserId     any         //
	Related    any         //
	After      any         //
	Before     any         //
	Amount     any         //
	Mode       any         //
	Type       any         //
	Remark     any         //
	CreateTime *gtime.Time //
}
