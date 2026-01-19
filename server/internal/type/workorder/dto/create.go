package dto_workorder

type Create struct {
	Title    string `p:"title" v:"required#请输入工单标题" dc:"工单标题"`
	Content  string `p:"content" v:"required#请输入上报内容"     dc:"上报内容"`
	Priority int    `p:"priority"  dc:"优先级"`
}
