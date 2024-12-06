package consts

import "github.com/gogf/gf/v2/os/gtime"

const (
	PageSize = 10 //分页长度
)

var StartTime = gtime.New()

type EnumResponseStatus int

const (
	EnumResponseSuccess EnumResponseStatus = 0 //成功
	EnumResponseWarning EnumResponseStatus = 2 //警告
)

const HKDToRMB = 0.93
