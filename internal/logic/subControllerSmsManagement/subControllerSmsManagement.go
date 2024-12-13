package subControllerSmsManagement

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"io/ioutil"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/consts"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
)

func New() *sSubControllerSmsManagement {
	return &sSubControllerSmsManagement{}
}

func init() {
	service.RegisterSubControllerSmsManagement(New())
}

type sSubControllerSmsManagement struct{}

func (s *sSubControllerSmsManagement) GetTaskList(ctx context.Context, req *sms.SubTaskListReq) (res *sms.SubTaskListRes, err error) {
	sand := dao.SmsMissionReport.Ctx(ctx).Page(req.PageNum, req.PageSize).Where("owner_account_id = ?", req.SubUserID)
	if req.ProjectID != 0 {
		sand = sand.Where("project_id = ?", req.ProjectID)
	}
	if len(req.TaskName) != 0 {
		sand = sand.Where("task_name like ?", "%"+req.TaskName+"%")
	}

	if len(req.DateRange) > 0 {
		sand = sand.Where("create_at >= ? AND create_at <= ?", req.DateRange[0], req.DateRange[1])
	}
	var data []entity.SmsMissionReport
	var totalCount int
	if err = sand.ScanAndCount(&data, &totalCount, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询DB SmsMissionReport 错误")
	}
	res.Total = totalCount
	res.Data = make([]sms.SubTaskListResData, len(data))
	for i := range data {
		res.Data[i] = sms.SubTaskListResData{
			ID:                data[i].Id,
			ProjectID:         data[i].ProjectId,
			ProjectName:       data[i].ProjectName,
			TaskName:          data[i].TaskName,
			FileName:          data[i].FileName,
			DeviceQuota:       data[i].DeviceQuota,
			TaskStatus:        data[i].TaskStatus,
			SmsQuantity:       data[i].SmsQuantity,
			SurplusQuantity:   data[i].SurplusQuantity,
			QuantitySent:      data[i].QuantitySent,
			AssociatedAccount: data[i].AssociatedAccount,
			IntervalTime:      data[i].IntervalTime,
			StartTime:         data[i].StartTime.String(),
			CreateTime:        data[i].CreatedAt.String(),
		}
	}
	return
}

// Create Task

type FileData struct {
	TargetPhoneNumber []string `json:"target_phone_number"`
	Content           []string `json:"content"`
}

func (s *sSubControllerSmsManagement) TaskCreate(ctx context.Context, req *sms.SubTaskCreateReq) (res *sms.SubTaskCreateRes, err error) {

	filename, err := req.File.Save(consts.TaskFilePath, true)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("文件存储错误")
	}
	// Check 文件名称是否重复
	if c, err := dao.SmsMissionReport.Ctx(ctx).Where("file_name = ?", filename).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	} else if c != 0 {
		return nil, errors.New("文件名称随机重复")
	}

	content, err := ioutil.ReadFile(consts.TaskFilePath + "/" + filename)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("文件打开错误")
	}
	//g.Log().Infof(ctx, "content 文件内容 = %s", string(content))
	// Now let's unmarshall the data into `payload`
	var payload FileData
	err = json.Unmarshal(content, &payload)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("文件解析json错误")
	}

	if len(payload.Content) == 0 || len(payload.TargetPhoneNumber) == 0 || len(payload.Content) != len(payload.TargetPhoneNumber) {
		return nil, errors.New("文件格式错误")
	}

	//将文件存储到 Redis DB
	g.Log().Infof(ctx, "filename===%s", filename)

	if _, err = g.Redis().Do(ctx, "SET", filename, string(content)); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("将文件内容存储到Redis 失败")
	}

	var project entity.ProjectList
	if err = dao.ProjectList.Ctx(ctx).Where("id=?", req.ProjectID).Scan(&project); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询ProjectList错误")
	}

	data := entity.SmsMissionReport{
		ProjectId:       req.ProjectID,
		TaskName:        req.TaskName,
		GroupId:         req.GroupID,
		FileName:        filename,
		TaskStatus:      1,
		SmsQuantity:     len(payload.Content),
		SurplusQuantity: len(payload.TargetPhoneNumber),
		QuantitySent:    0,
		//todo 根据子账号id查询子账号名称
		AssociatedAccount:   "todo 根据子账号id查询子账号名称",
		IntervalTime:        req.IntervalTime,
		StartTime:           req.TimingStartTime,
		ProjectName:         project.ProjectName,
		AssociatedAccountId: req.SubUserId,
	}

	var mrID int64
	if mrID, err = dao.SmsMissionReport.Ctx(ctx).Data(data).InsertAndGetId(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("插入DB SmsMissionReport 错误")
	}
	res = &sms.SubTaskCreateRes{
		ID: mrID,
	}
	return

}

// Download File
func (s *sSubControllerSmsManagement) TaskFileDownload(ctx context.Context, req *sms.TaskFileDownloadReq) (res *sms.TaskFileDownloadRes, err error) {
	g.Log().Infof(ctx, "FileDownloadReq: %v", req)
	res.R.ServeFileDownload(consts.TaskFilePath, req.FileName)
	return
}

// Delete Task

func (s *sSubControllerSmsManagement) TaskReportDelete(ctx context.Context, req *sms.SubTaskDeleteReq) (res *sms.SubTaskDeleteRes, err error) {
	// todo 撤销任务时做任务状态检查
	if _, err = dao.SmsMissionReport.Ctx(ctx).Where("id = ?", req.TaskID).Delete(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return
}

func (s *sSubControllerSmsManagement) GetTaskRecordList(ctx context.Context, req *sms.SubTaskRecordReq) (res *sms.SubTaskRecordRes, err error) {
	sand := dao.SmsMissionRecord.Ctx(ctx).Page(req.PageNum, req.PageSize).Where("associated_account_id = ?", req.SubUserID)
	if req.ProjectID != 0 {
		sand = sand.Where("project_id = ?", req.ProjectID)
	}

	if req.SmsStatus != 0 {
		sand = sand.Where("sms_status = ?", req.SmsStatus)
	}

	if len(req.TaskName) != 0 {
		sand = sand.Where("task_name like ?", "%"+req.TaskName+"%")
	}

	if len(req.TargetPhoneNumber) != 0 {
		sand = sand.Where("target_phone_number like ?", "%"+req.TargetPhoneNumber+"%")
	}

	if len(req.DeviceNumber) != 0 {
		sand = sand.Where("device_number like ?", "%"+req.DeviceNumber+"%")
	}

	if len(req.SentDateRange) > 0 {
		sand = sand.Where("start_time >= ? AND start_time <= ?", req.SentDateRange[0], req.SentDateRange[1])
	}

	if len(req.CreateDateRange) > 0 {
		sand = sand.Where("created_at >= ? AND created_at <= ?", req.CreateDateRange[0], req.CreateDateRange[1])
	}

	var data []entity.SmsMissionRecord
	var totalCount int
	if err = sand.ScanAndCount(&data, &totalCount, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询DB SmsMissionRecord 错误")
	}
	res.Total = totalCount
	res.Data = make([]sms.SubTaskRecordResData, len(data))
	for i := range data {
		res.Data[i] = sms.SubTaskRecordResData{
			ID:                data[i].Id,
			TaskName:          data[i].TaskName,
			SubTaskID:         data[i].SubTaskId,
			TargetPhoneNumber: data[i].TargetPhoneNumber,
			DeviceNumber:      data[i].DeviceNumber,
			SmsTopic:          data[i].SmsTopic,
			SmsContent:        data[i].SmsContent,
			SmsStatus:         data[i].SmsStatus,
			AssociatedAccount: data[i].AssociatedAccount,
			ProjectName:       data[i].ProjectName,
			StartTime:         data[i].StartTime.String(),
			CreateTime:        data[i].CreatedAt.String(),
		}
	}
	return
}
