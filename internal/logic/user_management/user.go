package user_management

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"sms_backend/api/v1/user"
	"sms_backend/internal/consts"
	"sms_backend/internal/dao"
	"sms_backend/internal/model"
	"sms_backend/internal/model/do"
	"sms_backend/internal/service"
	"sms_backend/library/libUtils"
	"sms_backend/library/liberr"
	"sms_backend/utility"
)

type sUser struct{}

func New() *sUser { return &sUser{} }

func init() { service.RegisterUser(New()) }

func (s *sUser) GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var info []*model.User
		systemId := service.Context().GetSystemId(ctx)

		orm := g.Model(info).Ctx(ctx).WithAll().Safe()
		if systemId != 1 {
			orm = orm.Where(dao.User.Columns().SystemId, systemId)
		}
		if req.Keyword != "" {
			orm = orm.Where("name LIKE ?", "%"+req.Keyword+"%")
		}
		if req.UserStatus != 0 {
			checker := consts.EnumsUserStatus(req.UserStatus)
			if checker.IsValid() {
				orm = orm.Where("status = ?", req.UserStatus)
			} else {
				liberr.ErrIsNil(ctx, errors.New("未满足枚举条件"))
			}
		}

		if req.PageNum != 0 && req.PageSize != 0 {
			orm = orm.Page(req.PageNum, req.PageSize)
		}
		if req.OrderBy != "" {
			orm = orm.Order(req.OrderBy)
		}
		res = &user.GetListRes{List: make([]*model.User, 0)}
		err = orm.ScanAndCount(&res.List, &res.Total, false)
		if req.PageNum > 0 {
			res.Current = req.PageNum
		} else if req.PageSize == 0 {
			res.Current = 1
		}
	})
	return
}

func (s *sUser) CreatedUser(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		systemId := service.Context().GetSystemId(ctx)
		if systemId == 0 {
			liberr.ErrIsNil(ctx, errors.New("获取当前用户失败"))
		}

		exist, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Name, req.Username).Exist()
		liberr.ErrIsNil(ctx, err)
		if exist {
			liberr.ErrIsNil(ctx, errors.New("该用户名已被注册"))
		}

		salt := grand.Letters(10)
		password := libUtils.EncryptPassword(req.Password, salt)
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Name:     req.Username,
			Password: password,
			Salt:     salt,
			Status:   req.Status,
			RoleId:   req.RoleId,
			SystemId: systemId,
			Note:     req.Note,
		}).Insert()
		liberr.ErrIsNil(ctx, err)

		data := do.Log{
			UserId:   service.Context().GetUserId(ctx),
			UserName: service.Context().GetUsername(ctx),
			ClientIp: libUtils.GetClientIp(ctx),
			Function: "账号管理",
			Note:     "新增账号" + req.Username,
			SystemId: service.Context().GetSystemId(ctx),
		}
		err = utility.CreatedLog(ctx, data)
		liberr.ErrIsNil(ctx, err)

	})
	return
}

func (s *sUser) UpdateUser(ctx context.Context, req *user.UpdateReq) (res *user.UpdateRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.User.Ctx(ctx).WherePri(req.ID).Data(do.User{
			Note:   req.Note,
			RoleId: req.RoleId,
		}).Update()
		liberr.ErrIsNil(ctx, err)

		username, err := dao.User.Ctx(ctx).Fields(dao.User.Columns().Name).WherePri(req.ID).Value()
		liberr.ErrIsNil(ctx, err)

		data := do.Log{
			UserId:   service.Context().GetUserId(ctx),
			UserName: service.Context().GetUsername(ctx),
			ClientIp: libUtils.GetClientIp(ctx),
			Function: "账号管理",
			Note:     "编辑账号" + username.String(),
			SystemId: service.Context().GetSystemId(ctx),
		}
		err = utility.CreatedLog(ctx, data)
		liberr.ErrIsNil(ctx, err)

	})
	return
}

func (s *sUser) DeleteUser(ctx context.Context, req *user.DeleteReq) (res *user.DeleteRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.User.Ctx(ctx).WherePri(req.ID).Delete()
		liberr.ErrIsNil(ctx, err)

		username, err := dao.User.Ctx(ctx).Fields(dao.User.Columns().Name).WherePri(req.ID).Value()
		liberr.ErrIsNil(ctx, err)

		data := do.Log{
			UserId:   service.Context().GetUserId(ctx),
			UserName: service.Context().GetUsername(ctx),
			ClientIp: libUtils.GetClientIp(ctx),
			Function: "账号管理",
			Note:     "删除账号" + username.String(),
			SystemId: service.Context().GetSystemId(ctx),
		}
		err = utility.CreatedLog(ctx, data)
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (s *sUser) ChangeStatus(ctx context.Context, req *user.ChangeStatusReq) (res *user.ChangeStatusRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.User.Ctx(ctx).WherePri(req.ID).Data(do.User{
			Status: req.Status,
		}).Update()
		liberr.ErrIsNil(ctx, err)

		var status string
		if req.Status == consts.EnumUserStatusDisable {
			status = "停用"
		} else if req.Status == consts.EnumUserStatusEnable {
			status = "启用"
		}

		username, err := dao.User.Ctx(ctx).Fields(dao.User.Columns().Name).WherePri(req.ID).Value()
		liberr.ErrIsNil(ctx, err)

		data := do.Log{
			UserId:   service.Context().GetUserId(ctx),
			UserName: service.Context().GetUsername(ctx),
			ClientIp: libUtils.GetClientIp(ctx),
			Function: "账号管理",
			Note:     status + "账号" + username.String(),
			SystemId: service.Context().GetSystemId(ctx),
		}
		err = utility.CreatedLog(ctx, data)
		liberr.ErrIsNil(ctx, err)

	})
	return
}
