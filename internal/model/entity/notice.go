// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notice is the golang structure for table notice.
type Notice struct {
	Id                  uint        `json:"id"                     orm:"id"                     description:""`                             //
	UserType            int         `json:"user_type"              orm:"user_type"              description:"推送目标用户类型1管理端用户2企业端用户"`         // 推送目标用户类型1管理端用户2企业端用户
	UserId              int         `json:"user_id"                orm:"user_id"                description:"推送目标用户id"`                     // 推送目标用户id
	OrderId             int         `json:"order_id"               orm:"order_id"               description:"订单id"`                         // 订单id
	Type                string      `json:"type"                   orm:"type"                   description:"通知类型"`                         // 通知类型
	Title               string      `json:"title"                  orm:"title"                  description:"标题"`                           // 标题
	Content             string      `json:"content"                orm:"content"                description:"内容"`                           // 内容
	Read                int         `json:"read"                   orm:"read"                   description:"是否已读0未读1已读"`                   // 是否已读0未读1已读
	CreatedAt           *gtime.Time `json:"created_at"             orm:"created_at"             description:"推送时间"`                         // 推送时间
	DeletedAt           *gtime.Time `json:"deleted_at"             orm:"deleted_at"             description:"删除时间"`                         // 删除时间
	DeviceGrantImportId int         `json:"device_grant_import_id" orm:"device_grant_import_id" description:"设备授权导入任务id"`                   // 设备授权导入任务id
	ImportType          int         `json:"import_type"            orm:"import_type"            description:"1-授权导入  2-终端设备标签导入 3-app标签导入"` // 1-授权导入  2-终端设备标签导入 3-app标签导入
}
