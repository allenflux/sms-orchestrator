package sms

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	commonApi "sms_backend/api/v1/common"
	"sms_backend/internal/model"
)

type TaskListReq struct {
	model.PageReq
	g.Meta    `path:"/task/list" tags:"群发短信" method:"get" dc:"查看群发短信列表" `
	ProjectID int      `json:"project_id" dc:"所属项目id" example:"1"`
	TaskName  string   `json:"task_name" dc:"任务名称"`
	DateRange []string `json:"date_range" p:"dateRange" description:"日期范围"`
}

type TaskListResData struct {
	ID                int    `json:"id" dc:"序号"`
	ProjectID         int    `json:"project_id" dc:"项目id"`
	ProjectName       string `json:"project_name" dc:"项目name"`
	TaskName          string `json:"task_name" dc:"任务名称"`
	FileName          string `json:"file_name" dc:"文件名"`
	DeviceQuota       string `json:"device_quota" dc:"执行设备"`
	TaskStatus        int    `json:"task_status" dc:"任务状态任务状态，1-待发送 2-发送中 3-已发送 4-已撤销"`
	SmsQuantity       int    `json:"sms_quantity" dc:"SMS Quantity 短信总条数"`
	SurplusQuantity   int    `json:"surplus_quantity" dc:"剩余数量"`
	QuantitySent      int    `json:"quantity_sent" dc:"以发送数量"`
	AssociatedAccount string `json:"associated_account" dc:"所属子账号"`
	IntervalTime      string `json:"interval_time" dc:"间隔时间"`
	StartTime         string `json:"start_time" dc:"开始时间"`
	CreateTime        string `json:"create_time" dc:"创建时间"`
}

type TaskListRes struct {
	commonApi.ListRes
	Data []TaskListResData `json:"data"`
}

type TaskRecordReq struct {
	model.PageReq
	g.Meta            `path:"/task/record" tags:"群发记录" method:"get" dc:"查看群发短信列表" `
	ProjectID         int      `json:"project_id" dc:"所属项目id" example:"1"`
	SmsStatus         int      `json:"sms_status" dc:"SMS Status 发送状态 1发送成功 2发送失败" `
	TaskName          string   `json:"task_name" dc:"任务名称"`
	TargetPhoneNumber string   `json:"target_phone_number" dc:"Target Phone Number 目标手机号"`
	DeviceNumber      string   `json:"device_number" dc:"Device Number 执行设备号"`
	SentDateRange     []string `json:"sent_date_range" p:"dateRange 1" description:"发送日期范围"`
	CreateDateRange   []string `json:"create_date_range" p:"dateRange 2" description:"创建日期范围"`
}

type TaskRecordResData struct {
	ID                int    `json:"id" dc:"序号"`
	TaskName          string `json:"task_name" dc:"任务名称"`
	SubTaskID         int    `json:"sub_task_id" dc:"Sub Task ID 子任务ID"`
	TargetPhoneNumber string `json:"target_phone_number" dc:"Target Phone Number 目标手机号"`
	DeviceNumber      string `json:"device_number" dc:"Device Number 执行设备号"`
	SmsTopic          string `json:"sms_topic" dc:"SMS Topic 主题"`
	Reason            string `json:"reason" dc:"消息发送失败原因"`
	SmsContent        string `json:"sms_content" dc:"SMS Content 短信内容"`
	SmsStatus         string `json:"sms_status" dc:"SMS Status 短信发送状态"`
	AssociatedAccount string `json:"associated_account" dc:"所属子账号"`
	ProjectName       string `json:"project_name" dc:"Project Name 所属项目"`
	StartTime         string `json:"start_time" dc:"开始时间"`
	CreateTime        string `json:"create_time" dc:"创建时间"`
}

type TaskRecordRes struct {
	commonApi.ListRes
	Data []TaskRecordResData `json:"data"`
}

type ConversationListReq struct {
	g.Meta `path:"/conversation/list" tags:"消息对话" method:"get" dc:"查看群发短信列表" `
	model.PageReq
	ProjectID          int    `json:"project_id" v:"required"`
	SearchWord         string `json:"search_word" dc:"搜索对话"`
	ConversationStatus int    `json:"conversation_status" dc:"对话状态 1-所有对话 2-正在对话" default:"1"`
}
type ConversationListRes struct {
	*SubGetConversationRecordListRes
}

type ConversationRecordReq struct {
	g.Meta    `path:"/conversation/record" tags:"消息对话" method:"get" dc:"单点对话记录" `
	ChatLogID int `json:"chat_log_id" v:"required" dc:"需要查看的chart id"`
}

type ConversationRecordRes struct {
	*SubGetConversationRecordRes
}

type TaskDevicesReq struct {
	model.PageReq
	g.Meta `path:"/task/device/list" tags:"群发短信" method:"get" dc:"根据任务id获取执行设备列表" `
	TaskID int `json:"task_id" v:"required" dc:"任务id"`
}

type TaskDevicesResData struct {
	ID        int    `json:"id"`
	TaskID    int    `json:"task_id" dc:"任务id"`
	DeviceNum string `json:"device_num"`
}

type TaskDevicesRes struct {
	commonApi.ListRes
	Data []TaskDevicesResData `json:"data"`
}

type SubTaskDevicesReq struct {
	model.PageReq
	g.Meta `path:"/sub/task/device/list" tags:"群发短信" method:"get" dc:"根据任务id获取执行设备列表" `
	TaskID int `json:"task_id" v:"required" dc:"任务id"`
}

type SubTaskDevicesRes struct {
	commonApi.ListRes
	Data []TaskDevicesResData `json:"data"`
}

type PendingTaskReq struct {
	g.Meta `path:"/task/device/pending/list" tags:"群发短信" method:"get" dc:"待发送任务列表" `
	TaskID int `json:"task_id" dc:"任务id"`
}

type PendingTaskRes struct {
	commonApi.ListRes
	Data []PendingTaskResData `json:"data"`
}

type SubPendingTaskReq struct {
	g.Meta `path:"/sub/task/device/pending/list" tags:"子平台群发短信" method:"get" dc:"待发送任务列表" `
	TaskID int `json:"task_id" dc:"任务id"`
}

type PendingTaskResData struct {
	TaskID            int         `json:"task_id" dc:"任务id"`
	TargetPhoneNumber string      `json:"target_phone_number"`
	Content           string      `json:"content"`
	StartAt           *gtime.Time `json:"start_at"`
}

type SubPendingTaskRes struct {
	commonApi.ListRes
	Data []PendingTaskResData `json:"data"`
}
