// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppUser is the golang structure of table app_user for DAO operations like Where/Data.
type AppUser struct {
	g.Meta              `orm:"table:app_user, do:true"`
	Id                  interface{} //
	Password            interface{} // 密码
	Salt                interface{} // 盐
	TransactionPassword interface{} // 交易密码
	TransactionSalt     interface{} // 交易密码的盐
	CreatedAt           *gtime.Time // 创建时间
	UpdatedAt           *gtime.Time // 修改时间
	DeletedAt           *gtime.Time // 删除时间
	Uid                 interface{} // 银行uid
	Mobile              interface{} // 手机号
	Status              interface{} // 1-待KYC 2-待上传 3-待配卡 4-待申请 5-已申请
	KycStatus           interface{} // 1-待kyc 2-已kyc
	UserCardStatus      interface{} // 1-未上传 2-已上传
	BankCardStatus      interface{} // 1-未配卡 2-已配卡
	BankImage           interface{} // 万事达照片
	CardLevelId         interface{} // 卡片等级表ID
	UserCardImage       interface{} // 用户卡照片
	Remark              interface{} // 备注
	Email               interface{} // 邮箱
}
