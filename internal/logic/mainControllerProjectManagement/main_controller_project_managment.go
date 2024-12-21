package mainControllerProjectManagement

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
)

func New() *sMainControllerProjectManagement {
	return &sMainControllerProjectManagement{}
}

func init() {
	service.RegisterMainControllerProjectManagement(New())
}

type sMainControllerProjectManagement struct{}

func (s *sMainControllerProjectManagement) ProjectListListForFront(ctx context.Context, req *sms.ProjectListForFrontReq) (res *sms.ProjectListForFrontRes, err error) {
	dbTemper := dao.ProjectList.Ctx(ctx)
	if req.SubUserID != 0 {
		g.Log().Info(ctx, "查询子账号下的项目信息")
		dbTemper = dbTemper.Where("associated_account_id = ?", req.SubUserID)
	}
	var projects []*entity.ProjectList
	c := 0
	if err = dbTemper.ScanAndCount(&projects, &c, false); err != nil {
		return nil, errors.New("DB错误 查询ProjectList错误")
	}
	data := make([]*sms.ProjectListForFrontResData, len(projects))
	for i, project := range projects {
		data[i] = &sms.ProjectListForFrontResData{
			ProjectName: project.ProjectName,
			ProjectId:   project.Id,
		}
	}

	res = &sms.ProjectListForFrontRes{
		Data: data,
	}

	return
}

// Project List

func (s *sMainControllerProjectManagement) ProjectList(ctx context.Context, req *sms.ProjectListReq) (res *sms.ProjectListRes, err error) {
	dbTemper := dao.ProjectList.Ctx(ctx).Page(req.PageNum, req.PageSize).Order("id desc")
	if req.KeyWordSearch != "" {
		dbTemper = dbTemper.Where("name like ?", "%"+req.KeyWordSearch+"%")
	}
	var data []entity.ProjectList
	var totalCount int
	if err := dbTemper.ScanAndCount(&data, &totalCount, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("ScanAndCount错误 查询ProjectList")
	}
	res = new(sms.ProjectListRes)
	res.Total = totalCount
	res.Data = make([]sms.ProjectListResData, len(data))
	for i := range data {
		res.Data[i] = sms.ProjectListResData{
			ID:                  data[i].Id,
			Name:                data[i].ProjectName,
			QuantityDevice:      data[i].QuantityDevice,
			AssociatedAccount:   data[i].AssociatedAccount,
			AssociatedAccountId: data[i].AssociatedAccountId,
			Note:                data[i].Note,
			UpdateTime:          data[i].UpdateAt.String(),
		}
	}
	return
}

// Create Project

func (s *sMainControllerProjectManagement) CreateProject(ctx context.Context, req *sms.ProjectCreateReq) (res *sms.ProjectCreateRes, err error) {
	row := &entity.ProjectList{
		ProjectName: req.ProjectName,
		Note:        req.Note,
	}
	// 查询 project重名
	if count, err := dao.ProjectList.Ctx(ctx).Where("project_name = ? ", req.ProjectName).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询ProjectList 错误")
	} else if count > 0 {
		return nil, errors.New("项目名称已存在")
	}
	var rowID int64
	if rowID, err = dao.ProjectList.Ctx(ctx).Data(row).InsertAndGetId(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("插入数据库错误")
	}
	res = &sms.ProjectCreateRes{
		ID: rowID,
	}
	return
}

// Delete Project

func (s *sMainControllerProjectManagement) DeleteProject(ctx context.Context, req *sms.ProjectDeleteReq) (res *sms.ProjectDeleteRes, err error) {

	if _, err := dao.ProjectList.Ctx(ctx).Where("id = ?", req.ProjectID).Insert(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("数据库删除错误")
	}
	return
}

// Edit Project

func (s *sMainControllerProjectManagement) UpdateProject(ctx context.Context, req *sms.ProjectUpdateReq) (res *sms.ProjectUpdateRes, err error) {
	data := make([]entity.ProjectList, 0)
	err = dao.ProjectList.Ctx(ctx).Where("project_name = ?", req.ProjectName).Scan(&data)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	if len(data) > 0 {
		return nil, errors.New("名称已存在")
	}

	if _, err := dao.ProjectList.Ctx(ctx).Data(g.Map{"project_name": req.ProjectName, "note": req.Note}).Where("id = ?", req.ProjectID).Update(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("数据库更新错误")
	}
	return
}

// AllocateAccount2Project

func (s *sMainControllerProjectManagement) AllocateAccount2Project(ctx context.Context, req *sms.AllocateAccount2ProjectReq) (res *sms.AllocateAccount2ProjectRes, err error) {
	data := &entity.ProjectList{}
	if err = dao.ProjectList.Ctx(ctx).Where("id = ?", req.ProjectId).Scan(data); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询数据库错误")
	}
	// todo 查询子用户表 分配相关的获得子账号名

	accountName := "todo 测试子账号"

	// 更新Device list 表
	if _, err = dao.DeviceList.Ctx(ctx).Data(g.Map{"owner_account": accountName, "owner_account_id": req.AccountId}).Where("assigned_items_id = ? ", req.ProjectId).Update(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("Update Error Table  DeviceList")
	}

	if _, err := dao.ProjectList.Ctx(ctx).Data(g.Map{"associated_account_id": req.AccountId, "associated_account": accountName}).Where("id = ?", req.ProjectId).Update(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("数据库更新错误 ProjectList")
	}
	return
}
