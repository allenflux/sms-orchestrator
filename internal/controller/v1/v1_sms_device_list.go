package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) DeviceList(ctx context.Context, req *sms.DeviceListReq) (res *sms.DeviceListRes, err error) {
	return service.MainControllerDeviceManagement().GetDeviceList(ctx, req)
}
