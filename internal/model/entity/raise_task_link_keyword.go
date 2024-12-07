// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RaiseTaskLinkKeyword is the golang structure for table raise_task_link_keyword.
type RaiseTaskLinkKeyword struct {
	Id          uint `json:"id"            orm:"id"            description:""`       //
	RaiseTaskId int  `json:"raise_task_id" orm:"raise_task_id" description:"养号任务id"` // 养号任务id
	KeywordId   int  `json:"keyword_id"    orm:"keyword_id"    description:"关键词id"`  // 关键词id
}
