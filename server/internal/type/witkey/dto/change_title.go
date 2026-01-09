package dto_witkey

type ChangeTitle struct {
	Id      int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"头衔id"`
	TitleId int64 `p:"titleId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"头衔id"`
	GameId  int64 `p:"gameId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"所属分类id"`
	// Type        int    `p:"type" v:"required|in:1,2,3,4#请选择类型|类型值只能在1,2,3,4" dc:"类型"`
	Description string `p:"description" v:"required#请输入变更描述"  dc:"变更描述"`
}
