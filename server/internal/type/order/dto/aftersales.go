package dto_order

type Aftersales struct {
	Type          int     `p:"type" v:"required|in:1,2,3#请选择售后类型|售后类型不对" dc:"售后类型"`
	Amount        float64 `p:"amount" v:"required#请设置金额"`
	Reason        string  `p:"reason" v:"required#请输入售后原因" dc:"售后原因"`
	Id            int64   `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	ContactNumber string  `p:"contactNumber" v:"required#请输入联系号码" dc:"联系号码"`
	ContactMode   int     `p:"contactMode" v:"required|in:1,2,3#请选择联系方式|联系方式不对" dc:"联系方式"`
	Content       string  `p:"content" v:"required#请输入售后内容" dc:"售后内容"`
}
