// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatTask is the golang structure for table chat_task.
type ChatTask struct {
	Id              uint        `json:"id"                orm:"id"                description:""`                                  //
	Name            string      `json:"name"              orm:"name"              description:"任务名称"`                              // 任务名称
	ChatNum         int         `json:"chat_num"          orm:"chat_num"          description:"私信数量"`                              // 私信数量
	ChatContentType int         `json:"chat_content_type" orm:"chat_content_type" description:"私信内容类型1话术+自定义2话术3自定义"`              // 私信内容类型1话术+自定义2话术3自定义
	WordsGroupId    int         `json:"words_group_id"    orm:"words_group_id"    description:"话术组id"`                             // 话术组id
	CustomContent   string      `json:"custom_content"    orm:"custom_content"    description:"自定义内容"`                             // 自定义内容
	SendGap         int         `json:"send_gap"          orm:"send_gap"          description:"发生时间间隔秒数"`                          // 发生时间间隔秒数
	CreatedAt       *gtime.Time `json:"created_at"        orm:"created_at"        description:"创建时间"`                              // 创建时间
	CompletedAt     *gtime.Time `json:"completed_at"      orm:"completed_at"      description:"完成时间"`                              // 完成时间
	Status          int         `json:"status"            orm:"status"            description:"任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成"` // 任务状态1待执行2执行中3客户取消4客服取消5已暂停6异常7已完成
	Remark          string      `json:"remark"            orm:"remark"            description:"备注"`                                // 备注
	ClientId        int         `json:"client_id"         orm:"client_id"         description:"用户ID"`                              // 用户ID
	HasEmpty        int         `json:"has_empty"         orm:"has_empty"         description:"1-未清空 2-已清空"`                       // 1-未清空 2-已清空
}
