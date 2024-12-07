package sms

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "sms_backend/api/v1/common"
	"sms_backend/internal/model"
)

type SubTaskListReq struct {
	model.PageReq
	g.Meta    `path:"/sub/task/list" tags:"子平台群发短信" method:"get" dc:"查看群发短信列表" `
	ProjectID int      `json:"project_id" dc:"所属项目id" example:"1"`
	TaskName  string   `json:"task_name" dc:"任务名称"`
	DateRange []string `json:"date_range" p:"dateRange" description:"日期范围"`
}

type SubTaskListRes struct {
	commonApi.ListRes
	Data struct {
		ID                int    `json:"id" dc:"序号"`
		TaskID            int    `json:"task_id" dc:"任务id"`
		TaskName          string `json:"task_name" dc:"任务名称"`
		FileName          string `json:"file_name" dc:"文件名"`
		DeviceQuota       int    `json:"device_quota" dc:"执行设备"`
		TaskStatus        string `json:"task_status" dc:"任务状态"`
		SmsQuantity       string `json:"sms_quantity" dc:"SMS Quantity 短信总条数"`
		SurplusQuantity   string `json:"surplus_quantity" dc:"剩余数量"`
		QuantitySent      string `json:"quantity_sent" dc:"以发送数量"`
		AssociatedAccount string `json:"associated_account" dc:"所属子账号"`
		IntervalTime      string `json:"interval_time" dc:"间隔时间"`
		StartTime         string `json:"start_time" dc:"开始时间"`
		CreateTime        string `json:"create_time" dc:"创建时间"`
	} `json:"data"`
}
