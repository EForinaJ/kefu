// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkeyChangeTitle is the golang structure of table sys_witkey_change_title for DAO operations like Where/Data.
type SysWitkeyChangeTitle struct {
	g.Meta      `orm:"table:sys_witkey_change_title, do:true"`
	Id          any         //
	ManageId    any         //
	WitkeyId    any         //
	AfterGame   any         //
	AfterTitle  any         //
	BeforeTitle any         //
	BeforeGame  any         //
	Status      any         //
	Reason      any         //
	Description any         //
	CreateTime  *gtime.Time //
}
