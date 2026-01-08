package v1

import (
	dao_witkey "server/internal/type/witkey/dao"
	dto_witkey "server/internal/type/witkey/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/witkey/list" method:"get" tags:"威客" summary:"威客列表"`
	*dto_witkey.Query
}
type GetListRes struct {
	Total int                `json:"total" dc:"总数"`
	List  []*dao_witkey.List `json:"list" dc:"威客列表"`
}

type CreateReq struct {
	g.Meta `path:"/witkey/create" method:"post" tags:"威客" summary:"创建威客"`
	*dto_witkey.Create
}
type CreateRes struct{}

type ChangeTitleReq struct {
	g.Meta `path:"/witkey/change/title" method:"post" tags:"威客" summary:"修改威客头衔"`
	*dto_witkey.ChangeTitle
}
type ChangeTitleRes struct{}

type GetDetailReq struct {
	g.Meta `path:"/witkey/detail" method:"get" tags:"威客" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_witkey.Detail
}
