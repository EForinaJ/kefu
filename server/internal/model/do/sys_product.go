// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysProduct is the golang structure of table sys_product for DAO operations like Where/Data.
type SysProduct struct {
	g.Meta              `orm:"table:sys_product, do:true"`
	Id                  any         //
	Pic                 any         //
	Code                any         //
	Type                any         // 1护航，2代肝
	CategoryId          any         //
	GameId              any         //
	Name                any         //
	Description         any         //
	Content             any         //
	WitkeyCount         any         //
	Price               any         //
	Unit                any         //
	SalesCount          any         //
	Rate                any         //
	PurchaseLimit       any         //
	SkuType             any         //
	WxMiniProgramQrCode any         //
	Status              any         // 状态：1-禁用，2-启用
	CreateTime          *gtime.Time //
	UpdateTime          *gtime.Time //
}
