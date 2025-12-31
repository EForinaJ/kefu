package service

// 定义显示接口
type IOrder interface {
	// GetList(ctx context.Context) (res *dao_Order.Detail, err error)
}

// 定义接口变量
var localOrder IOrder

// 定义一个获取接口的方法
func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

// 定义一个接口实现的注册方法
func RegisterOrder(i IOrder) {
	localOrder = i
}
