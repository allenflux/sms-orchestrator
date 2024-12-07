// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpecialOperate is the golang structure for table special_operate.
type SpecialOperate struct {
	Id         int         `json:"id"           orm:"id"           description:""`               //
	Type       int         `json:"type"         orm:"type"         description:"1-绑定 2-解绑 3-延期"` // 1-绑定 2-解绑 3-延期
	CreatedBy  int         `json:"created_by"   orm:"created_by"   description:"操作人"`            // 操作人
	Days       int         `json:"days"         orm:"days"         description:"延期操作前的租赁天数"`     // 延期操作前的租赁天数
	LeaseEndAt *gtime.Time `json:"lease_end_at" orm:"lease_end_at" description:"延期前的租赁到期时间"`     // 延期前的租赁到期时间
	CreatedAt  *gtime.Time `json:"created_at"   orm:"created_at"   description:""`               //
	DeletedAt  *gtime.Time `json:"deleted_at"   orm:"deleted_at"   description:""`               //
}
