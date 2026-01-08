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
	g.Meta      `orm:"table:sys_witkey_onboarding, do:true"`
	Id          interface{} // 用户ID
	Name        interface{} //
	TitleId     interface{} //
	GameId      interface{} //
	Phone       interface{} //
	Address     interface{} //
	Sex         interface{} //
	Description interface{} //
	Status      interface{} //
	CreateTime  *gtime.Time // 创建时间
}
