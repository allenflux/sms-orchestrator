// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceList is the golang structure for table device_list.
type DeviceList struct {
	Id              int         `json:"id"                orm:"id"                description:""`                    //
	DeviceId        string      `json:"device_id"         orm:"device_id"         description:"设备id"`                // 设备id
	Number          string      `json:"number"            orm:"number"            description:"设备号码"`                // 设备号码
	ActiveTime      *gtime.Time `json:"active_time"       orm:"active_time"       description:"激活时间"`                // 激活时间
	OwnerAccount    string      `json:"owner_account"     orm:"owner_account"     description:"子账号"`                 // 子账号
	AssignedItems   string      `json:"assigned_items"    orm:"assigned_items"    description:"所属项目"`                // 所属项目
	SentStatus      int         `json:"sent_status"       orm:"sent_status"       description:"发送状态，1-异常 2-占用 3-空闲"` // 发送状态，1-异常 2-占用 3-空闲
	QuantitySent    int         `json:"quantity_sent"     orm:"quantity_sent"     description:"发送数量"`                // 发送数量
	QuantityAll     int         `json:"quantity_all"      orm:"quantity_all"      description:"全部数量"`                // 全部数量
	DeviceStatus    int         `json:"device_status"     orm:"device_status"     description:"设备状态，1-异常 2-正常"`      // 设备状态，1-异常 2-正常
	GroupName       string      `json:"group_name"        orm:"group_name"        description:"分组名称"`                // 分组名称
	GroupId         int         `json:"group_id"          orm:"group_id"          description:"分组id"`                // 分组id
	CreatedAt       *gtime.Time `json:"created_at"        orm:"created_at"        description:"创建时间"`                // 创建时间
	UpdateAt        *gtime.Time `json:"update_at"         orm:"update_at"         description:"修改时间"`                // 修改时间
	DeleteAt        *gtime.Time `json:"delete_at"         orm:"delete_at"         description:"删除时间"`                // 删除时间
	DeviceNumber    string      `json:"device_number"     orm:"device_number"     description:"设备序列号"`               // 设备序列号
	StartAt         *gtime.Time `json:"start_at"          orm:"start_at"          description:"发送时间"`                // 发送时间
	AssignedItemsId int         `json:"assigned_items_id" orm:"assigned_items_id" description:"项目id"`                // 项目id
	OwnerAccountId  int         `json:"owner_account_id"  orm:"owner_account_id"  description:"子账号id"`               // 子账号id
}
