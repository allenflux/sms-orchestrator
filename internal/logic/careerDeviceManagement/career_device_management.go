package careerDeviceManagement

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"sms_backend/api/v1/career"
	"sms_backend/internal/consts"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
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

	if _, err = dao.DeviceList.Ctx(ctx).Data(raw).Insert(); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("注册数据库错误")
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
	// get Group id
	var device entity.DeviceList
	if err = dao.DeviceList.Ctx(ctx).Where("device_number = ?", req.DeviceNumber).Scan(&device); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("获取group id失败")
	}
	if &device == nil {
		return nil, errors.New("未查询到device信息")
	}
	var job entity.SmsMissionReport
	// 任务状态，1-待发送 2-发送中 3-已发送 4-已撤销
	if err = dao.SmsMissionReport.Ctx(ctx).Where("group_id = ?", device.GroupId).Where("task_status = ?", 1).Scan(&job); err != nil {
		return nil, errors.New("查询数据库Mission Report失败")
	}

	if &job == nil {
		return nil, errors.New("目前设备无可执行任务")
	}

	// Get File
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	g.Log().Infof(ctx, "读取文件中 %s", job.FileName)
	content, err := ioutil.ReadFile(consts.TaskFilePath + "/" + job.FileName)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("文件打开错误")
	}

	// Now let's unmarshall the data into `payload`
	var payload FileData
	err = json.Unmarshal(content, &payload)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("文件解析json错误")
	}

	if len(payload.Content) <= len(payload.DeviceNumber) {
		return nil, errors.New("文件已经被执行完 无任务可以返回")
	}

	payload.DeviceNumber = append(payload.DeviceNumber, req.DeviceNumber)

	// 将结构体的格式，转为json字符串的格式。这里用的到库包是"github.com/pquerna/ffjson/ffjson"
	data, err := ffjson.Marshal(content)
	if err != nil {
		return nil, errors.New("文件格式转换错误")
	}
	// 将json格式的数据写入文件
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if err = ioutil.WriteFile(consts.TaskFilePath+"/"+job.FileName, data, 0777); err != nil {
		g.Log().Error(ctx, err)
		return nil, errors.New("文件写入错误")
	}

	return
}
