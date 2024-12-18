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
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"sms_backend/library/libUtils"
	"sms_backend/library/liberr"
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
		token, err := service.GfToken().GenerateToken(ctx, key, consts.GetUser(map[string]interface{}{
			"user_name": req.Username,
			"user_id":   user.Id,
			"system_id": user.SystemId,
		}, consts.EnumEntUser))
		liberr.ErrIsNil(ctx, err, "登录失败，后端服务出现错误")

		var permission []*model.Permission
		pidList, err := dao.RolePermission.Ctx(ctx).Fields("permission_id").Where("role_id", user.RoleId).Array()
		liberr.ErrIsNil(ctx, err)

		err = dao.Permission.Ctx(ctx).Fields("id,name,redirect").Where("id IN(?)", pidList).Scan(&permission)
		liberr.ErrIsNil(ctx, err)
		res = &allUser.LoginRes{
			Token:      token,
			Permission: permission,
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
