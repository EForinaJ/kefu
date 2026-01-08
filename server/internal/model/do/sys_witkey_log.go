// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkeyLog is the golang structure of table sys_witkey_log for DAO operations like Where/Data.
type SysWitkeyLog struct {
	g.Meta     `orm:"table:sys_witkey_log, do:true"`
	Id         interface{} //
	ManageId   interface{} //
	WitkeyId   interface{} //
	Type       interface{} //
	Content    interface{} //
	CreateTime *gtime.Time //
}
