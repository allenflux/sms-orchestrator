package subUser

import (
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/allUser"
	"sms_backend/internal/consts"
	"sms_backend/internal/model"
)

type SubRegisterReq struct {
	g.Meta   `path:"/register" tags:"子账号后台(admin)" method:"post" sm:"添加子账号"`
	Name     string                 `json:"name" v:"passport" dc:"用户名"`
	Password string                 `json:"password" v:"password" dc:"密码"`
	Project  []int                  `json:"project" dc:"项目id，不分配则提交空数组"`
	Status   consts.EnumsUserStatus `json:"status" v:"enums" dc:"状态,1-停用 2-启用"`
	Note     string                 `json:"note" dc:"备注"`
}

type SubRegisterRes struct{}

type SubGetListReq struct {
	g.Meta `path:"/list" tags:"子账号后台(admin)" method:"post" sm:"获取子账号列表"`
	allUser.GeneralReq
}

type SubGetListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.SubUser
	allUser.GeneralRes
}

type SubUpdateReq struct {
	g.Meta  `path:"/update" tags:"子账号后台(admin)" method:"put" sm:"修改子账号"`
	ID      int    `json:"id" v:"required" dc:"子账号id"`
	Project []int  `json:"project" v:"required" dc:"项目id，不分配则提交空数组"`
	Note    string `json:"note" dc:"备注"`
}

type SubUpdateRes struct{}

type SubChangeStatusReq struct {
	g.Meta `path:"/change-status" tags:"子账号后台(admin)" method:"put" sm:"修改子账号状态"`
	ID     int                    `json:"id" v:"required" dc:"子账号id"`
	Status consts.EnumsUserStatus `json:"status" v:"enums" dc:"状态,1-停用 2-启用"`
}

type SubChangeStatusRes struct{}

type SubDeleteReq struct {
	g.Meta `path:"/delete" tags:"子账号后台(admin)" method:"delete" sm:"删除子账号"`
	ID     int `json:"id" v:"min:2" dc:"子账号id"`
}

type SubDeleteRes struct{}
