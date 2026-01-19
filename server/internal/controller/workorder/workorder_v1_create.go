package workorder

import (
	"context"

	v1 "server/api/workorder/v1"
	"server/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.WorkOrder().Create(ctx, req.Create)
	return
}
