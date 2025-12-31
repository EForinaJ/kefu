package product

import (
	"context"

	v1 "kefu-server/api/product/v1"
	"kefu-server/internal/service"
)

func (c *ControllerV1) Purchase(ctx context.Context, req *v1.PurchaseReq) (res *v1.PurchaseRes, err error) {
	code, err := service.Product().Purchase(ctx, req.Purchase)
	if err != nil {
		return nil, err
	}
	res = &v1.PurchaseRes{
		Code: code,
	}
	return
}
