package workorder

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/do"
	dto_workorder "server/internal/type/workorder/dto"
	utils_code "server/internal/utils/code"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Create implements service.IWorkOrder.
func (s *sWorkOrder) Create(ctx context.Context, req *dto_workorder.Create) (err error) {
	var entity *do.SysWorkOrder
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	entity.Code = utils_code.GetCode(ctx, "WO")
	entity.CreateTime = gtime.Now()
	entity.Status = consts.WorkOrderStatusAllocate
	entity.CreatorId = ctx.Value("userId")
	_, err = dao.SysWorkOrder.Ctx(ctx).Data(&entity).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	return
}
