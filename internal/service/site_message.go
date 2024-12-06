// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"upay_backend/api/v1/admin"
	"upay_backend/api/v1/app"
	"upay_backend/internal/model/entity"
)

type (
	ISiteMessage interface {
		Create(ctx context.Context, req *admin.SiteMessageCreateReq) (res *admin.SiteMessageCreateRes, err error)
		GetOneById(ctx context.Context, id int) (res *entity.SiteMessage, err error)
		Delete(ctx context.Context, id int) (res *admin.SiteMessageDeleteRes, err error)
		// 管理后台列表
		List(ctx context.Context, req *admin.SiteMessageListReq) (res *admin.SiteMessageListRes, err error)
		// App列表
		AppList(ctx context.Context, req *app.SiteMessageListReq) (res *app.SiteMessageListRes, err error)
		GetReadMessagesByAppUserId(ctx context.Context, appUserId int) (res []*entity.SiteMessageRead, err error)
		// 通过appUserId 获取已读 meaasgId的map
		GetReadMessageIdMapByAppUserId(ctx context.Context, appUserId int) (res map[int]struct{}, err error)
		// App站内信列表
		Detail(ctx context.Context, req *app.SiteMessageDetailReq) (res *app.SiteMessageDetailRes, err error)
		// 是否有未读消息
		HasNoRead(ctx context.Context) (res *app.SiteMessageHasNoReadRes, err error)
	}
)

var (
	localSiteMessage ISiteMessage
)

func SiteMessage() ISiteMessage {
	if localSiteMessage == nil {
		panic("implement not found for interface ISiteMessage, forgot register?")
	}
	return localSiteMessage
}

func RegisterSiteMessage(i ISiteMessage) {
	localSiteMessage = i
}
