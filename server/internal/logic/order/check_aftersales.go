package order

import (
	"context"

	"server/internal/consts"
	"server/internal/dao"
	dto_order "server/internal/type/order/dto"
	utils_error "server/internal/utils/error"
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// CheckAftersales implements service.IOrder.
func (s *sOrder) CheckAftersales(ctx context.Context, req *dto_order.Aftersales) (err error) {
	orderActualAmount, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, req.Id).
		Value(dao.SysOrder.Columns().ActualAmount)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}

	if !decimal.NewFromFloat(req.Amount).LessThanOrEqual(decimal.NewFromFloat(orderActualAmount.Float64())) {
		return utils_error.Err(response.FAILD, "退款金额超过实付金额")
	}

	exist, err := dao.SysAftersales.Ctx(ctx).
		Where(dao.SysAftersales.Columns().OrderId, req.Id).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if exist {
		return utils_error.Err(response.FAILD, "该订单已经提交售后工单，请等待处理")
	}
	status, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, req.Id).
		Value(dao.SysOrder.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if gconv.Int(status) == consts.OrderStatusPendingPayment {
		return utils_error.Err(response.FAILD, "订单未支付")
	}

	if gconv.Int(status) == consts.OrderStatusCancel {
		return utils_error.Err(response.FAILD, "订单已取消")
	}

	if gconv.Int(status) == consts.OrderStatusRefund {
		return utils_error.Err(response.FAILD, "订单已退款")
	}

	return nil
}
