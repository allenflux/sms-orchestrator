package allUser

import "github.com/gogf/gf/v2/frame/g"

type GeneralReq struct {
	PageNum  int    `json:"page_num" dc:"页码"`
	PageSize int    `json:"page_size" dc:"每页数量"`
	Keyword  string `json:"keyword" dc:"搜索关键字"`
}

type GeneralRes struct {
	Current int `json:"current" dc:"当前页码"`
	Total   int `json:"total" dc:"总数"`
}

type LoginReq struct {
	g.Meta   `path:"/login" tags:"全用户登录" method:"post" sm:"登录"`
	Username string `json:"username" v:"passport" dc:"用户名"`
	Password string `json:"password" v:"password" dc:"密码"`
}

type LoginRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"全用户登录" method:"post" sm:"退出登录"`
}

type LogoutRes struct{}

type ChangePasswordReq struct {
	g.Meta      `path:"/change-password" tags:"全用户登录" method:"put" sm:"修改密码"`
	OldPassword string `json:"old_password" v:"password" dc:"旧密码"`
	NewPassword string `json:"new_password" v:"password" dc:"新密码"`
}

type ChangePasswordRes struct{}
