package role_management

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/role"
	"sms_backend/internal/dao"
	"sms_backend/internal/model"
	"sms_backend/internal/model/do"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"sms_backend/library/liberr"
)

type sRole struct{}

func New() *sRole { return &sRole{} }

func init() { service.RegisterRole(New()) }

func (s *sRole) CreatedRole(ctx context.Context, req *role.CreatedReq) (res *role.CreatedRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		exist, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns().Name, req.Name).Exist()
		liberr.ErrIsNil(ctx, err)
		if exist {
			liberr.ErrIsNil(ctx, errors.New("角色已存在"))
		}
		if len(req.Permission) == 0 {
			liberr.ErrIsNil(ctx, errors.New("权限不能为空"))
		}

		roleID, err := dao.Role.Ctx(ctx).Data(do.Role{Name: req.Name, Note: req.Note}).InsertAndGetId()
		liberr.ErrIsNil(ctx, err)

		data := make([]*do.RolePermission, 0)
		for _, v := range req.Permission {
			data = append(data, &do.RolePermission{
				RoleId:       int(roleID),
				PermissionId: v,
			})
		}
		_, err = dao.RolePermission.Ctx(ctx).Data(data).Insert()
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (s *sRole) GetList(ctx context.Context, req *role.ListReq) (res *role.ListRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var info []*model.Role
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

		res = &role.ListRes{List: make([]*model.Role, 0)}
		err := orm.ScanAndCount(&res.List, &res.Total, false)
		liberr.ErrIsNil(ctx, err)

	})
	return
}

func (s *sRole) UpdatedRole(ctx context.Context, req *role.UpdatedReq) (res *role.UpdatedRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err := dao.Role.Ctx(ctx).Data(do.Role{Name: req.Name, Note: req.Note}).WherePri(req.ID).Update()
		liberr.ErrIsNil(ctx, err)

		_, err = dao.RolePermission.Ctx(ctx).Where(dao.RolePermission.Columns().RoleId, req.ID).Delete()
		liberr.ErrIsNil(ctx, err)

		data := make([]*do.RolePermission, 0)
		for _, v := range req.Permission {
			data = append(data, &do.RolePermission{
				RoleId:       req.ID,
				PermissionId: v,
			})
		}
		_, err = dao.RolePermission.Ctx(ctx).Data(data).Insert()
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (s *sRole) DeleteRole(ctx context.Context, req *role.DeletedReq) (res *role.DeletedRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err := dao.RolePermission.Ctx(ctx).Where(dao.RolePermission.Columns().RoleId, req.ID).Delete()
		liberr.ErrIsNil(ctx, err)

		_, err = dao.Role.Ctx(ctx).WherePri(req.ID).Delete()
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (s *sRole) GetPermissionList(ctx context.Context, req *role.GetPermissionReq) (res *role.GetPermissionRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var data []*entity.Permission
		err := dao.Permission.Ctx(ctx).Scan(&data)
		liberr.ErrIsNil(ctx, err)

		res = &role.GetPermissionRes{List: data}
	})
	return
}
