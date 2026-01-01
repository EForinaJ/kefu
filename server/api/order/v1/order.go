package v1

import (
	dao_order "server/internal/type/order/dao"
	dto_order "server/internal/type/order/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"订单" summary:"订单列表"`
	*dto_order.Query
}
type GetListRes struct {
	Total int               `json:"total" dc:"总数"`
	List  []*dao_order.List `json:"list" dc:"订单列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/order/detail" method:"get" tags:"订单" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_order.Detail
}

type AftersalesReq struct {
	g.Meta `path:"/order/aftersales" method:"post" tags:"订单" summary:"添加售后工单"`
	*dto_order.Aftersales
}
type AftersalesRes struct{}

type AddDiscountReq struct {
	g.Meta `path:"/order/add/discount" method:"post" tags:"订单" summary:"订单添加折扣"`
	*dto_order.AddDiscount
}
type AddDiscountRes struct{}

type PaidReq struct {
	g.Meta `path:"/order/paid" method:"post" tags:"订单" summary:"订单确认收款"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type PaidRes struct{}

type CancelReq struct {
	g.Meta `path:"/order/cancel" method:"post" tags:"订单" summary:"关闭订单"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type CancelRes struct{}

type StartServiceReq struct {
	g.Meta `path:"/order/start" method:"post" tags:"订单" summary:"开始服务"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type StartServiceRes struct{}

type CompleteReq struct {
	g.Meta `path:"/order/complete" method:"post" tags:"订单" summary:"完成服务"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type CompleteRes struct{}

type DistributeReq struct {
	g.Meta `path:"/order/distribute" method:"post" tags:"订单" summary:"派发威客"`
	*dto_order.Distribute
}
type DistributeRes struct{}

type GetDistributeListReq struct {
	g.Meta `path:"/order/distribute/list" method:"get" tags:"订单" summary:"派单列表"`
	*dto_order.DistributeQuery
}
type GetDistributeListRes struct {
	Total int                         `json:"total" dc:"总数"`
	List  []*dao_order.DistributeList `json:"list" dc:"派单列表"`
}

type DistributeCancelReq struct {
	g.Meta `path:"/order/distribute/cancel" method:"post" tags:"订单" summary:"派发取消"`
	*dto_order.DistributeCancel
}
type DistributeCancelRes struct{}
