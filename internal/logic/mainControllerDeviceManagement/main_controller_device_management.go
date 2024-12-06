package mainControllerDeviceManagement

import (
	"context"
	"sms_backend/api/v1/sms"
	"sms_backend/internal/service"
)

func New() *sMainControllerDeviceManagement {
	return &sMainControllerDeviceManagement{}
}

func init() {
	service.RegisterMainControllerDeviceManagement(New())
}

type sMainControllerDeviceManagement struct{}

func (s *sMainControllerDeviceManagement) CetCardList(ctx context.Context, req *sms.AllocateDevice2ProjectReq) (res *sms.AllocateDevice2ProjectRes, err error) {
	return nil, err
}
