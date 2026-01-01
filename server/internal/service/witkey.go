package service

import (
	"context"
	dao_witkey "server/internal/type/witkey/dao"
	dto_witkey "server/internal/type/witkey/dto"
)

// 定义显示接口
type IWitkey interface {
	GetList(ctx context.Context, req *dto_witkey.Query) (total int, res []*dao_witkey.List, err error)
}

// 定义接口变量
var localWitkey IWitkey

// 定义一个获取接口的方法
func Witkey() IWitkey {
	if localWitkey == nil {
		panic("implement not found for interface IWitkey, forgot register?")
	}
	return localWitkey
}

// 定义一个接口实现的注册方法
func RegisterWitkey(i IWitkey) {
	localWitkey = i
}
