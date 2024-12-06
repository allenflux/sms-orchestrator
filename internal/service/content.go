// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"upay_backend/api/v1/admin"
	"upay_backend/api/v1/app"
)

type (
	IContent interface {
		Create(ctx context.Context, req *admin.ContentCreateReq) (res *admin.ContentCreateRes, err error)
		Delete(ctx context.Context, req *admin.ContentDeleteReq) (res *admin.ContentDeleteRes, err error)
		Update(ctx context.Context, req *admin.ContentUpdateReq) (res *admin.ContentUpdateRes, err error)
		List(ctx context.Context, req *admin.ContentListReq) (res *admin.ContentListRes, err error)
		AppList(ctx context.Context, req *app.ContentListReq) (res *app.ContentListRes, err error)
	}
)

var (
	localContent IContent
)

func Content() IContent {
	if localContent == nil {
		panic("implement not found for interface IContent, forgot register?")
	}
	return localContent
}

func RegisterContent(i IContent) {
	localContent = i
}
