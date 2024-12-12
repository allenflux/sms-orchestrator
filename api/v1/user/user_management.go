package user

import (
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/allUser"
	"sms_backend/internal/consts"
	"sms_backend/internal/model/entity"
)

type RegisterReq struct {
	g.Meta   `path:"/register" tags:"账号管理" method:"post" sm:"添加账号"`
	Username string                 `json:"username" v:"passport" dc:"用户名"`
	Password string                 `json:"password" v:"password" dc:"密码"`
	Role     int                    `json:"role" v:"required" dc:"角色id"`
	Status   consts.EnumsUserStatus `json:"status" v:"enums" dc:"状态,1-停用 2-启用"`
	Note     string                 `json:"note" dc:"备注"`
}

type RegisterRes struct{}

type GetListReq struct {
	g.Meta `path:"/list" tags:"账号管理" method:"post" sm:"获取账号列表"`
	allUser.GeneralReq
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.User
	allUser.GeneralRes
}

type UpdateReq struct {
	g.Meta `path:"/update" tags:"账号管理" method:"put" sm:"修改账号"`
	ID     int    `json:"id" v:"required" dc:"账号id"`
	Role   int    `json:"role" v:"required" dc:"角色id"`
	Note   string `json:"note" dc:"备注"`
}

type UpdateRes struct{}

type ChangeStatusReq struct {
	g.Meta `path:"/change-status" tags:"账号管理" method:"put" sm:"修改账号状态"`
	ID     int                    `json:"id" v:"required" dc:"账号id"`
	Status consts.EnumsUserStatus `json:"status" v:"enums" dc:"状态,1-停用 2-启用"`
}

type ChangeStatusRes struct{}

type DeleteReq struct {
	g.Meta `path:"/delete" tags:"账号管理" method:"delete" sm:"删除账号"`
	ID     int `json:"id" v:"min:2" dc:"账号id"`
}

type DeleteRes struct{}
