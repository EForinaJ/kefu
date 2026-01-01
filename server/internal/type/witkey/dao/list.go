package dao_witkey

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	Name       string      `json:"name" dc:"威客名称"`
	User       *User       `json:"user" dc:"所属用户"`
	Game       string      `json:"game" dc:"游戏领域"`
	Title      string      `json:"title" dc:"头衔勋章"`
	Status     int         `json:"status" dc:"状态"`
	Rate       int         `json:"rate" dc:"评分"`
	CreateTime *gtime.Time `json:"createTime" dc:"入驻时间"`
}

type User struct {
	Name   string `json:"name" dc:"用户名称"`
	Avatar string `json:"avatar" dc:"用户头像"`
}
