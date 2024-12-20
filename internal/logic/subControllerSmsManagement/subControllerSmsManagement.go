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
	sand := dao.SmsMissionReport.Ctx(ctx).Page(req.PageNum, req.PageSize).Where("associated_account_id = ?", req.SubUserID)
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
	res = &sms.SubTaskListRes{}
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

	c := 0
	var group entity.SubGroup
	if err = dao.SubGroup.Ctx(ctx).Where("id = ?", req.GroupID).ScanAndCount(&group, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询SubGroup 错误")
	}
	if c == 0 {
		return nil, errors.New("未查询到相关group id，请检查 group id是否正确")
	}

	var project entity.ProjectList
	c = 0
	if err = dao.ProjectList.Ctx(ctx).Where("id=?", group.ProjectId).ScanAndCount(&project, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询ProjectList错误")
	}
	if c == 0 {
		return nil, errors.New("所查询的Project 不存在 请检查 Project Id 是否正确")
	}
	if c, err := dao.SmsMissionReport.Ctx(ctx).Where("task_name = ?", req.TaskName).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	} else if c != 0 {
		return nil, errors.New("重复的任务名")
	}

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

	data := entity.SmsMissionReport{
		ProjectId:       group.ProjectId,
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
	g.Log().Infof(ctx, "FileDownloadReq: %v", req.FileName)
	var content []byte
	if content, err = ioutil.ReadFile(consts.TaskFilePath + "/" + req.FileName); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("下载时存储的读取文件失败")
	}
	res = &sms.TaskFileDownloadRes{
		JsonData: string(content),
	}
	return
}

// Delete Task

func (s *sSubControllerSmsManagement) TaskReportDelete(ctx context.Context, req *sms.SubTaskDeleteReq) (res *sms.SubTaskDeleteRes, err error) {
	// todo 撤销任务时做任务状态检查
	if _, err = dao.SmsMissionReport.Ctx(ctx).Where("id = ?", req.TaskID).Delete(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("删除记录失败 SmsMissionReport")
	}
	return
}

func (s *sSubControllerSmsManagement) GetSubGetConversationRecord(ctx context.Context, req *sms.SubGetConversationRecordReq) (res *sms.SubGetConversationRecordRes, err error) {
	var chatLog entity.SmsChartLog
	if err = dao.SmsChartLog.Ctx(ctx).Where("id = ?", req.ChatLogID).Scan(&chatLog); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询 DB SmsChartLog 错误")
	}
	if chatLog.AssociatedAccountId != req.SubUserID {
		return nil, errors.New("sub user id 验证错误")
	}
	var chatLogList []*entity.SmsChartLog
	if err = dao.SmsChartLog.Ctx(ctx).Where("target_phone_number = ?", chatLog.TargetPhoneNumber).Where("device_number = ?", chatLog.DeviceNumber).OrderDesc("id").Scan(&chatLogList); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询 DB SmsChartLog 错误2")
	}

	data := make([]sms.SubGetConversationRecordListResData, len(chatLogList))
	for i := range chatLogList {
		data[i].ChatLogID = chatLogList[i].Id
		data[i].RecordTime = chatLogList[i].CreatedAt
		data[i].Content = chatLogList[i].SmsContent
		data[i].TargetPhoneNumber = chatLogList[i].TargetPhoneNumber
		data[i].SentOrReceive = chatLogList[i].SentOrReceive
	}
	res = &sms.SubGetConversationRecordRes{
		Data: data,
	}

	return
}

func (s *sSubControllerSmsManagement) SubGetConversationRecordList(ctx context.Context, req *sms.SubGetConversationRecordListReq) (res *sms.SubGetConversationRecordListRes, err error) {
	var chatLogsId []*entity.SmsChartLog
	var chatLogs []*entity.SmsChartLog
	if err = dao.SmsChartLog.Ctx(ctx).Where("project_id=?", req.ProjectID).Where("associated_account_id=?", req.SubUserID).FieldMax("id", "id").Group("target_phone_number").Group("device_number").Scan(&chatLogsId); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询SmsChartLog IDs错误")
	}
	idList := make([]int, len(chatLogsId))
	for i := range chatLogsId {
		idList[i] = chatLogsId[i].Id
	}

	if err = dao.SmsChartLog.Ctx(ctx).WhereIn("id", idList).Scan(&chatLogs); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询SmsChartLog错误")
	}
	data := make([]sms.SubGetConversationRecordListResData, len(chatLogs))
	for i := range chatLogs {
		data[i].ChatLogID = chatLogs[i].Id
		data[i].RecordTime = chatLogs[i].CreatedAt
		data[i].Content = chatLogs[i].SmsContent
		data[i].TargetPhoneNumber = chatLogs[i].TargetPhoneNumber
		data[i].SentOrReceive = chatLogs[i].SentOrReceive
	}
	res = &sms.SubGetConversationRecordListRes{
		Data: data,
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

// 接收Sub User输入的内容放入 Cache，作为优先任务暴露给Device执行

func (s *sSubControllerSmsManagement) PostConversationRecord(ctx context.Context, req *sms.SubPostConversationRecordReq) (res *sms.SubPostConversationRecordRes, err error) {
	// search DB 获取chart log信息
	var chartLog entity.SmsChartLog
	c := 0
	if err = dao.SmsChartLog.Ctx(ctx).Where("id = ?", req.ChartLogID).ScanAndCount(&chartLog, &c, false); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("数据库操作错误 查询 Task ChartLogID 错误")
	}
	if c == 0 {
		return nil, errors.New("未查询到 Task ChartLog")
	}

	// 验证Sub User id
	if req.SubUserID != chartLog.AssociatedAccountId {
		return nil, errors.New("验证sub user id失败，发送消息不属于原任务 report 的信息")
	}
	// 生成任务信息
	// {TargetPhoneNumber: ... DeviceNumber: ... Content: ... TaskID ...}
	message := sms.SubPostConversationRecordData{
		TaskID:            chartLog.TaskId,
		Content:           req.Content,
		DeviceNumber:      chartLog.DeviceNumber,
		TargetPhoneNumber: chartLog.TargetPhoneNumber,
	}
	// 生成任务队列
	var rIndex int64
	if rIndex, err = g.Redis().LPush(ctx, chartLog.DeviceNumber, message); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("插入redis list 错误")
	}
	res = &sms.SubPostConversationRecordRes{}
	res.TaskItemName = chartLog.DeviceNumber
	res.UnSendTaskNum = rIndex
	return
}
