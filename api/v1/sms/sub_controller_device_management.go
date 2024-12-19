package sms

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "sms_backend/api/v1/common"
	"sms_backend/internal/model"
)

type SubDeviceListReq struct {
	g.Meta `path:"/sub/device/list" tags:"子平台设备列表" method:"get" dc:"查看设备列表" `
	model.PageReq
	SubUserID       int      `json:"sub_user_id" v:"required" dc:"子平台用户id"`
	SentDateRange   []string `json:"sent_date_range" p:"dateRange" description:"发送时间日期范围"`
	CreateDateRange []string `json:"create_date_range" p:"dateRange" description:"创建时间日期范围"`
	ProjectID       string   `json:"project_id" dc:"查询条件中的项目ID"`
	SentStatus      int      `json:"sent_status" dc:"需要查询的设备状态"`
	TaskName        string   `json:"task_name" dc:"任务名称"`
	Number          string   `json:"number" dc:"设备号码"`
	DeviceNumber    string   `json:"device_number" dc:"设备序列号"`
}

type SubDeviceListResData struct {
	ID            int    `json:"id"`
	ProjectID     int    `json:"project_id"`
	DeviceNumber  string `json:"device_number"`
	Number        string `json:"number"`
	ActiveDays    int    `json:"active_days"`
	OwnerAccount  string `json:"owner_account"`
	AssignedItems string `json:"assigned_items"`
	SentStatus    int    `json:"sent_status" dc:"1 空闲 2 异常 3 占用"`
	QuantitySent  string `json:"quantity_sent"`
	DeviceStatus  int    `json:"device_status" dc:"1 空闲 2 异常 3 占用"`
	ActiveTime    string `json:"active_time"`
}

type SubDeviceListRes struct {
	Data []SubDeviceListResData `json:"data"`
	commonApi.ListRes
}

// Sub Group
// Get Group
type SubGroupListReq struct {
	g.Meta    `path:"/sub/group/list" tags:"子平台设备列表" method:"get" dc:"查看当前用户下的分组列表" `
	SubUserID int `json:"sub_user_id"  v:"required" dc:"子平台用户账号ID"`
}

type SubGroupListResData struct {
	GroupName string `json:"group_name"`
	GroupID   int    `json:"group_id"`
}

type SubGroupListRes struct {
	Data []SubGroupListResData `json:"data"`
}

// Create Group

type SubCreateGroupReq struct {
	g.Meta    `path:"/sub/group" tags:"子平台设备列表" method:"post" dc:"创建当前用户下的分组列表" `
	SubUserID int    `json:"sub_user_id"  v:"required" dc:"子平台用户账号ID"`
	ProjectID int    `json:"project_id"  v:"required" dc:"平台当前子用户下的projectID"`
	GroupName string `json:"group_name" v:"required"`
}
type SubCreateGroupRes struct {
	ID int64 `json:"id"`
}

// Update Group
type SubUpdateGroupReq struct {
	g.Meta    `path:"/sub/group" tags:"子平台设备列表" method:"put" dc:"更新当前用户下的分组列表" `
	SubUserID int    `json:"sub_user_id"  v:"required" dc:"子平台用户账号ID"`
	GroupID   int    `json:"group_id" v:"required"`
	GroupName string `json:"group_name" v:"required"`
}
type SubUpdateGroupRes struct{}

// Delete Group
type SubDeleteGroupReq struct {
	g.Meta  `path:"/sub/group" tags:"子平台设备列表" method:"delete" dc:"删除当前用户下的分组列表" `
	GroupID int `json:"group_id" v:"required"`
}
type SubDeleteGroupRes struct{}

//AllocateDevice2Group

type AllocateDevice2GroupReq struct {
	g.Meta       `path:"/sub/device/group" tags:"子平台设备列表" method:"post" dc:"分配设备给分组" `
	SubUserID    int      `json:"sub_user_id"  v:"required" dc:"子平台用户账号ID"`
	DeviceIdList []string `json:"device_id_list" v:"required"`
	GroupID      int      `json:"group_id" v:"required"`
}
type AllocateDevice2GroupRes struct {
}
