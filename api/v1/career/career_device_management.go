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
type RegisterRes struct{}

// Reporting Device Information

type ReportInfoReq struct {
	g.Meta       `path:"/device/info" tags:"Device使用的API" method:"post" dc:"根据设备号获取设备信息"`
	DeviceNumber string `json:"device_number" dc:"设备号" v:"required"`
}
type ReportInfoRes struct{}

// Fetch Task

type FetchTaskReq struct {
	g.Meta       `path:"/device/task" tags:"Device使用的API" method:"post" dc:"根据设备号获取设备任务"`
	DeviceNumber string `json:"device_number" dc:"设备号" v:"required"`
}
type FetchTaskRes struct{}

// Report Task Results

type ReportTaskResultReq struct {
	g.Meta       `path:"/device/task/result" tags:"Device使用的API" method:"post" dc:"根据设备号上报任务执行结果"`
	DeviceNumber string `json:"device_number" dc:"设备号" v:"required"`
	SendStatus   int    `json:"send_status" dc:"1-成功 2-失败" v:"required"`
	Reason       string `json:"reason" dc:"如果不成功告知失败原因"`
}
type ReportTaskResultRes struct{}

// Report Receive Content

type ReportReceiveContentReq struct {
	g.Meta            `path:"/device/task/content" tags:"Device使用的API" method:"post" dc:"根据设备号上报接收的内容"`
	DeviceNumber      string `json:"device_number" dc:"设备号" v:"required"`
	SmsContent        string `json:"sms_content" v:"required"`
	SmsTopic          string `json:"sms_topic" v:"required"`
	ReceiveTime       string `json:"receive_time" v:"required" dc:"接收时间戳"`
	TargetPhoneNumber string `json:"target_phone_number" v:"required" dc:"接收的手机号"`
}
type ReportReceiveContentRes struct{}
