package log

import (
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/allUser"
	"sms_backend/internal/model"
	"sms_backend/internal/model/entity"
)

type ListReq struct {
	g.Meta `path:"/list" tags:"操作日志" method:"post" sm:"获取日志列表"`
	model.PageReq
}

type ListRes struct {
	g.Meta             `mime:"application/json"`
	List               []*entity.Log `json:"list" dc:"日志列表"`
	allUser.GeneralRes `json:"general" dc:"分页信息"`
}
