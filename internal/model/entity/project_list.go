// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProjectList is the golang structure for table project_list.
type ProjectList struct {
	Id                  int         `json:"id"                    orm:"id"                    description:""`      //
	ProjectName         string      `json:"project_name"          orm:"project_name"          description:"项目名称"`  // 项目名称
	QuantityDevice      int         `json:"quantity_device"       orm:"quantity_device"       description:"设备数量"`  // 设备数量
	AssociatedAccount   string      `json:"associated_account"    orm:"associated_account"    description:"关联账号"`  // 关联账号
	Note                string      `json:"note"                  orm:"note"                  description:"备注"`    // 备注
	CreatedAt           *gtime.Time `json:"created_at"            orm:"created_at"            description:"创建时间"`  // 创建时间
	UpdateAt            *gtime.Time `json:"update_at"             orm:"update_at"             description:"修改时间"`  // 修改时间
	DeleteAt            *gtime.Time `json:"delete_at"             orm:"delete_at"             description:"删除时间"`  // 删除时间
	AssociatedAccountId int         `json:"associated_account_id" orm:"associated_account_id" description:"子账户ID"` //
}
