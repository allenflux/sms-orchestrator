// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Client is the golang structure for table client.
type Client struct {
	Id           int         `json:"id"            orm:"id"            description:""`               //
	LabelId      int         `json:"label_id"      orm:"label_id"      description:"标签id"`           // 标签id
	UserName     string      `json:"user_name"     orm:"user_name"     description:"账户"`             // 账户
	UserPassword string      `json:"user_password" orm:"user_password" description:"密码"`             // 密码
	UserSalt     string      `json:"user_salt"     orm:"user_salt"     description:"密码加密的盐"`         // 密码加密的盐
	UserNickname string      `json:"user_nickname" orm:"user_nickname" description:"昵称"`             // 昵称
	Mobile       string      `json:"mobile"        orm:"mobile"        description:"手机号"`            // 手机号
	UserEmail    string      `json:"user_email"    orm:"user_email"    description:"邮箱"`             // 邮箱
	Avatar       string      `json:"avatar"        orm:"avatar"        description:"头像"`             // 头像
	Balance      float64     `json:"balance"       orm:"balance"       description:"余额"`             // 余额
	Credit       float64     `json:"credit"        orm:"credit"        description:"授信余额"`           // 授信余额
	Advance      float64     `json:"advance"       orm:"advance"       description:"预扣金额"`           // 预扣金额
	UseCredit    float64     `json:"use_credit"    orm:"use_credit"    description:"已使用的授信金额"`       // 已使用的授信金额
	UserStatus   int         `json:"user_status"   orm:"user_status"   description:"客户状态0停用1启用2未验证"` // 客户状态0停用1启用2未验证
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间"`           // 创建时间
	UpdatedAt    *gtime.Time `json:"updated_at"    orm:"updated_at"    description:"修改时间"`           // 修改时间
	DeletedAt    *gtime.Time `json:"deleted_at"    orm:"deleted_at"    description:"删除时间"`           // 删除时间
	RoleId       int         `json:"role_id"       orm:"role_id"       description:"角色ID"`           // 角色ID
}
