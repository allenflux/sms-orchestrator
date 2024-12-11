package user

import (
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/internal/consts"
)

type LoginReq struct {
	g.Meta   `path:"/user/login" tags:"用户管理" method:"post" summary:"用户登录"`
	Username string `json:"username" v:"passport" dc:"用户名"`
	Password string `json:"password" v:"password" dc:"密码"`
}

type LoginRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token"`
}

type SubRegisterReq struct {
	g.Meta   `path:"/user/sub-register" tags:"用户管理" method:"post" summary:"添加子账号(仅admin可用)"`
	Username string                 `json:"username" v:"passport" dc:"用户名"`
	Password string                 `json:"password" v:"password" dc:"密码"`
	Status   consts.EnumsUserStatus `json:"status" v:"enums" dc:"状态,1-停用 2-启用"`
	Note     string                 `json:"note" dc:"备注"`
}

type SubRegisterRes struct{}
