package dao_workorder

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	Code       string      `json:"code" dc:"工单号"`
	Title      string      `json:"title" dc:"工单标题"`
	Priority   int         `json:"priority" dc:"紧急状态"`
	Content    string      `json:"content" dc:"上报内容"`
	Status     int         `json:"status" dc:"状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"入驻时间"`
}
