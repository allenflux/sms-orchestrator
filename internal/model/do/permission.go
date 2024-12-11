// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Permission is the golang structure of table permission for DAO operations like Where/Data.
type Permission struct {
	g.Meta    `orm:"table:permission, do:true"`
	Id        interface{} //
	Name      interface{} // 功能名称
	CreatedAd *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 修改时间
	DeleteAt  *gtime.Time // 删除时间
}
