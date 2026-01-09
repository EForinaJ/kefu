package dashboard

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_dashboard "server/internal/type/dashboard/dao"
	dao_product "server/internal/type/product/dao"
	utils_common "server/internal/utils/common"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IDashboard.
func (s *sDashboard) GetDetail(ctx context.Context) (res *dao_dashboard.Detail, err error) {
	// 获取昨天收益
	yesterday, err := dao.SysCapital.Ctx(ctx).
		Where(dao.SysCapital.Columns().Type, consts.CapitalPaymentOrder).
		Where("DATE(create_time) = ?", gtime.Now().AddDate(0, 0, -1).Format("Y-m-d")).
		Sum("amount")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	today, err := dao.SysCapital.Ctx(ctx).
		Where(dao.SysCapital.Columns().Type, consts.CapitalPaymentOrder).
		Where("DATE(create_time) = ?", gtime.Now().Format("Y-m-d")).
		Sum("amount")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	calculate := utils_common.CalculateChangePercent(today, yesterday)

	detail := dao_dashboard.Detail{
		TodaySales:      today,
		SalesComparison: calculate,
	}

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

	return &detail, nil
}
