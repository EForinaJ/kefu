package v1

import (
	dao_product "kefu-server/internal/type/product/dao"
	dto_product "kefu-server/internal/type/product/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/product/list" method:"get" tags:"商品" summary:"商品列表"`
	*dto_product.Query
}
type GetListRes struct {
	Total int                 `json:"total" dc:"总数"`
	List  []*dao_product.List `json:"list" dc:"商品列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/product/detail" method:"get" tags:"商品" summary:"商品详情"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"商品id"`
}
type GetDetailRes struct {
	*dao_product.Detail
}

type PurchaseReq struct {
	g.Meta `path:"/product/purchase" method:"post" tags:"商品" summary:"商品购买"`
	*dto_product.Purchase
}
type PurchaseRes struct {
	Code string `json:"code" dc:"订单编码"`
}
