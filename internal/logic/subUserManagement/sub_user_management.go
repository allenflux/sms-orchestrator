package subUserManagement

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"sms_backend/api/v1/subUser"
	"sms_backend/internal/consts"
	"sms_backend/internal/dao"
	"sms_backend/internal/model"
	"sms_backend/internal/model/do"
	"sms_backend/internal/service"
	"sms_backend/library/libUtils"
	"sms_backend/library/liberr"
	"sms_backend/utility"
)

type sSubUser struct{}

func New() *sSubUser { return &sSubUser{} }

func (s *sSubUser) GetList(ctx context.Context, req *subUser.SubGetListReq) (res *subUser.SubGetListRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var info []*model.SubUser
		orm := g.Model(info).Ctx(ctx).WithAll().Safe()

		if req.Keyword != "" {
			orm = orm.WhereLike("name", "%"+req.Keyword+"%")
		}
		if req.PageNum != 0 && req.PageSize != 0 {
			orm = orm.Page(req.PageNum, req.PageSize)
		}
		if req.OrderBy != "" {
			orm = orm.Order(req.OrderBy)
		}

		res = &subUser.SubGetListRes{List: make([]*model.SubUser, 0)}
		err = orm.ScanAndCount(&res.List, &res.Total, false)
		if req.PageNum > 0 {
			res.Current = req.PageNum
		} else if req.PageSize == 0 {
			res.Current = 1
		}
	})
	return
}

func (s *sSubUser) CreatedSubUser(ctx context.Context, req *subUser.SubRegisterReq) (res *subUser.SubRegisterRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		orm := dao.User.Ctx(ctx)
		maxSystemId, err := orm.Fields("MAX(system_id)").Value()
		liberr.ErrIsNil(ctx, err)

		exist, err := orm.Where(dao.User.Columns().Name, req.Username).Exist()
		liberr.ErrIsNil(ctx, err)
		if exist {
			liberr.ErrIsNil(ctx, errors.New("该用户名已被注册"))
		}

		salt := grand.Letters(10)
		password := libUtils.EncryptPassword(req.Password, salt)
		_, err = orm.Data(do.User{
			Name:     req.Username,
			Password: password,
			Salt:     salt,
			Status:   req.Status,
			RoleId:   2,
			SystemId: maxSystemId.Int() + 1,
			Note:     req.Note,
		}).Insert()
		liberr.ErrIsNil(ctx, err)

		go func() {
			data := do.Log{
				UserId:   service.Context().GetUserId(ctx),
				UserName: service.Context().GetUsername(ctx),
				ClientIp: libUtils.GetClientIp(ctx),
				Function: "子账号后台",
				Note:     "添加子账号" + req.Username,
			}
			err := utility.CreatedLog(ctx, data)
			liberr.ErrIsNil(ctx, err)
		}()
	})
	return
}

func (s *sSubUser) UpdateSubUser(ctx context.Context, req *subUser.SubUpdateReq) (res *subUser.SubUpdateRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		orm := dao.User.Ctx(ctx)
		_, err = orm.WherePri(req.ID).Data(do.User{Note: req.Note}).Update()
		liberr.ErrIsNil(ctx, err)

		username, err := orm.Fields(dao.User.Columns().Name).WherePri(req.ID).Value()
		liberr.ErrIsNil(ctx, err)

		go func() {
			data := do.Log{
				UserId:   service.Context().GetUserId(ctx),
				UserName: service.Context().GetUsername(ctx),
				ClientIp: libUtils.GetClientIp(ctx),
				Function: "子账号后台",
				Note:     "修改子账号" + username.String(),
			}
			err := utility.CreatedLog(ctx, data)
			liberr.ErrIsNil(ctx, err)
		}()
	})
	return
}

func (s *sSubUser) ChangeStatus(ctx context.Context, req *subUser.SubChangeStatusReq) (res *subUser.SubChangeStatusRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		orm := dao.User.Ctx(ctx)
		_, err = orm.WherePri(req.ID).Data(do.User{
			Status: req.Status,
		}).Update()
		liberr.ErrIsNil(ctx, err)

		var status string
		if req.Status == consts.EnumUserStatusDisable {
			status = "停用"
		} else if req.Status == consts.EnumUserStatusEnable {
			status = "启用"
		}

		username, err := orm.Fields(dao.User.Columns().Name).WherePri(req.ID).Value()
		liberr.ErrIsNil(ctx, err)

		go func() {
			data := do.Log{
				UserId:   service.Context().GetUserId(ctx),
				UserName: service.Context().GetUsername(ctx),
				ClientIp: libUtils.GetClientIp(ctx),
				Function: "子账号后台",
				Note:     status + "子账号" + username.String(),
			}
			err := utility.CreatedLog(ctx, data)
			liberr.ErrIsNil(ctx, err)
		}()
	})
	return
}

func (s *sSubUser) DeleteSubUser(ctx context.Context, req *subUser.SubDeleteReq) (res *subUser.SubDeleteRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.User.Ctx(ctx).WherePri(req.ID).Delete()
		liberr.ErrIsNil(ctx, err)

		username, err := dao.User.Ctx(ctx).Fields(dao.User.Columns().Name).WherePri(req.ID).Value()
		liberr.ErrIsNil(ctx, err)

		go func() {
			data := do.Log{
				UserId:   service.Context().GetUserId(ctx),
				UserName: service.Context().GetUsername(ctx),
				ClientIp: libUtils.GetClientIp(ctx),
				Function: "子账号后台",
				Note:     "删除子账号" + username.String(),
			}
			err := utility.CreatedLog(ctx, data)
			liberr.ErrIsNil(ctx, err)
		}()
	})
	return
}
