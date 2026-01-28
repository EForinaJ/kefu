export enum ApplyStatus {
    Pending = 1, // 待审核
    Success = 2, // 已通过
    Fail = 3, // 拒绝
}

export enum Status {
    Disable = 1, // 禁用
    Enable = 2, // 启用
}


export enum ProductStatus {
    Disable = 1, // 下架
    Enable = 2, // 上架
}
export enum PayStatus {
    Pending = 1, // 待付款 
    Success = 2, // 支付成功
    Fail = 3, // 支付失败
}

export enum OrderStatus {
    PendingPayment = 1, // 待付款 
    PendingService = 2, // 待服务
    InProgress = 3, // 进行中
    Completed = 4, // 已完成
    Cancel = 5, // 已取消
    AfterSales = 6, // 已售后
}
export enum DistributeStatus {
    Pending = 1, // 待服务
    InProgress = 2, // 待服务
    Completed = 3, // 已完成
    Settlementing = 4, // 结算中
    Settlemented = 5, // 已结算
    Cancel = 6, // 已取消
}

export enum WorkOrderStatus {
    Allocate = 1, // 待分配
    Processing = 2, // 处理中
    Resolved = 3, // 已解决
    Cancel = 4, // 已取消
}
