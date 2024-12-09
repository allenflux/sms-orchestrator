package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/career"
)

func (c *ControllerCareer) Register(ctx context.Context, req *career.RegisterReq) (res *career.RegisterRes, err error) {
	return service.CareerDeviceManagement().DeviceRegister(ctx, req)
}
