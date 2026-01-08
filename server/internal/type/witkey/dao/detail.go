package dao_witkey

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id         int64       `json:"id" dc:"ID"`
	Avatar     string      `json:"avatar" dc:"威客头像"`
	Address    []string    `json:"address" dc:"所在地址"`
	Birthday   *gtime.Time `json:"birthday" dc:"出生日期"`
	Sex        int         `json:"sex" dc:"威客性别"`
	Name       string      `json:"name" dc:"威客昵称"`
	Game       string      `json:"game" dc:"游戏领域"`
	Title      string      `json:"title" dc:"头衔勋章"`
	Album      []string    `json:"album" dc:"相册"`
	Rate       int         `json:"rate" dc:"评分"`
	Status     int         `json:"status" dc:"接单状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}
