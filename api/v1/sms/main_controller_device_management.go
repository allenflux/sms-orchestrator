package sms

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "sms_backend/api/v1/common"
	"sms_backend/internal/model"
)

type ProjectListReq struct {
	g.Meta `path:"/project/list" tags:"项目管理" method:"get" dc:"查看项目列表" `
	model.PageReq
	KeyWordSearch string `json:"key_word_search" dc:"关键字查询"`
}

type ProjectListResData struct {
	ID                  int    `json:"id"`
	Name                string `json:"name" `
	QuantityDevice      int    `json:"quantity_device"`
	AssociatedAccountId int    `json:"associated_account_id"`
	AssociatedAccount   string `json:"associated_account"`
	Note                string `json:"note"`
	UpdateTime          string `json:"update_time"`
}

type ProjectListRes struct {
	Data []ProjectListResData `json:"data"`
	commonApi.ListRes
}

type ProjectListForFrontReq struct {
	g.Meta    `path:"/project/items" tags:"项目管理" method:"get" dc:"查询项目列表 给前端作为筛选条件使用" `
	SubUserID int `json:"sub_user_id" default:"0" dc:"这个字段在子平台查询时使用 如果传递此字段 将返回子平台下的项目 总平台不用传递"`
}
type ProjectListForFrontResData struct {
	ProjectName string `json:"project_name"`
	ProjectId   int    `json:"project_id"`
}
type ProjectListForFrontRes struct {
	Data []*ProjectListForFrontResData `json:"data"`
}

type ProjectCreateReq struct {
	g.Meta      `path:"/project" tags:"项目管理" method:"post" dc:"新增项目" `
	ProjectName string `json:"project_name" v:"required"`
	Note        string `json:"note" `
}

type ProjectCreateRes struct {
	ID int64 `json:"id"`
}

type ProjectUpdateReq struct {
	g.Meta      `path:"/project" tags:"项目管理" method:"put" dc:"编辑项目" `
	ProjectID   int    `json:"project_id" v:"required"`
	ProjectName string `json:"project_name" `
	Note        string `json:"note" `
}

type ProjectUpdateRes struct {
}

type ProjectDeleteReq struct {
	g.Meta    `path:"/project" tags:"项目管理" method:"delete" dc:"删除项目" `
	ProjectID int `json:"project_id" v:"required"`
}
type ProjectDeleteRes struct {
}

type DeviceListReq struct {
	g.Meta `path:"/device/list" tags:"设备列表" method:"get" dc:"查看设备列表" `
	model.PageReq
	SentDateRange   []string `json:"sent_date_range" p:"dateRange" description:"发送时间日期范围"`
	CreateDateRange []string `json:"create_date_range" p:"dateRange" description:"创建时间日期范围"`
	ProjectID       string   `json:"project_id" dc:"查询条件中的项目ID"`
	SentStatus      int      `json:"sent_status" dc:"需要查询的设备状态" default:"0" dc:"发送状态，1-异常 2-占用 3-空闲"`
	TaskName        string   `json:"task_name" dc:"任务名称"`
	Number          string   `json:"number" dc:"设备号码"`
	DeviceID        string   `json:"device_id" dc:"设备标识"`
	DeviceNumber    string   `json:"device_number" dc:"设备序列号"`
}
type DeviceListResData struct {
	ID            int    `json:"id"`
	ProjectID     int    `json:"project_id"`
	DeviceNumber  string `json:"device_number"`
	DeviceID      string `json:"device_id"`
	Number        string `json:"number"`
	ActiveDays    int    `json:"active_days"`
	OwnerAccount  string `json:"owner_account"`
	AssignedItems string `json:"assigned_items"`
	SentStatus    int    `json:"sent_status" dc:"1 空闲 2 异常 3 占用"`
	QuantitySent  string `json:"quantity_sent"`
	DeviceStatus  int    `json:"device_status" dc:"1 空闲 2 异常 3 占用"`
	ActiveTime    string `json:"active_time"`
}
type DeviceListRes struct {
	Data []DeviceListResData `json:"data"`
	commonApi.ListRes
}

type AllocateDevice2ProjectReq struct {
	g.Meta       `path:"/device/project" tags:"设备列表" method:"post" dc:"分配设备给项目" `
	DeviceIdList []int `json:"device_id_list" v:"required"`
	ProjectID    int   `json:"project_id" v:"required"`
}
type AllocateDevice2ProjectRes struct {
}

type AllocateAccount2ProjectReq struct {
	g.Meta    `path:"/account/project" tags:"设备列表" method:"post" dc:"分配子账号给项目" `
	AccountId string `json:"account_id" v:"required" dc:"子账号id"`
	ProjectId int    `json:"project_id" v:"required"`
}
type AllocateAccount2ProjectRes struct {
}
