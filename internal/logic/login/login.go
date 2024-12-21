package login

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"sms_backend/api/v1/allUser"
	"sms_backend/internal/consts"
	"sms_backend/internal/dao"
	"sms_backend/internal/model"
	"sms_backend/internal/model/do"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"sms_backend/library/libUtils"
	"sms_backend/library/liberr"
	"sms_backend/utility"
)

type sLogin struct{}

func New() *sLogin { return &sLogin{} }

func init() { service.RegisterLogin(New()) }

func (s *sLogin) Login(ctx context.Context, req *allUser.LoginReq) (res *allUser.LoginRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var user entity.User
		err = dao.User.Ctx(ctx).Where(dao.User.Columns().Name, req.Username).Scan(&user)
		liberr.ErrIsNil(ctx, err)
		if user.Id == 0 {
			liberr.ErrIsNil(ctx, errors.New("用户名不存在"))
		}

		if libUtils.EncryptPassword(req.Password, user.Salt) != user.Password {
			liberr.ErrIsNil(ctx, gerror.New("密码错误"))
		}

		ip := libUtils.GetClientIp(ctx)
		userAgent := libUtils.GetUserAgent(ctx)
		key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(req.Username) + gmd5.MustEncryptString(user.Password)
		if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
			key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(req.Username) + gmd5.MustEncryptString(user.Password+ip+userAgent)
		}

		pidArray, err := dao.RolePermission.Ctx(ctx).Fields(dao.RolePermission.Columns().PermissionId).Where("role_id = ?", user.RoleId).Array()
		liberr.ErrIsNil(ctx, err)
		var pidList []int
		for _, v := range pidArray {
			pidList = append(pidList, v.Int())
		}
		token, err := service.GfToken().GenerateToken(ctx, key, consts.GetUser(map[string]interface{}{
			"user_name":       req.Username,
			"user_id":         user.Id,
			"system_id":       user.SystemId,
			"permission_list": pidList,
		}, consts.EnumEntUser))
		liberr.ErrIsNil(ctx, err, "登录失败，后端服务出现错误")

		var permission []*entity.Permission
		err = dao.Permission.Ctx(ctx).Where("id IN(?)", pidArray).Scan(&permission)
		liberr.ErrIsNil(ctx, err)

		var userPermission []*model.UserPermission
		for _, v := range permission {
			userPermission = append(userPermission, &model.UserPermission{
				Id:        v.Id,
				Pid:       v.Pid,
				Name:      v.Name,
				MenuType:  v.MenuType,
				Path:      v.Path,
				Component: v.Component,
				Meta: model.Meta{
					Title:    v.Title,
					Icon:     v.Icon,
					IsHide:   v.IsHide,
					IsCached: v.IsCached,
				},
			})
		}
		tree := s.UserListTree(0, userPermission)

		data := do.Log{
			UserId:   user.Id,
			UserName: req.Username,
			ClientIp: libUtils.GetClientIp(ctx),
			Function: "用户登录",
			Note:     req.Username + "登录到系统",
			SystemId: user.SystemId,
		}
		err = utility.CreatedLog(ctx, data)
		liberr.ErrIsNil(ctx, err)

		res = &allUser.LoginRes{
			Token:      token,
			Permission: tree,
			Info: &model.UserInfo{
				Id:       user.Id,
				Name:     user.Name,
				SystemId: user.SystemId,
			},
		}
	})
	return
}

func (s *sLogin) ChangePassword(ctx context.Context, req *allUser.ChangePasswordReq) (res *allUser.ChangePasswordRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		userId := service.Context().GetUserId(ctx)
		var user entity.User
		err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, userId).Scan(&user)
		liberr.ErrIsNil(ctx, err)

		// 先判断旧密码是否正确，再更新密码
		if libUtils.EncryptPassword(req.OldPassword, user.Salt) != user.Password {
			liberr.ErrIsNil(ctx, gerror.New("原密码错误"))
		} else {
			salt := grand.Letters(10)
			password := libUtils.EncryptPassword(req.NewPassword, salt)
			user.Password = password
			user.Salt = salt
			_, err = dao.User.Ctx(ctx).Data(user).Update()
			liberr.ErrIsNil(ctx, err)
		}

		data := do.Log{
			UserId:   user.Id,
			UserName: user.Name,
			ClientIp: libUtils.GetClientIp(ctx),
			Function: "设置",
			Note:     user.Name + "修改了密码",
			SystemId: user.SystemId,
		}
		err := utility.CreatedLog(ctx, data)
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (s *sLogin) Logout(ctx context.Context, req *allUser.LogoutReq) (res *allUser.LogoutRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = service.GfToken().RemoveToken(ctx, service.GfToken().GetRequestToken(g.RequestFromCtx(ctx)))
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (s *sLogin) UserListTree(pid int, list []*model.UserPermission) []*model.UserPermissionTree {
	tree := make([]*model.UserPermissionTree, 0) // 初始化树切片
	for _, v := range list {
		if v.Pid == pid { // 如果当前节点的父ID与传入的pid匹配
			node := &model.UserPermissionTree{
				UserPermission: v, // 直接嵌套原始数据
			}
			node.Children = s.UserListTree(v.Id, list) // 递归查找子节点
			tree = append(tree, node)                  // 将节点添加到树中
		}
	}
	return tree
}
