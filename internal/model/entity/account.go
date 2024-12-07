// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Account is the golang structure for table account.
type Account struct {
	Id             uint        `json:"id"               orm:"id"               description:""`                       //
	Username       string      `json:"username"         orm:"username"         description:"账号"`                     // 账号
	Password       string      `json:"password"         orm:"password"         description:"密码"`                     // 密码
	LoginType      int         `json:"login_type"       orm:"login_type"       description:"登录方式1账号密码2验证码"`          // 登录方式1账号密码2验证码
	LoginStatus    int         `json:"login_status"     orm:"login_status"     description:"登录状态1未分配2已登录3已退出4禁用5异常"` // 登录状态1未分配2已登录3已退出4禁用5异常
	UserInfo       string      `json:"user_info"        orm:"user_info"        description:"用户信息"`                   // 用户信息
	FansNum        int         `json:"fans_num"         orm:"fans_num"         description:"粉丝数"`                    // 粉丝数
	FocusNum       int         `json:"focus_num"        orm:"focus_num"        description:"关注数"`                    // 关注数
	LastLoginAt    *gtime.Time `json:"last_login_at"    orm:"last_login_at"    description:"最后登录时间"`                 // 最后登录时间
	CreatedAt      *gtime.Time `json:"created_at"       orm:"created_at"       description:"创建时间"`                   // 创建时间
	UpdatedAt      *gtime.Time `json:"updated_at"       orm:"updated_at"       description:"更新时间"`                   // 更新时间
	AccountGroupId int         `json:"account_group_id" orm:"account_group_id" description:"账户分组ID"`                 // 账户分组ID
	AccountType    int         `json:"account_type"     orm:"account_type"     description:"1-账号 2-邮箱 3-手机号"`        // 1-账号 2-邮箱 3-手机号
	ClientId       int         `json:"client_id"        orm:"client_id"        description:"客户ID"`                   // 客户ID
	DataImportsId  int         `json:"data_imports_id"  orm:"data_imports_id"  description:"数据导入ID"`                 // 数据导入ID
	VerifyUrl      string      `json:"verify_url"       orm:"verify_url"       description:"验证码地址"`                  // 验证码地址
	PlatformType   int         `json:"platform_type"    orm:"platform_type"    description:"平台 1-facebook"`          // 平台 1-facebook
	Sex            int         `json:"sex"              orm:"sex"              description:"性别: 1-女性 2-男性 3-非二次元性别"` // 性别: 1-女性 2-男性 3-非二次元性别
	Avatar         string      `json:"avatar"           orm:"avatar"           description:"头像"`                     // 头像
	Cover          string      `json:"cover"            orm:"cover"            description:"封面图"`                    // 封面图
	Nickname       string      `json:"nickname"         orm:"nickname"         description:"昵称"`                     // 昵称
	Code           string      `json:"code"             orm:"code"             description:"2FA編碼"`                  // 2FA編碼
}
