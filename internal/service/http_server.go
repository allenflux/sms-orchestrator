// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IHttpServer interface {
		SendHttpRequest(ctx context.Context, data map[string]interface{}, url string) (res string, err error)
	}
)

var (
	localHttpServer IHttpServer
)

func HttpServer() IHttpServer {
	if localHttpServer == nil {
		panic("implement not found for interface IHttpServer, forgot register?")
	}
	return localHttpServer
}

func RegisterHttpServer(i IHttpServer) {
	localHttpServer = i
}
