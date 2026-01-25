// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMedia is the golang structure of table sys_media for DAO operations like Where/Data.
type SysMedia struct {
	g.Meta     `orm:"table:sys_media, do:true"`
	Id         any         //
	Path       any         // 存放路径
	Name       any         // 文件名称
	OrName     any         // 原始文件名称
	Size       any         // 文件大小
	Ext        any         // 文件后缀
	CreateTime *gtime.Time // 创建时间
	DeleteTime *gtime.Time //
	MediaType  any         // 文件类型
}
