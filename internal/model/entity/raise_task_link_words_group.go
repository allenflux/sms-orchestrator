// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RaiseTaskLinkWordsGroup is the golang structure for table raise_task_link_words_group.
type RaiseTaskLinkWordsGroup struct {
	Id           uint `json:"id"             orm:"id"             description:""`       //
	RaiseTaskId  int  `json:"raise_task_id"  orm:"raise_task_id"  description:"养号任务id"` // 养号任务id
	WordsGroupId int  `json:"words_group_id" orm:"words_group_id" description:"话术分组id"` // 话术分组id
}
