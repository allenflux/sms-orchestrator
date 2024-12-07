// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Crons is the golang structure for table crons.
type Crons struct {
	Id          int         `json:"id"           orm:"id"           description:""`         //
	Type        int         `json:"type"         orm:"type"         description:"1-解绑过期设备"` // 1-解绑过期设备
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   description:""`         //
	ExecTime    int         `json:"exec_time"    orm:"exec_time"    description:"耗时"`       // 耗时
	CreatedDate string      `json:"created_date" orm:"created_date" description:"执行日期"`     // 执行日期
}
