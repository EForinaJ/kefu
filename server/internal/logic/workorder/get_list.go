package workorder

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	dao_workorder "server/internal/type/workorder/dao"
	dto_workorder "server/internal/type/workorder/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IWorkOrder.
func (s *sWorkOrder) GetList(ctx context.Context, req *dto_workorder.Query) (total int, res []*dao_workorder.List, err error) {
	m := dao.SysWorkOrder.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysWorkOrder.Columns().AssigneeId, ctx.Value("userId")).
		OrderDesc(dao.SysWorkOrder.Columns().CreateTime)

	if req.Code != "" {
		m = m.Where(dao.SysWorkOrder.Columns().Code, req.Code)
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	var list []*entity.SysWorkOrder
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	res = make([]*dao_workorder.List, len(list))
	for i, v := range list {
		var entity *dao_workorder.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
		}
		res[i] = entity
	}

	return
}
