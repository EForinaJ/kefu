package dto_witkey

type Create struct {
	Name        string   `p:"name" v:"required|max-length:8#请输入威客昵称|长度最长为8" dc:"昵称"`
	TitleId     int64    `p:"titleId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"头衔id"`
	GameId      int64    `p:"gameId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"所属分类id"`
	Phone       string   `p:"phone" v:"required|phone#请设置手机号码|手机号码格式错误"     dc:"手机号码"`
	Address     []string `p:"address"  dc:"地址"`
	Description string   `p:"description"  dc:"描述"`
	Sex         int      `p:"sex"  dc:"性别"`
	Birthday    int64    `p:"birthday" dc:"生日"`
	Password    string   `p:"password"  dc:"密码"`
	Status      int      `p:"status"  dc:"状态"`
}
