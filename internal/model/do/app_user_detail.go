// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppUserDetail is the golang structure of table app_user_detail for DAO operations like Where/Data.
type AppUserDetail struct {
	g.Meta            `orm:"table:app_user_detail, do:true"`
	Id                interface{} //
	AppUserId         interface{} // 关联app_user表的id
	FirstnameCn       interface{} // 中文姓
	LastnameCn        interface{} // 中文名
	FirstnamePinyin   interface{} // 拼音姓
	LastnamePinyin    interface{} // 拼音名
	Birthday          *gtime.Time // 出生日期
	AnnualIncome      interface{} // 年收入
	Career            interface{} // 职业
	Position          interface{} // 职位
	Region            interface{} // 地区
	Addr              interface{} // 地址
	IdCardNo          interface{} // 身份证号
	IdCardExpiredAt   *gtime.Time // 证件有效期
	LongTermEffective interface{} // 证件是否长期有效1代表是2代表不是
	IdCardFrontPic    interface{} // 身份证正面照片
	IdCardBackPic     interface{} // 身份证背面照片
	HasCommit         interface{} // 是否提交到万事达,1否2是
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 修改时间
	DeletedAt         *gtime.Time // 删除时间
}
