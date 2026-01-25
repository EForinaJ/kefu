// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysManage is the golang structure of table sys_manage for DAO operations like Where/Data.
type SysManage struct {
	g.Meta      `orm:"table:sys_manage, do:true"`
	Id          any         // 用户ID
	Name        any         // 用户昵称
	Phone       any         // 手机号码
	Email       any         //
	Tags        any         //
	Address     any         //
	Sex         any         //
	Avatar      any         // 头像地址
	Password    any         // 密码
	Salt        any         // 密码盐
	Status      any         // 帐号状态（1停用,2正常）
	Description any         //
	LoginIp     any         //
	LoginTime   *gtime.Time //
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
	Remark      any         // 备注
	DeleteTime  *gtime.Time // 软删除
}
