package consts

const (
	Disable = 1
	Enable  = 2
)

const (
	StatusApply   = 1 // 待审核
	StatusSuccess = 2 // 通过
	StatusFail    = 3 // 拒绝
)
const (
	PayStatusPending = 1 // 未支付
	PayStatusSuccess = 2 // 支付成功
	PayStatusRefund  = 3 // 已退款
)

const (
	OrderStatusPendingPayment = 1 // 待付款
	OrderStatusPendingOrder   = 2 // 待服务
	OrderStatusInProgress     = 3 // 进行中
	OrderStatusComplete       = 4 // 已完成
	OrderStatusCancel         = 5 // 已取消
	OrderStatusAftersales     = 6 // 已售后
)

const (
	DistributeStatusPending       = 1 // 待服务
	DistributeStatusInProgress    = 2 // 进行中
	DistributeStatusComplete      = 3 // 已完成
	DistributeStatusSettlementing = 4 // 结算中
	DistributeStatusSettlemented  = 5 // 已结算
	DistributeStatusCancel        = 6 // 已取消
)

const (
	WorkOrderStatusAllocate   = 1 // 待分配
	WorkOrderStatusProcessing = 2 // 处理中
	WorkOrderStatusResolved   = 3 // 已解决
	WorkOrderStatusCancel     = 4 // 已取消
)
