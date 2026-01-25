// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRecharge is the golang structure of table sys_recharge for DAO operations like Where/Data.
type SysRecharge struct {
	g.Meta     `orm:"table:sys_recharge, do:true"`
	Id         any         //
	UserId     any         //
	Code       any         //
	Amount     any         //
	PayMode    any         //
	Status     any         //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	Remark     any         //
}
