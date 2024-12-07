// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AllocationTaskFails is the golang structure for table allocation_task_fails.
type AllocationTaskFails struct {
	Id        int         `json:"id"         orm:"id"         description:""`      //
	AccountId int         `json:"account_id" orm:"account_id" description:"账号ID"`  // 账号ID
	DevId     string      `json:"dev_id"     orm:"dev_id"     description:"设备ID"`  // 设备ID
	TaskId    int         `json:"task_id"    orm:"task_id"    description:"任务表ID"` // 任务表ID
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`    // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`      //
	Username  string      `json:"username"   orm:"username"   description:"账号"`    // 账号
}
