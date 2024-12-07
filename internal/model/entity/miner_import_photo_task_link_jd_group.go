// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MinerImportPhotoTaskLinkJdGroup is the golang structure for table miner_import_photo_task_link_jd_group.
type MinerImportPhotoTaskLinkJdGroup struct {
	Id                     int         `json:"id"                         orm:"id"                         description:""`                                      //
	MinerImportPhotoTaskId int         `json:"miner_import_photo_task_id" orm:"miner_import_photo_task_id" description:"导入相册任务ID"`                              // 导入相册任务ID
	JdGroupName            string      `json:"jd_group_name"              orm:"jd_group_name"              description:"金盾分组名称"`                                // 金盾分组名称
	Num                    int         `json:"num"                        orm:"num"                        description:"分组数量"`                                  // 分组数量
	CreatedAt              *gtime.Time `json:"created_at"                 orm:"created_at"                 description:""`                                      //
	MinerTaskType          int         `json:"miner_task_type"            orm:"miner_task_type"            description:"1-导入相册 2-导入通讯录 3-矿机登录 4-日常任务 5-定时任务计划"` // 1-导入相册 2-导入通讯录 3-矿机登录 4-日常任务 5-定时任务计划
}
