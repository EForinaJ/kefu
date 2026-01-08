package dao_witkey

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	Name       string      `json:"name" dc:"威客昵称"`
	Avatar     string      `json:"avatar" dc:"威客头像"`
	Sex        int         `json:"sex" dc:"威客性别"`
	Game       string      `json:"game" dc:"游戏领域"`
	Title      string      `json:"title" dc:"头衔勋章"`
	Rate       int         `json:"rate" dc:"评分"`
	Status     int         `json:"status" dc:"接单状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"入驻时间"`
}
