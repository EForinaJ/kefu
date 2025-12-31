package dao_order

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id          int64       `json:"id" dc:"ID"`
	Code        string      `json:"code" dc:"订单号"`
	User        string      `json:"user" dc:"下单用户"`
	Product     *Product    `json:"product" dc:"订单商品"`
	TotalAmount float64     `json:"totalAmount" dc:"订单总额"`
	Status      int         `json:"status" dc:"订单状态"`
	CreateTime  *gtime.Time `json:"createTime" dc:"下单时间"`
}

type DistributeList struct {
	Id         int64       `json:"id" dc:"项目Id"`
	Witkey     string      `json:"witkey" dc:"接单威客"`
	Game       string      `json:"game" dc:"游戏领域"`
	Title      string      `json:"title" dc:"头衔勋章"`
	IsCancel   int         `json:"isCancel" dc:"是否取消"`
	CreateTime *gtime.Time `json:"createTime" dc:"下单时间"`
}
