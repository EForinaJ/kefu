package consts

const (
	RoleTypeKefu = 2 // 客服角色
)
const (
	ProfitTypeOrder = 1 // 订单收益
)

const (
	UserChangeBalanceTypeSystem    = 1 // 系统变更余额
	UserChangeBalanceTypeRecharge  = 2 // 用户充值
	UserChangeBalanceTypePaidOrder = 3 // 订单余额消费
)
const (
	BillTypeRecharge = 1 // 账户充值
	BillTypeRefund   = 2 // 商品退款
	BillTypeOrder    = 3 // 购买商品
)
const (
	WitkeyChangeCommissionTypeSystem     = 1 // 系统变更佣金
	WitkeyChangeCommissionTypeSettlement = 2 // 结算变更佣金
	WitkeyChangeCommissionTypeWithdraw   = 3 // 威客提现变更佣金
)
const (
	CapitalPaymentOrder       = 1 // 订单支付
	CapitalRefundOrder        = 2 // 订单退款
	CapitalPaymentRecharge    = 3 // 充值支付
	CapitalWithdrawCommission = 4 // 佣金提现
)

const (
	OrderLogTypeCreate           = 1 // 创建订单
	OrderLogTypeAddDiscount      = 2 // 添加优惠金额
	OrderLogTypeCancel           = 3 // 关闭订单
	OrderLogTypeComplete         = 4 // 完成订单
	OrderLogTypePaid             = 5 // 确认收款订单
	OrderLogTypeAfterSales       = 6 // 添加售后工单
	OrderLogTypeDistribute       = 7 // 派发威客
	OrderLogTypeDistributeCancel = 8 // 取消派单服务
	OrderLogTypeStart            = 9 // 开始服务
)

const (
	WitkeyLogTypeCreate      = 1 // 创建威客
	WitkeyLogTypeChangeTitle = 2 // 变更头衔
)

const (
	DistributeTypeTeam = 1 // 自带队伍
	DistributeTypeSelf = 2 //自己个人
)
