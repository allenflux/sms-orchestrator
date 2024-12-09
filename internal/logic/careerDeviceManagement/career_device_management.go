package careerDeviceManagement

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/career"
	"sms_backend/internal/dao"
	"sms_backend/internal/model/entity"
	"sms_backend/internal/service"
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
