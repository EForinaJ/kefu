export enum ProductStatus {
    Disable = 1, // 下架
    Enable = 2, // 上架
}
export enum PayStatus {
    PendingPayment = 1, // 待付款 
    Paid = 2, // 已付款
    Refunded = 3, // 已退款
}

export enum OrderStatus {
    PendingPayment = 1, // 待付款 
    PendingService = 2, // 待服务
    InProgress = 3, // 进行中
    Completed = 4, // 已完成
    Cancel = 5, // 已取消
    Refund = 6, // 已退款
}
