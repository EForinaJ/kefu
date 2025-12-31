package dto_product

type Purchase struct {
	Id           int64  `p:"Id" v:"required|integer|min:1#请设置商品|商品id类型必须是整型|商品id最小为1" dc:"商品id"`
	UserId       int64  `p:"userId" v:"required|integer|min:1#请设置用户|用户id类型必须是整型|用户id最小为1" dc:"用户id"`
	Quantity     int    `p:"quantity" v:"required|integer|min:1#请设置下单数量|下单数量类型必须是整型|下单数量最小为1" dc:"下单数量"`
	Requirements string `p:"requirements" dc:"下单要求"`
}
