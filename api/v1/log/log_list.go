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
	DateRange []string `json:"date_range" dc:"日期区间"`
}

type ListRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.Log `json:"list"`
	allUser.GeneralRes
}
