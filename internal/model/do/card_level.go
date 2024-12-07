// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CardLevel is the golang structure of table card_level for DAO operations like Where/Data.
type CardLevel struct {
	g.Meta    `orm:"table:card_level, do:true"`
	Id        interface{} //
	Name      interface{} // 卡片等级名称
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	File      interface{} // 文件base64值
}
