package workorder

import (
	"context"

	v1 "server/api/workorder/v1"
	"server/internal/service"
)

func (c *ControllerV1) Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error) {
	err = service.WorkOrder().Cancel(ctx, req.Id)
	return
}
