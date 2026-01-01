package product

import (
	"context"
	"server/internal/dao"
	dao_product "server/internal/type/product/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IProduct.
func (s *sProduct) GetDetail(ctx context.Context, id int64) (res *dao_product.Detail, err error) {
	obj, err := dao.SysProduct.Ctx(ctx).Where(dao.SysProduct.Columns().Id, id).One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if obj.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}

	var detail *dao_product.Detail
	err = gconv.Scan(obj.Map(), &detail)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	category, err := dao.SysCategory.Ctx(ctx).
		Where(dao.SysCategory.Columns().Id, obj.GMap().Get(dao.SysProduct.Columns().CategoryId)).
		Value(dao.SysCategory.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Category = category.String()

	game, err := dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Id, obj.GMap().Get(dao.SysProduct.Columns().GameId)).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	detail.Game = game.String()
	return detail, nil
}
