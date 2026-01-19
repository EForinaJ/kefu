package workorder

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// Cancel implements service.IWorkOrder.
func (s *sWorkOrder) Cancel(ctx context.Context, id int64) (err error) {
	obj, err := dao.SysWorkOrder.Ctx(ctx).
		Where(dao.SysWorkOrder.Columns().AssigneeId, ctx.Value("userId")).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if obj.IsEmpty() {
		return utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}
	if gconv.Int(obj.GMap().Get(dao.SysWorkOrder.Columns().Status)) != consts.WorkOrderStatusProcessing {
		return utils_error.Err(response.FAILD, "工单已处理")
	}

	_, err = dao.SysWorkOrder.Ctx(ctx).
		Where(dao.SysWorkOrder.Columns().AssigneeId, ctx.Value("userId")).
		Data(g.Map{
			dao.SysWorkOrder.Columns().Status: consts.WorkOrderStatusCancel,
		}).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED, response.CodeMsg(response.UPDATE_FAILED))
	}
	return
}
