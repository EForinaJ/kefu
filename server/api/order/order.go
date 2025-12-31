// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package order

import (
	"context"

	"kefu-server/api/order/v1"
)

type IOrderV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	Aftersales(ctx context.Context, req *v1.AftersalesReq) (res *v1.AftersalesRes, err error)
	AddDiscount(ctx context.Context, req *v1.AddDiscountReq) (res *v1.AddDiscountRes, err error)
	Paid(ctx context.Context, req *v1.PaidReq) (res *v1.PaidRes, err error)
	Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error)
	StartService(ctx context.Context, req *v1.StartServiceReq) (res *v1.StartServiceRes, err error)
	Complete(ctx context.Context, req *v1.CompleteReq) (res *v1.CompleteRes, err error)
	Distribute(ctx context.Context, req *v1.DistributeReq) (res *v1.DistributeRes, err error)
	GetDistributeList(ctx context.Context, req *v1.GetDistributeListReq) (res *v1.GetDistributeListRes, err error)
	DistributeCancel(ctx context.Context, req *v1.DistributeCancelReq) (res *v1.DistributeCancelRes, err error)
}
