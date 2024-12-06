// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ICommonMiddleware interface {
		MiddlewareCORS(r *ghttp.Request)
	}
)

var (
	localCommonMiddleware ICommonMiddleware
)

func CommonMiddleware() ICommonMiddleware {
	if localCommonMiddleware == nil {
		panic("implement not found for interface ICommonMiddleware, forgot register?")
	}
	return localCommonMiddleware
}

func RegisterCommonMiddleware(i ICommonMiddleware) {
	localCommonMiddleware = i
}
