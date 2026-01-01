package v1

import (
	dao_site "server/internal/type/site/dao"

	"github.com/gogf/gf/v2/frame/g"
)

// ---------------------- 基础配置
type GetSiteReq struct {
	g.Meta `path:"/site" method:"get" tags:"站点" summary:"获取系统站点信息"`
}
type GetSiteRes struct {
	*dao_site.Detail
}

type GetGameOptionsReq struct {
	g.Meta `path:"/site/game/options" method:"get" tags:"站点" summary:"游戏选项"`
}
type GetGameOptionsRes struct {
	List []*dao_site.Options `json:"list" dc:"游戏选项列表"`
}

type GetUserOptionsReq struct {
	g.Meta `path:"/site/user/options" method:"get" tags:"站点" summary:"用户选项"`
	Phone  string `p:"phone" v:"required#请输入用户手机号" dc:"手机号"`
}
type GetUserOptionsRes struct {
	List []*dao_site.Options `json:"list" dc:"用户选项列表"`
}

type GetTitleOptionsReq struct {
	g.Meta `path:"/site/title/options" method:"get" tags:"站点" summary:"所属头衔选项"`
	GameId int64 `p:"gameId" v:"required#请输入游戏id" dc:"游戏id"`
}
type GetTitleOptionsRes struct {
	List []*dao_site.Options `json:"list" dc:"所属头衔选项列表"`
}
