// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AccountGroup is the golang structure for table account_group.
type AccountGroup struct {
	Id           uint        `json:"id"            orm:"id"            description:""`                //
	Name         string      `json:"name"          orm:"name"          description:"账号分组"`            // 账号分组
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间"`            // 创建时间
	UpdatedAt    *gtime.Time `json:"updated_at"    orm:"updated_at"    description:"修改时间"`            // 修改时间
	ClientId     int         `json:"client_id"     orm:"client_id"     description:"客户ID"`            // 客户ID
	PlatformType int         `json:"platform_type" orm:"platform_type" description:"平台类型 1-facebook"` // 平台类型 1-facebook
}
