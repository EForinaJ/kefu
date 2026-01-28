package dashboard

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_dashboard "server/internal/type/dashboard/dao"
	dao_product "server/internal/type/product/dao"
	dao_workorder "server/internal/type/workorder/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IDashboard.
func (s *sDashboard) GetDetail(ctx context.Context) (res *dao_dashboard.Detail, err error) {

	detail := dao_dashboard.Detail{}

	m := dao.SysProduct.Ctx(ctx).
		Page(1, 10).
		OrderDesc(dao.SysProduct.Columns().SalesCount)

	var list []*entity.SysProduct
	err = m.Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	productList := make([]*dao_product.List, len(list))
	for i, v := range list {
		var entity *dao_product.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		game, err := dao.SysGame.Ctx(ctx).
			WherePri(v.GameId).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Game = game.String()

		category, err := dao.SysCategory.Ctx(ctx).
			WherePri(v.CategoryId).
			Value(dao.SysCategory.Columns().Name)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		entity.Category = category.String()

		productList[i] = entity
	}
	detail.HotProductList = productList

	workOrderObj := dao.SysWorkOrder.Ctx(ctx).
		Page(1, 10).
		Where(dao.SysWorkOrder.Columns().AssigneeId, ctx.Value("userId")).
		Where(dao.SysWorkOrder.Columns().Status, consts.WorkOrderStatusProcessing).
		OrderDesc(dao.SysWorkOrder.Columns().CreateTime)
	var workOrderList []*entity.SysWorkOrder
	err = workOrderObj.Scan(&workOrderList)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	wkList := make([]*dao_workorder.List, len(workOrderList))
	for i, v := range workOrderList {
		var wkEntity *dao_workorder.List
		err = gconv.Scan(v, &wkEntity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}

		wkList[i] = wkEntity
	}
	detail.TodoList = wkList

	return &detail, nil
}
