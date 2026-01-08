// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysWitkeyChangeTitle is the golang structure for table sys_witkey_change_title.
type SysWitkeyChangeTitle struct {
	Id          int64       `json:"id"          orm:"id"           description:""` //
	ManageId    int64       `json:"manageId"    orm:"manage_id"    description:""` //
	WitkeyId    int64       `json:"witkeyId"    orm:"witkey_id"    description:""` //
	OldGameId   int64       `json:"oldGameId"   orm:"old_game_id"  description:""` //
	OldTitleId  int64       `json:"oldTitleId"  orm:"old_title_id" description:""` //
	NewGameId   int64       `json:"newGameId"   orm:"new_game_id"  description:""` //
	NewTitleId  int64       `json:"newTitleId"  orm:"new_title_id" description:""` //
	Type        int         `json:"type"        orm:"type"         description:""` //
	Description string      `json:"description" orm:"description"  description:""` //
	Status      int         `json:"status"      orm:"status"       description:""` //
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:""` //
}
