// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sms_backend/api/v1/log"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ILog interface {
		GetLogList(ctx context.Context, req *log.ListReq) (res *log.ListRes, err error)
		CreatedLog(r *ghttp.Request)
	}
)

var (
	localLog ILog
)

func Log() ILog {
	if localLog == nil {
		panic("implement not found for interface ILog, forgot register?")
	}
	return localLog
}

func RegisterLog(i ILog) {
	localLog = i
}
