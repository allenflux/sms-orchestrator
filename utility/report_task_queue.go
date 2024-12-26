package utility

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"strconv"
)

type SentStatus int

const reportTaskQueueLogger = "reportTaskQueue"

const (
	UnknownSentStatus SentStatus = iota
	SentSuccess
	SentFailure
	EndSentStatus
)

func (s SentStatus) IsValid() bool {
	if s > UnknownSentStatus && s < EndSentStatus {
		return true
	}
	return false
}

func (s SentStatus) Value() SentStatus {
	if s == 1 {
		return SentSuccess
	}
	if s == 2 {
		return SentFailure
	}
	return UnknownSentStatus
}

func GenHash(taskName, taskId, targetPhoneNumber, deviceNumber, content, sendStatus, aa, aaId, projectName, projectId, st, SentOrReceive string) string {
	// SentOrReceive : 1是为了char log表 表示此短信是发送的 2表示接收
	hashData := taskName + taskId + targetPhoneNumber + deviceNumber + content + sendStatus + aa + aaId + projectName + projectId + st + SentOrReceive
	hashByte := sha256.Sum256([]byte(hashData))
	rowHash := hex.EncodeToString(hashByte[:])
	return rowHash
}

func NewSmsTaskPayload(taskID int, targetPhoneNumber, deviceNumber, content string, sendStatus int, startTime *gtime.Time, reason string) *SmsTaskPayload {

	return &SmsTaskPayload{
		TaskID:            taskID,
		TargetPhoneNumber: targetPhoneNumber,
		DeviceNumber:      deviceNumber,
		Content:           content,
		SendStatus:        sendStatus,
		StartTime:         startTime,
		Reason:            reason,
	}
}

type SmsTaskPayload struct {
	TaskID            int         `json:"task_id"`             // 任务 ID
	TargetPhoneNumber string      `json:"target_phone_number"` // 目标手机号
	DeviceNumber      string      `json:"device_number"`       // 设备编号
	Content           string      `json:"content"`             // 短信内容
	SendStatus        int         `json:"send_status"`         // 发送状态
	StartTime         *gtime.Time `json:"start_time"`          // 任务开始时间
	Reason            string      `json:"reason"`              // 发送失败原因或备注
}

func UpdateTaskResultAndLog(ctx context.Context, taskID int, targetPhoneNumber, deviceNumber, content string, sendStatus int, startTime *gtime.Time, reason string) error {
	var mission entity.SmsMissionReport
	count := 0
	if err := dao.SmsMissionReport.Ctx(ctx).
		Where("id = ?", taskID).
		ScanAndCount(&mission, &count, false); err != nil {
		g.Log(reportTaskQueueLogger).Errorf(ctx, "Failed to query SmsMissionReport for TaskId=%d: %v", taskID, err)
		return fmt.Errorf("查询 SmsMissionReport 失败: %w", err)
	}
	if count == 0 {
		g.Log(reportTaskQueueLogger).Warningf(ctx, "Invalid TaskId=%d: Task does not exist", taskID)
		return errors.New("非法提交: 不存在的任务 ID")
	}

	// 生成行数据哈希
	rowHash := GenHash(
		mission.TaskName,
		strconv.Itoa(taskID),
		targetPhoneNumber,
		deviceNumber,
		content,
		strconv.Itoa(sendStatus),
		mission.AssociatedAccount,
		strconv.Itoa(mission.AssociatedAccountId),
		mission.ProjectName,
		strconv.Itoa(mission.ProjectId),
		startTime.String(),
		"1",
	)
	g.Log().Infof(ctx, "Generated row hash for TaskId=%d: %s", taskID, rowHash)
	// 检查是否重复提交
	if recordCount, err := dao.SmsMissionRecord.Ctx(ctx).
		Where("row_hash = ?", rowHash).
		Count(); err != nil {
		g.Log(reportTaskQueueLogger).Errorf(ctx, "Failed to query SmsMissionRecord for RowHash=%s: %v", rowHash, err)
		return fmt.Errorf("查询 SmsMissionRecord 错误: %w", err)
	} else if recordCount > 0 {
		g.Log(reportTaskQueueLogger).Warningf(ctx, "Duplicate submission detected for RowHash=%s", rowHash)
		return errors.New("数据记录已存在，请勿重复提交")
	}

	// 校验 SendStatus 值
	if !SentStatus(sendStatus).IsValid() {
		g.Log(reportTaskQueueLogger).Errorf(ctx, "Invalid SendStatus=%d for TaskId=%d", sendStatus, taskID)
		Debug(ctx, fmt.Errorf("Invalid SendStatus=%d for TaskId=%d ", sendStatus, taskID).Error(), "filter.txt")
		return errors.New("SendStatus 验证错误，请提交合法的 SendStatus 值")
	}
	status := SentStatus(sendStatus).Value()

	// Retrieve Device information by DeviceNumber
	var device entity.DeviceList
	if err := dao.DeviceList.Ctx(ctx).
		Where("device_number = ?", deviceNumber).
		Scan(&device); err != nil {
		g.Log(reportTaskQueueLogger).Errorf(ctx, "Failed to fetch Device information for DeviceNumber='%s': %v", deviceNumber, err)
		return fmt.Errorf("failed to fetch Device information for DeviceNumber='%s': %w", deviceNumber, err)
	}

	// 准备插入 SmsMissionRecord 数据
	recordData := entity.SmsMissionRecord{
		TaskName:            mission.TaskName,
		SubTaskId:           taskID,
		TargetPhoneNumber:   targetPhoneNumber,
		DeviceNumber:        deviceNumber,
		SmsTopic:            "Placeholder ~Topic",
		SmsContent:          content,
		SmsStatus:           strconv.Itoa(sendStatus),
		AssociatedAccount:   mission.AssociatedAccount,
		AssociatedAccountId: mission.AssociatedAccountId,
		ProjectName:         mission.ProjectName,
		ProjectId:           mission.ProjectId,
		Reason:              reason,
		StartTime:           startTime,
		RowHash:             rowHash,
	}

	// 准备插入 SmsChartLog 数据
	chartLogData := entity.SmsChartLog{
		TaskId:              mission.Id,
		ProjectName:         mission.ProjectName,
		ProjectId:           mission.ProjectId,
		TargetPhoneNumber:   targetPhoneNumber,
		DeviceNumber:        deviceNumber,
		SmsTopic:            "Placeholder 短信内容没有 topic",
		SmsContent:          content,
		SmsStatus:           sendStatus,
		SentOrReceive:       1, // 接收的数据都属于发送结果
		AssociatedAccount:   mission.AssociatedAccount,
		AssociatedAccountId: mission.AssociatedAccountId,
		RowHash:             rowHash,
	}

	// 开启事务
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log(reportTaskQueueLogger).Errorf(ctx, "Failed to start transaction: %v", err)
		return fmt.Errorf("事务操作开启失败: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			g.Log(reportTaskQueueLogger).Errorf(ctx, "Transaction rolled back due to panic: %v", p)
		}
	}()

	// 插入 SmsMissionRecord 数据
	rawId, err := tx.Model("sms_mission_record").Ctx(ctx).Data(recordData).InsertAndGetId()
	if err != nil {
		g.Log(reportTaskQueueLogger).Errorf(ctx, "Failed to : SmsMissionRecord: %v", err)
		_ = tx.Rollback()
		return fmt.Errorf("插入 SmsMissionRecord 失败: %w", err)
	}
	g.Log(reportTaskQueueLogger).Infof(ctx, "Inserted SmsMissionRecord with ID=%d", rawId)

	// 更新 SmsMissionReport 表，根据发送状态调整数量
	var updateData g.Map
	var updateDeviceData g.Map
	switch status {
	case SentSuccess:
		updateData = g.Map{
			"quantity_sent":         mission.QuantitySent + 1,
			"sent_success_quantity": mission.SentSuccessQuantity + 1,
		}
		updateDeviceData = g.Map{
			"quantity_sent": device.QuantitySent + 1,
			"quantity_all":  device.QuantityAll + 1,
			"sent_status":   3, // update status to idle
		}
	case SentFailure:
		updateData = g.Map{
			"quantity_sent":      mission.QuantitySent + 1,
			"sent_fail_quantity": mission.SentFailQuantity + 1,
		}
		updateDeviceData = g.Map{
			"quantity_all": device.QuantityAll + 1,
			"sent_status":  3, // update status to idle
		}
	default:
		g.Log(reportTaskQueueLogger).Errorf(ctx, "Invalid SentStatus: %d", status)
		_ = tx.Rollback()
		return errors.New("未知的发送状态，请检查验证逻辑是否成功")
	}

	//Debug(ctx, strconv.Itoa(mission.QuantitySent+1)+"---- quantity_sent", "quantity_sent.txt")

	// 更新 SmsMissionReport
	if _, err := tx.Model("sms_mission_report").Ctx(ctx).Data(updateData).Where("id = ?", mission.Id).Update(); err != nil {
		g.Log(reportTaskQueueLogger).Errorf(ctx, "Failed to update SmsMissionReport for MissionId=%d: %v", mission.Id, err)
		_ = tx.Rollback()
		return fmt.Errorf("更新 SmsMissionReport 失败: %w", err)
	}

	// 更新 DeviceList
	if _, err := tx.Model("device_list").Ctx(ctx).Data(updateDeviceData).Where("device_number = ?", deviceNumber).Update(); err != nil {
		g.Log(reportTaskQueueLogger).Errorf(ctx, "Failed to update DeviceList for DeviceNumber='%s': %v", deviceNumber, err)
		_ = tx.Rollback()
		return fmt.Errorf("failed to update DeviceList for DeviceNumber='%s': %w", deviceNumber, err)
	}

	// 插入 SmsChartLog 数据
	if _, err := tx.Model("sms_chart_log").Ctx(ctx).Data(chartLogData).Insert(); err != nil {
		g.Log(reportTaskQueueLogger).Errorf(ctx, "Failed to insert SmsChartLog: %v", err)
		_ = tx.Rollback()
		return fmt.Errorf("插入 SmsChartLog 失败: %w", err)
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		g.Log(reportTaskQueueLogger).Errorf(ctx, "Failed to commit transaction: %v", err)
		return fmt.Errorf("事务提交失败: %w", err)
	}

	g.Log(reportTaskQueueLogger).Infof(ctx, "Transaction successfully committed for TaskId=%d", mission.Id)
	return nil
}

var SmsTaskPayloadChan = make(chan *SmsTaskPayload, 100)

func ConsumerReportTaskResult(ctx context.Context, taskChan <-chan *SmsTaskPayload) {
	for {
		select {
		case taskData, ok := <-taskChan: // 从管道中获取任务
			if !ok {
				g.Log(reportTaskQueueLogger).Info(ctx, "Task channel is closed, stopping consumer")
				return
			}
			// 处理任务 taskID int, targetPhoneNumber, deviceNumber, content string, sendStatus int, startTime *gtime.Time, reason string
			if err := UpdateTaskResultAndLog(ctx, taskData.TaskID, taskData.TargetPhoneNumber, taskData.DeviceNumber, taskData.Content, taskData.SendStatus, taskData.StartTime, taskData.Reason); err != nil {
				g.Log(reportTaskQueueLogger).Errorf(ctx, "ReportTaskResultDBAction Error: %v", err)

			} else {
				g.Log(reportTaskQueueLogger).Infof(ctx, "Task successfully processed: %+v", taskData)

			}

		case <-ctx.Done(): // 监听 context 的取消信号
			g.Log(reportTaskQueueLogger).Info(ctx, "Consumer stopped due to context cancellation")
			return
		}
	}
}
