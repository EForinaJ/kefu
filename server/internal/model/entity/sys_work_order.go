// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWorkOrder is the golang structure for table sys_work_order.
type SysWorkOrder struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	CreatorId  int64       `json:"creatorId"  orm:"creator_id"  description:""` //
	AssigneeId int64       `json:"assigneeId" orm:"assignee_id" description:""` //
	Code       string      `json:"code"       orm:"code"        description:""` //
	Priority   int         `json:"priority"   orm:"priority"    description:""` //
	Title      string      `json:"title"      orm:"title"       description:""` //
	Content    string      `json:"content"    orm:"content"     description:""` //
	Status     int         `json:"status"     orm:"status"      description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
}
