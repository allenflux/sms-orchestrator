/*
* @desc:用户处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/23 15:08
 */

package sysUser

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"sms_backend/internal/consts"
	"sms_backend/internal/service"
	"sms_backend/library/libResponse"
)

func init() {
	service.RegisterSysUser(New())
}

type sSysUser struct{}

func New() *sSysUser { return &sSysUser{} }

func (s *sSysUser) CheckLogin(r *ghttp.Request) {
	b, failed := service.GfToken().IsLogin(r)
	if b == false {
		g.Log().Error(r.GetCtx(), failed)
		libResponse.RJson(r, failed.Code, failed.Message)
		return
	}
	r.Middleware.Next()
}

func (s *sSysUser) CheckUserAuth(r *ghttp.Request) {
	ctx := r.GetCtx()
	pidList := service.Context().GetPidList(ctx)
	auth := false
	for _, v := range pidList {
		if v == consts.UserPermissionAuth {
			auth = true
			break
		}
	}
	if auth == true {
		r.Middleware.Next()
	} else {
		libResponse.RJson(r, 434, "没有权限")
	}
}

func (s *sSysUser) CheckSubUserAuth(r *ghttp.Request) {
	ctx := r.GetCtx()
	pidList := service.Context().GetPidList(ctx)
	g.Dump(pidList)
	auth := false
	for _, v := range pidList {
		if v == consts.SubUserPermissionAuth {
			auth = true
			break
		}
	}
	if auth == true {
		r.Middleware.Next()
	} else {
		libResponse.RJson(r, 434, "没有权限")
	}
}

func (s *sSysUser) CheckRoleAuth(r *ghttp.Request) {
	ctx := r.GetCtx()
	pidList := service.Context().GetPidList(ctx)
	auth := false
	for _, v := range pidList {
		if v == consts.RolePermissionAuth {
			auth = true
			break
		}
	}
	if auth == true {
		r.Middleware.Next()
	} else {
		libResponse.RJson(r, 434, "没有权限")
	}
}

func (s *sSysUser) CheckLogAuth(r *ghttp.Request) {
	ctx := r.GetCtx()
	pidList := service.Context().GetPidList(ctx)
	auth := false
	for _, v := range pidList {
		if v == consts.LogPermissionAuth {
			auth = true
			break
		}
	}
	if auth == true {
		r.Middleware.Next()
	} else {
		libResponse.RJson(r, 434, "没有权限")
	}
}
