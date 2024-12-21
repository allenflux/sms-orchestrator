// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ISystemMiddleware interface {
		// Ctx 自定义上下文对象
		Ctx(r *ghttp.Request)
		PrintParams(r *ghttp.Request)
		PrintAndHideError(r *ghttp.Request)
	}
)

var (
	localSystemMiddleware ISystemMiddleware
)

func SystemMiddleware() ISystemMiddleware {
	if localSystemMiddleware == nil {
		panic("implement not found for interface ISystemMiddleware, forgot register?")
	}
	return localSystemMiddleware
}

func RegisterSystemMiddleware(i ISystemMiddleware) {
	localSystemMiddleware = i
}
