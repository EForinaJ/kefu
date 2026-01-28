
export enum AfterSalesType {
  Refund = 1, // 仅退款
  Compensate = 2, // 仅补偿
  RefundCompensate = 3, // 退款加补偿
}

export enum SexType {
  Male = 1,      // 男
  Female = 2,    // 女
  Other = 3,     // 其他
}
export enum DistributeType {
  Team = 1,      // 自带队伍
  Self = 2,    // 个人接单
}

export enum ChangeTitleType {
  Downgrade = 1,      // 降级
  Upgrade = 2,    // 升级
  ChangeGame = 3,    // 换游戏
}

export enum PriorityType {
  urgent = 1, // 紧急
  secondary = 2, // 高
  tertiary = 3, // 中
  quaternary = 4, // 低
}
