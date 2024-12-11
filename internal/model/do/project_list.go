// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProjectList is the golang structure of table project_list for DAO operations like Where/Data.
type ProjectList struct {
	g.Meta              `orm:"table:project_list, do:true"`
	Id                  interface{} //
	ProjectName         interface{} // 项目名称
	QuantityDevice      interface{} // 设备数量
	AssociatedAccount   interface{} // 关联账号
	Note                interface{} // 备注
	CreatedAt           *gtime.Time // 创建时间
	UpdateAt            *gtime.Time // 修改时间
	DeleteAt            *gtime.Time // 删除时间
	AssociatedAccountId interface{} //
}
