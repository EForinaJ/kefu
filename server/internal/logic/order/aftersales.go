package order

import (
	"context"
	"server/internal/consts"
	"server/internal/dao"
	dto_order "server/internal/type/order/dto"
	utils_code "server/internal/utils/code"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Aftersales implements service.IOrder.
func (s *sOrder) Aftersales(ctx context.Context, req *dto_order.Aftersales) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_TX_ERROR, response.CodeMsg(response.DB_TX_ERROR))
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	orderUserId, err := tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, req.Id).
		Value(dao.SysOrder.Columns().UserId)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	_, err = tx.Model(dao.SysAftersales.Table()).
		Data(g.Map{
			dao.SysAftersales.Columns().OrderId:    req.Id,
			dao.SysAftersales.Columns().UserId:     orderUserId.Int64(),
			dao.SysAftersales.Columns().Code:       utils_code.GetCode(ctx, consts.AS),
			dao.SysAftersales.Columns().Amount:     req.Amount,
			dao.SysAftersales.Columns().Type:       req.Type,
			dao.SysAftersales.Columns().ManageId:   ctx.Value("userId"),
			dao.SysAftersales.Columns().Reason:     req.Reason,
			dao.SysAftersales.Columns().Status:     consts.StatusApply,
			dao.SysAftersales.Columns().CreateTime: gtime.Now(),
		}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	//  添加订单日志
	_, err = tx.Model(dao.SysOrderLog.Table()).Data(g.Map{
		dao.SysOrderLog.Columns().CreateTime: gtime.Now(),
		dao.SysOrderLog.Columns().Content:    "售后工单",
		dao.SysOrderLog.Columns().ManageId:   ctx.Value("userId"),
		dao.SysOrderLog.Columns().OrderId:    req.Id,
		dao.SysOrderLog.Columns().Type:       consts.OrderLogTypeAfterSales,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	return
}
