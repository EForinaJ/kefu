// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure of table sys_config for DAO operations like Where/Data.
type SysConfig struct {
	g.Meta     `orm:"table:sys_config, do:true"`
	Id         any         // 参数主键
	Name       any         // 参数名称
	Key        any         // 参数键名
	Value      any         // 参数键值
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	Remark     any         // 备注
}
