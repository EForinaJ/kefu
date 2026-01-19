package service

import (
	"context"
	dao_workorder "server/internal/type/workorder/dao"
	dto_workorder "server/internal/type/workorder/dto"
)

// 定义显示接口
type IWorkOrder interface {
	GetList(ctx context.Context, req *dto_workorder.Query) (total int, res []*dao_workorder.List, err error)
	Create(ctx context.Context, req *dto_workorder.Create) (err error)

	Solve(ctx context.Context, id int64) (err error)
	Cancel(ctx context.Context, id int64) (err error)
}

// 定义接口变量
var localWorkOrder IWorkOrder

// 定义一个获取接口的方法
func WorkOrder() IWorkOrder {
	if localWorkOrder == nil {
		panic("implement not found for interface IWorkOrder, forgot register?")
	}
	return localWorkOrder
}

// 定义一个接口实现的注册方法
func RegisterWorkOrder(i IWorkOrder) {
	localWorkOrder = i
}
