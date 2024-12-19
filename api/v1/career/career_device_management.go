package career

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Register with the platform

type RegisterReq struct {
	g.Meta       `path:"/device" tags:"Device使用的API" method:"post" dc:"注册设备到平台" `
	DeviceNumber string      `json:"device_number" dc:"设备号" v:"required"`
	PhoneNumber  string      `json:"phone_number" dc:"设备手机号" v:"required"`
	ActiveTime   *gtime.Time `json:"active_time" dc:"激活时间 时间戳" v:"required"`
}
type RegisterRes struct {
	ID int64 `json:"id"`
}

// Reporting Device Information

//type ReportInfoReq struct {
//	g.Meta       `path:"/device/info" tags:"Device使用的API" method:"post" dc:"根据设备号上报设备信息"`
//	DeviceNumber string `json:"device_number" dc:"设备号" v:"required"`
//	Status       string `json:"status" dc:"设备状态" v:"required"`
//}
//type ReportInfoRes struct{}

// Fetch Task

type FetchTaskReq struct {
	g.Meta       `path:"/device/task" tags:"Device使用的API" method:"post" dc:"根据设备号获取设备任务，每次调用只取回一个任务"`
	DeviceNumber string `json:"device_number" dc:"设备号" v:"required"`
}

type FetchTaskRes struct {
	TargetPhoneNumber string      `json:"target_phone_number"`
	Content           string      `json:"content"`
	DeviceNumber      string      `json:"device_number" dc:"有时为空 若不为空则需要使用特定设备发送此条信息"`
	Interval          string      `json:"interval"`
	TaskId            int         `json:"task_id"`
	StartAt           *gtime.Time `json:"start_at"`
}

//type FetchTaskByGroupRes struct {
//	Data []FetchTaskResData `json:"data"`
//}

// Report Task Results

type ReportTaskResultReq struct {
	g.Meta            `path:"/device/task/result" tags:"Device使用的API" method:"post" dc:"根据设备号上报任务执行结果"`
	DeviceNumber      string      `json:"device_number" dc:"设备号" v:"required"`
	Content           string      `json:"content" v:"required"`
	SendStatus        int         `json:"send_status" dc:"1-成功 2-失败" v:"required"`
	Reason            string      `json:"reason" dc:"如果不成功告知失败原因" v:"required"`
	TargetPhoneNumber string      `json:"target_phone_number" v:"required"`
	TaskId            int         `json:"task_id" v:"required"`
	StartTime         *gtime.Time `json:"start_time" v:"required"`
}
type ReportTaskResultRes struct {
	ID int64 `json:"id"`
}

// Report Receive Content

type ReportReceiveContentReq struct {
	g.Meta `path:"/device/task/content" tags:"Device使用的API" method:"post" dc:"根据设备号上报接收的内容"`
	//TaskId            int         `json:"task_id" v:"required" dc:"任务id platform需要记录任务的关联"`
	DeviceNumber      string      `json:"device_number" dc:"设备号" v:"required"`
	TargetPhoneNumber string      `json:"target_phone_number" dc:"设备号" v:"required"`
	SmsContent        string      `json:"sms_content" v:"required" dc:"接收的内容"`
	StartTime         *gtime.Time `json:"start_time" v:"required" dc:"设备内容接收时间"`
}
type ReportReceiveContentRes struct {
	ID int64 `json:"id"`
}
