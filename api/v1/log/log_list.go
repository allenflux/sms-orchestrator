package log

import (
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/allUser"
	"sms_backend/internal/model/entity"
)

type ListReq struct {
	g.Meta    `path:"/list" tags:"操作日志" method:"post" sm:"获取日志列表"`
	PageNum   int      `json:"page_num" dc:"页码"`
	PageSize  int      `json:"page_size" dc:"每页数量"`
	OrderBy   string   `json:"order_by" dc:"排序,例如【id desc】为id倒叙"`
	DateRange []string `json:"date_range" dc:"日期范围"`
}

type ListRes struct {
	g.Meta             `mime:"application/json"`
	List               []*entity.Log `json:"list" dc:"日志列表"`
	allUser.GeneralRes `json:"general" dc:"分页信息"`
}
