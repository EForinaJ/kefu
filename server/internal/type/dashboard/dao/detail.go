package dao_dashboard

import (
	dao_product "server/internal/type/product/dao"
	dao_workorder "server/internal/type/workorder/dao"
)

type Detail struct {
	TodaySales      float64               `json:"todaySales" dc:"今日销售额度"`
	SalesComparison float64               `json:"salesComparison" dc:"较比昨日"`
	HotProductList  []*dao_product.List   `json:"hotProductList" dc:"热门商品"`
	TodoList        []*dao_workorder.List `json:"todoList" dc:"待办工单"`
}
