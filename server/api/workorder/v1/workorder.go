package v1

import (
	dao_workorder "server/internal/type/workorder/dao"
	dto_workorder "server/internal/type/workorder/dto"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta `path:"/workorder/create" method:"post" tags:"工单" summary:"上报工单"`
	*dto_workorder.Create
}
type CreateRes struct{}

type GetListReq struct {
	g.Meta `path:"/workorder/list" method:"get" tags:"工单" summary:"工单列表"`
	*dto_workorder.Query
}
type GetListRes struct {
	Total int                   `json:"total" dc:"总数"`
	List  []*dao_workorder.List `json:"list" dc:"工单列表"`
}

type CancelReq struct {
	g.Meta `path:"/workorder/cancel" method:"post" tags:"工单" summary:"取消工单"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type CancelRes struct{}

type SolveReq struct {
	g.Meta `path:"/workorder/solve" method:"post" tags:"工单" summary:"解决工单"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type SolveRes struct{}
