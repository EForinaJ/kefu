// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package witkey

import (
	"context"

	"server/api/witkey/v1"
)

type IWitkeyV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	ChangeTitle(ctx context.Context, req *v1.ChangeTitleReq) (res *v1.ChangeTitleRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
}
