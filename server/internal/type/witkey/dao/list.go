package dao_witkey

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	User       *User       `json:"user" dc:"所属用户"`
	Game       string      `json:"game" dc:"游戏领域"`
	Title      string      `json:"title" dc:"头衔勋章"`
	Commission float64     `json:"commission" dc:"佣金"`
	Rate       int         `json:"rate" dc:"评分"`
	CreateTime *gtime.Time `json:"createTime" dc:"入驻时间"`
}

type User struct {
	Name   string `json:"name" dc:"用户名称"`
	Avatar string `json:"avatar" dc:"用户头像"`
	Phone  string `json:"phone" dc:"用户手机号"`
}
