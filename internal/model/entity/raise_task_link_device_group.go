// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RaiseTaskLinkDeviceGroup is the golang structure for table raise_task_link_device_group.
type RaiseTaskLinkDeviceGroup struct {
	Id            uint `json:"id"              orm:"id"              description:""`       //
	RaiseTaskId   int  `json:"raise_task_id"   orm:"raise_task_id"   description:"养号任务id"` // 养号任务id
	DeviceGroupId int  `json:"device_group_id" orm:"device_group_id" description:"任务组id"`  // 任务组id
}
