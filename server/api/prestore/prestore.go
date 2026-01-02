// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package prestore

import (
	"context"

	"server/api/prestore/v1"
)

type IPrestoreV1 interface {
	GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
}
