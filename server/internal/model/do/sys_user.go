// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta      `orm:"table:sys_user, do:true"`
	Id          any         // 用户ID
	LevelId     any         //
	WxUnionId   any         //
	WxOpenId    any         //
	Name        any         // 用户昵称
	Phone       any         // 手机号码
	Sex         any         // 用户性别（1男 2女 3未知）
	Address     any         //
	Birthday    *gtime.Time //
	Avatar      any         // 头像地址
	Password    any         // 密码
	Salt        any         // 密码盐
	Cover       any         //
	Experience  any         //
	Balance     any         // 余额
	Description any         //
	Status      any         // 帐号状态（1停用,2正常）
	LoginIp     any         //
	LoginTime   *gtime.Time //
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
	DeleteTime  *gtime.Time // 软删除
	Remark      any         // 备注
}
