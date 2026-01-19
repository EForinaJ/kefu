// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package workorder

import (
	"context"

	"server/api/workorder/v1"
)

type IWorkorderV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error)
	Solve(ctx context.Context, req *v1.SolveReq) (res *v1.SolveRes, err error)
}
