// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkey is the golang structure of table sys_witkey for DAO operations like Where/Data.
type SysWitkey struct {
	g.Meta      `orm:"table:sys_witkey, do:true"`
	Id          any         // 用户ID
	Name        any         //
	TitleId     any         //
	GameId      any         //
	Phone       any         //
	Avatar      any         //
	Address     any         //
	Salt        any         //
	Password    any         //
	Album       any         //
	Sex         any         //
	Birthday    *gtime.Time //
	Rate        any         //
	Commission  any         //
	Description any         //
	Status      any         //
	LoginIp     any         //
	LoginTime   *gtime.Time //
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
