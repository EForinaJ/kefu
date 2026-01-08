// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkeyTitle is the golang structure of table sys_witkey_title for DAO operations like Where/Data.
type SysWitkeyTitle struct {
	g.Meta      `orm:"table:sys_witkey_title, do:true"`
	Id          interface{} //
	ManageId    interface{} //
	WitkeyId    interface{} //
	OldGameId   interface{} //
	OldTitleId  interface{} //
	NewGameId   interface{} //
	NewTitleId  interface{} //
	Type        interface{} //
	Description interface{} //
	Status      interface{} //
	CreateTime  *gtime.Time //
}
