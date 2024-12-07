// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ChatTaskLinkDeviceGroup is the golang structure for table chat_task_link_device_group.
type ChatTaskLinkDeviceGroup struct {
	Id            uint `json:"id"              orm:"id"              description:""`       //
	ChatTaskId    int  `json:"chat_task_id"    orm:"chat_task_id"    description:"私信任务id"` // 私信任务id
	DeviceGroupId int  `json:"device_group_id" orm:"device_group_id" description:"设备分组id"` // 设备分组id
}
