package v1

import (
	dao_witkey "kefu-server/internal/type/witkey/dao"
	dto_witkey "kefu-server/internal/type/witkey/dto"

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
