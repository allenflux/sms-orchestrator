// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// InstallTaskLinkDeviceGroup is the golang structure for table install_task_link_device_group.
type InstallTaskLinkDeviceGroup struct {
	Id            uint `json:"id"              orm:"id"              description:""`       //
	InstallTaskId int  `json:"install_task_id" orm:"install_task_id" description:"安装任务id"` // 安装任务id
	DeviceGroupId int  `json:"device_group_id" orm:"device_group_id" description:"设备组id"`  // 设备组id
}
