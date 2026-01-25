// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWithdraw is the golang structure of table sys_withdraw for DAO operations like Where/Data.
type SysWithdraw struct {
	g.Meta        `orm:"table:sys_withdraw, do:true"`
	Id            any         //
	Code          any         //
	ManageId      any         //
	WitkeyId      any         //
	Amount        any         //
	SettledAmount any         //
	ServiceFee    any         //
	Type          any         //
	Number        any         //
	Name          any         //
	Status        any         //
	Reason        any         //
	CreateTime    *gtime.Time //
	UpdateTime    *gtime.Time //
}
