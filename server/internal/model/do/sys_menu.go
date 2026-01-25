// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta        `orm:"table:sys_menu, do:true"`
	Id            any // 菜单ID
	PId           any // 父菜单ID
	MenuType      any // 菜单类型（1目录 2菜单）
	Name          any // 请求地址
	Path          any //
	Label         any // 菜单名称
	Component     any // 组件地址
	Icon          any //
	IsEnable      any //
	Sort          any //
	IsMenu        any //
	KeepAlive     any //
	IsHide        any // 菜单状态（2显示 1隐藏）
	IsHideTab     any //
	Link          any // 跳转
	IsIframe      any //
	ShowBadge     any //
	ShowTextBadge any //
	FixedTab      any //
	ActivePath    any //
	IsFullPage    any //
	AuthName      any //
	AuthLabel     any //
	AuthIcon      any //
	AuthSort      any //
}
