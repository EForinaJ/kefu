// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCommission is the golang structure of table sys_commission for DAO operations like Where/Data.
type SysCommission struct {
	g.Meta     `orm:"table:sys_commission, do:true"`
	Id         any         //
	WitkeyId   any         //
	Related    any         //
	After      any         //
	Before     any         //
	Amount     any         //
	Mode       any         //
	Type       any         //
	Remark     any         //
	CreateTime *gtime.Time //
}
