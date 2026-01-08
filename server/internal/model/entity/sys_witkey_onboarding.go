// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkeyOnboarding is the golang structure for table sys_witkey_onboarding.
type SysWitkeyOnboarding struct {
	Id          int64       `json:"id"          orm:"id"          description:"用户ID"` // 用户ID
	Name        string      `json:"name"        orm:"name"        description:""`     //
	TitleId     int64       `json:"titleId"     orm:"title_id"    description:""`     //
	GameId      int64       `json:"gameId"      orm:"game_id"     description:""`     //
	Phone       string      `json:"phone"       orm:"phone"       description:""`     //
	Address     string      `json:"address"     orm:"address"     description:""`     //
	Sex         int         `json:"sex"         orm:"sex"         description:""`     //
	Description string      `json:"description" orm:"description" description:""`     //
	Status      int         `json:"status"      orm:"status"      description:""`     //
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time" description:"创建时间"` // 创建时间
}
