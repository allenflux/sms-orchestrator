// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AccountLinkDevice is the golang structure for table account_link_device.
type AccountLinkDevice struct {
	Id           int         `json:"id"            orm:"id"            description:""`               //
	DevId        string      `json:"dev_id"        orm:"dev_id"        description:"设备ID"`           // 设备ID
	AccountId    int         `json:"account_id"    orm:"account_id"    description:"账号ID"`           // 账号ID
	DeviceId     int         `json:"device_id"     orm:"device_id"     description:"设备表主键ID"`        // 设备表主键ID
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:""`               //
	UpdatedAt    *gtime.Time `json:"updated_at"    orm:"updated_at"    description:""`               //
	PlatformType int         `json:"platform_type" orm:"platform_type" description:"平台： 1-facebook"` // 平台： 1-facebook
}
