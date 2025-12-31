package order

import (
	"context"

	v1 "kefu-server/api/order/v1"
	"kefu-server/internal/service"
)

func (c *ControllerV1) Aftersales(ctx context.Context, req *v1.AftersalesReq) (res *v1.AftersalesRes, err error) {
	err = service.Order().CheckAftersales(ctx, req.Aftersales)
	if err != nil {
		return nil, err
	}
	err = service.Order().Aftersales(ctx, req.Aftersales)
	return
}
