package workorder

import (
	"context"

	v1 "server/api/workorder/v1"
	"server/internal/service"
)

func (c *ControllerV1) Solve(ctx context.Context, req *v1.SolveReq) (res *v1.SolveRes, err error) {
	err = service.WorkOrder().Solve(ctx, req.Id)
	return
}
