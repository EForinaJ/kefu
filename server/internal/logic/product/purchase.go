package product

import (
	"context"
	"kefu-server/internal/consts"
	"kefu-server/internal/dao"
	dto_product "kefu-server/internal/type/product/dto"
	utils_code "kefu-server/internal/utils/code"
	utils_error "kefu-server/internal/utils/error"
	"kefu-server/internal/utils/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// Purchase implements service.IProduct.
func (s *sProduct) Purchase(ctx context.Context, req *dto_product.Purchase) (code string, err error) {
	// 判断商品是否存在
	productObj, err := dao.SysProduct.Ctx(ctx).Where(dao.SysProduct.Columns().Id, req.Id).One()
	if err != nil {
		return "", utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if productObj.IsEmpty() {
		return "", utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}
	// 判断用户是否存在
	userExist, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, req.UserId).Exist()
	if err != nil {
		return "", utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if !userExist {
		return "", utils_error.Err(response.NOT_FOUND, response.CodeMsg(response.NOT_FOUND))
	}

	// 判断是否限购
	count, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().UserId, req.UserId).
		WhereNot(dao.SysOrder.Columns().Status, consts.OrderStatusPendingPayment).
		WhereNot(dao.SysOrder.Columns().Status, consts.OrderStatusCancel).
		WhereNot(dao.SysOrder.Columns().Status, consts.OrderStatusRefund).
		Where(dao.SysOrder.Columns().ProductId, req.Id).Count()
	if err != nil {
		return "", utils_error.Err(response.DB_READ_ERROR, response.CodeMsg(response.DB_READ_ERROR))
	}
	if gconv.Int(productObj.GMap().Get(dao.SysProduct.Columns().PurchaseLimit)) != 0 && count >= gconv.Int(productObj.GMap().Get(dao.SysProduct.Columns().PurchaseLimit)) {
		return "", utils_error.Err(response.FAILD, "该商品是限购商品，您已购买过")
	}
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return "", utils_error.Err(response.DB_TX_ERROR, response.CodeMsg(response.DB_TX_ERROR))
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	orderCode := utils_code.GetCode(ctx, consts.OD)
	price := productObj.GMap().Get(dao.SysProduct.Columns().Price)
	quantity := decimal.NewFromFloat(gconv.Float64(req.Quantity))
	totalAmount := decimal.NewFromFloat(gconv.Float64(price)).Mul(quantity)
	rs, err := tx.Model(dao.SysOrder.Table()).
		Data(g.Map{
			dao.SysOrder.Columns().Code:         orderCode,
			dao.SysOrder.Columns().ProductId:    req.Id,
			dao.SysOrder.Columns().UserId:       req.UserId,
			dao.SysOrder.Columns().WitkeyCount:  productObj.GMap().Get(dao.SysProduct.Columns().WitkeyCount),
			dao.SysOrder.Columns().Quantity:     req.Quantity,
			dao.SysOrder.Columns().Unit:         productObj.GMap().Get(dao.SysProduct.Columns().Unit),
			dao.SysOrder.Columns().UnitPrice:    productObj.GMap().Get(dao.SysProduct.Columns().Price),
			dao.SysOrder.Columns().TotalAmount:  totalAmount,
			dao.SysOrder.Columns().ActualAmount: totalAmount,
			dao.SysOrder.Columns().Requirements: req.Requirements,
			dao.SysOrder.Columns().Status:       consts.OrderStatusPendingPayment,
			dao.SysOrder.Columns().PayStatus:    consts.PayStatusPending,
			dao.SysOrder.Columns().CreateTime:   gtime.Now(),
		}).
		Insert()
	if err != nil {
		return "", utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}
	rid, err := rs.LastInsertId()
	if err != nil {
		return "", utils_error.Err(response.ADD_FAILED, response.CodeMsg(response.ADD_FAILED))
	}

	_, err = tx.Model(dao.SysOrderLog.Table()).
		Data(g.Map{
			dao.SysOrderLog.Columns().OrderId:    rid,
			dao.SysOrderLog.Columns().ManageId:   ctx.Value("userId"),
			dao.SysOrderLog.Columns().CreateTime: gtime.Now(),
			dao.SysOrderLog.Columns().Content:    "创建订单",
			dao.SysOrderLog.Columns().Type:       consts.OrderLogTypeCreate,
		}).
		Insert()

	return orderCode, nil
}
