// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ChatTaskLinkCollectDataGroup is the golang structure for table chat_task_link_collect_data_group.
type ChatTaskLinkCollectDataGroup struct {
	Id                 uint `json:"id"                    orm:"id"                    description:""`        //
	ChatTaskId         int  `json:"chat_task_id"          orm:"chat_task_id"          description:"私信任务id"`  // 私信任务id
	CollectDataGroupId int  `json:"collect_data_group_id" orm:"collect_data_group_id" description:"采集数据组id"` // 采集数据组id
}
