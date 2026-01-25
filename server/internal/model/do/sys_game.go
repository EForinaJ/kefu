// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGame is the golang structure of table sys_game for DAO operations like Where/Data.
type SysGame struct {
	g.Meta      `orm:"table:sys_game, do:true"`
	Id          any         //
	Pid         any         // 顶级分类
	Name        any         // 分类名称
	Pic         any         // 分类背景图
	Sort        any         // 分类排序
	Description any         // 分类描述
	CreateTime  *gtime.Time // 创建日期
	UpdateTime  *gtime.Time // 更新日期
}
