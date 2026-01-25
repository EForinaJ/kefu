// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkeyOnboarding is the golang structure of table sys_witkey_onboarding for DAO operations like Where/Data.
type SysWitkeyOnboarding struct {
	g.Meta     `orm:"table:sys_witkey_onboarding, do:true"`
	Id         any         // 用户ID
	ManageId   any         //
	Name       any         //
	TitleId    any         //
	GameId     any         //
	Phone      any         //
	Salt       any         //
	Password   any         //
	Birthday   *gtime.Time //
	Address    any         //
	Sex        any         //
	Status     any         //
	CreateTime *gtime.Time // 创建时间
}
