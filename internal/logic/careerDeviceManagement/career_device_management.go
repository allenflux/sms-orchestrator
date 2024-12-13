package careerDeviceManagement

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"sms_backend/api/v1/career"
	"sms_backend/internal/consts"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
	"strconv"
	"sync"
)

func New() *sCareerDeviceManagement {
	return &sCareerDeviceManagement{}
}

func init() {
	service.RegisterCareerDeviceManagement(New())
}

type sCareerDeviceManagement struct{}

func (s *sCareerDeviceManagement) DeviceRegister(ctx context.Context, req *career.RegisterReq) (res *career.RegisterRes, err error) {
	raw := entity.DeviceList{
		Number:       req.PhoneNumber,
		DeviceNumber: req.DeviceNumber,
		ActiveTime:   req.ActiveTime,
	}

	if count, err := dao.DeviceList.Ctx(ctx).Where("device_number = ?", raw.DeviceNumber).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	} else if count > 0 {
		return nil, errors.New("设备已注册")
	}
	var rowId int64
	if rowId, err = dao.DeviceList.Ctx(ctx).Data(raw).InsertAndGetId(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("注册数据库错误")
	}
	res = &career.RegisterRes{
		ID: rowId,
	}
	return
}

var rwMutex sync.RWMutex

type FileData struct {
	TargetPhoneNumber []string `json:"target_phone_number"`
	Content           []string `json:"content"`
	DeviceNumber      []string `json:"device_number"`
}

func (s *sCareerDeviceManagement) FetchTasks(ctx context.Context, req *career.FetchTaskReq) (res *career.FetchTaskRes, err error) {
	// todo 限制下device获取任务的次数 每太设备最多可以获取1条任务 上一条任务如果没有提交发送报告则不能再次获取
	var device entity.DeviceList
	if err = dao.DeviceList.Ctx(ctx).Where("device_number = ?", req.DeviceNumber).Scan(&device); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("获取group id失败")
	}
	if &device == nil {
		return nil, errors.New("未查询到device信息")
	}
	var jobs []*entity.SmsMissionReport
	// 任务状态，1-待发送 2-发送中 3-已发送 4-已撤销
	if err = dao.SmsMissionReport.Ctx(ctx).Where("group_id = ?", device.GroupId).Where("task_status = ?", 1).WhereOr("task_status = ?", 2).Scan(&jobs); err != nil {
		return nil, errors.New("查询数据库Mission Report失败")
	}

	if len(jobs) == 0 {
		return nil, errors.New("目前设备无可执行任务List")
	}

	var content []byte
	// 确定需要更新的任务report 条目
	var ii int

	// Get File
	for i, job := range jobs {
		g.Log().Infof(ctx, "读取文件中 %s", job.FileName)
		g.Log().Infof(ctx, "fetch task <<<<< filename===%s", job.FileName)
		if v, err := g.Redis().Do(ctx, "GET", job.FileName); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("Redis Get Error ")
		} else {
			content = gconv.Bytes(v)
			g.Log().Infof(ctx, "fetch content by redis <<<<< filename===%s", string(content))
		}
		ii = i
		if len(content) == 0 {
			g.Log().Info(ctx, "从 mos 中获取的 content 长度为0")
			g.Log().Info(ctx, "从文件中 content 长度为0")
			// 从文件中重新加载的数据属于新的task list【造成这种情况的原因是程序重启导致的偏差】
			if content, err = ioutil.ReadFile(consts.TaskFilePath + "/" + job.FileName); err != nil {
				g.Log().Error(ctx, err)
				return nil, errors.New("存储的读取文件失败")
			}
			if job.TaskStatus == 2 {
				//   从db中获取已经被发送的任务
				var record []*entity.SmsMissionRecord
				if err = dao.SmsMissionRecord.Ctx(ctx).Where("sub_task_id = ?", job.Id).Scan(record); err != nil {
					g.Log().Error(ctx, err)
					return nil, errors.New("查询SmsMissionRecord错误 程序健壮性错误")
				}
				if len(record) == 0 {
					//	有设备领了这个task的任务没有回报 重复发送任务
				} else {
					//	还原文件content 直接添加，因为device num 写入 mos 是顺序数组 所以只需要根据长度就可以还原未完成的任务
					var tpayload FileData
					err = json.Unmarshal(content, &tpayload)
					if err != nil {
						g.Log().Error(ctx, err)
						return nil, errors.New("还原mos时解析错误")
					}
					length := len(record)
					for i := 0; i < length; i++ {
						tpayload.DeviceNumber = append(tpayload.DeviceNumber, "1")
					}
					content, err = ffjson.Marshal(tpayload)
					if err != nil {
						return nil, errors.New("tpayload 文件格式转换错误")
					}

				}

			}
		} else {
			// go out loop
			// fetch base64
			//g.Log().Infof(ctx, "content ------- : %s", string(content))
			//g.Log().Infof(ctx, "content 000 : %s", string(content[0]))
			//g.Log().Infof(ctx, "content : %v", content)
			g.Log().Info(ctx, "go out loop")
			break
		}

	}

	g.Log().Infof(ctx, "ii = %d", ii)

	if len(content) == 0 {
		return nil, errors.New("最终获取的 content 长度为0 说明无可执行任务块 属于异常错误")
	}
	// Now let's unmarshall the data into `payload`
	var payload FileData
	err = json.Unmarshal(content, &payload)
	if err != nil {
		g.Log().Error(ctx, err)
		g.Log().Infof(ctx, "content : %s", string(content))
		return nil, errors.New("文件解析json错误")
	}

	if len(payload.Content) <= len(payload.DeviceNumber) {
		// 当前文件无可执行任务 需要更新挑选机制 遇到这种状况的原因是有device领取了任务没有即使回报 末尾添加发送数量来限制这种情况
		return nil, errors.New("文件已经被执行完 无任务可以返回")
	}

	i := len(payload.DeviceNumber)
	g.Log().Infof(ctx, "length : %d", i)
	res = &career.FetchTaskRes{
		TargetPhoneNumber: payload.TargetPhoneNumber[i],
		Content:           payload.Content[i],
		DeviceNumber:      req.DeviceNumber,
		Interval:          jobs[ii].IntervalTime,
		TaskId:            jobs[ii].Id,
		StartAt:           jobs[ii].StartTime,
	}

	payload.DeviceNumber = append(payload.DeviceNumber, req.DeviceNumber)

	g.Log().Infof(ctx, "before   ==== > %s", payload)
	g.Log().Infof(ctx, "payload.DeviceNumber = %s", payload.DeviceNumber)
	// 将结构体的格式，转为json字符串的格式。这里用的到库包是"github.com/pquerna/ffjson/ffjson"
	data, err := json.Marshal(&payload)
	if err != nil {
		return nil, errors.New("文件格式转换错误")
	}
	//g.Log().Infof(ctx, "data ---- > %s", string(data))
	// 更新后的任务块条目 将json格式的数据写入mos
	if _, err = g.Redis().Do(ctx, "SET", jobs[ii].FileName, string(data)); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("redis 写入文件错误")
	}
	// 判断剩余短信数量
	sq := jobs[ii].SurplusQuantity - 1
	if sq < 0 {
		return nil, errors.New("剩余短信数量不能为小于0的数 请检查程序逻辑")
	}
	// 判断report的sent status 如果 == 1 需要 更新为 == 2
	// 判断文件中最后一条任务 如果是最后一条可以更新DB报告状态为 更新report的sent status 已完成 -> 已完成的report不会被再次挑选出来
	if len(payload.DeviceNumber) == len(payload.Content) {
		if _, err = dao.SmsMissionReport.Ctx(ctx).Data(g.Map{"task_status": 3, "surplus_quantity": sq}).Where("id = ?", jobs[ii].Id).Update(); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("更新SmsMissionReport状态错误")
		}
	} else {
		if _, err = dao.SmsMissionReport.Ctx(ctx).Data(g.Map{"task_status": 2, "surplus_quantity": sq}).Where("id = ?", jobs[ii].Id).Update(); err != nil {
			g.Log().Error(ctx, err)
			return nil, errors.New("更新SmsMissionReport状态错误")
		}
	}

	//
	return
}

func (s *sCareerDeviceManagement) ReportTaskResult(ctx context.Context, req *career.ReportTaskResultReq) (res *career.ReportTaskResultRes, err error) {
	// todo 更新device list中 device 的状态
	var mission entity.SmsMissionReport
	// Get TaskName by task ID
	if err = dao.SmsMissionReport.Ctx(ctx).Where("id = ?", req.TaskId).Scan(&mission); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if &mission == nil {
		return nil, errors.New("非法提交 不存在的任务id")
	}
	// 生成hash串
	// 对整行数据进行hash加研
	hashData := mission.TaskName + strconv.Itoa(req.TaskId) + req.TargetPhoneNumber + req.DeviceNumber + req.Content + strconv.Itoa(req.SendStatus) + mission.AssociatedAccount + strconv.Itoa(mission.AssociatedAccountId) + mission.ProjectName + strconv.Itoa(mission.ProjectId) + req.StartTime.String()
	hashByte := sha256.Sum256([]byte(hashData))
	rowHash := hex.EncodeToString(hashByte[:])
	g.Log().Infof(ctx, "row hash -> %s", rowHash)

	if c, err := dao.SmsMissionRecord.Ctx(ctx).Where("row_hash = ?", rowHash).Count(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("查询数据库 SmsMissionRecord 错误")
	} else if c > 0 {
		return nil, errors.New("数据记录已存在 请勿重复提交")
	}

	data := entity.SmsMissionRecord{
		TaskName:            mission.TaskName,
		SubTaskId:           req.TaskId,
		TargetPhoneNumber:   req.TargetPhoneNumber,
		DeviceNumber:        req.DeviceNumber,
		SmsTopic:            "短信无topic 有只有个文件名比较合理",
		SmsContent:          req.Content,
		SmsStatus:           strconv.Itoa(req.SendStatus),
		AssociatedAccount:   mission.AssociatedAccount,
		AssociatedAccountId: mission.AssociatedAccountId,
		ProjectName:         mission.ProjectName,
		ProjectId:           mission.ProjectId,
		StartTime:           req.StartTime,
		RowHash:             rowHash,
	}

	var rawId int64
	if rawId, err = dao.SmsMissionRecord.Ctx(ctx).Data(data).InsertAndGetId(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	res = &career.ReportTaskResultRes{ID: rawId}
	return
}
