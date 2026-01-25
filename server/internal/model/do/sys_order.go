// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrder is the golang structure of table sys_order for DAO operations like Where/Data.
type SysOrder struct {
	g.Meta         `orm:"table:sys_order, do:true"`
	Id             any         //
	Code           any         //
	UserId         any         //
	ProductId      any         //
	WitkeyCount    any         //
	Quantity       any         //
	Unit           any         //
	UnitPrice      any         //
	TotalAmount    any         // 订单总额
	DiscountAmount any         // 优惠金额
	ActualAmount   any         // 实付金额
	PayAmount      any         //
	Requirements   any         //
	Status         any         //
	PayCode        any         //
	PayMode        any         //
	PayStatus      any         //
	PayTime        *gtime.Time //
	StartTime      *gtime.Time //
	FinishTime     *gtime.Time //
	CreateTime     *gtime.Time //
	UpdateTime     *gtime.Time //
	Remark         any         //
}
