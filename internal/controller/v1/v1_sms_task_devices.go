package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) TaskDevices(ctx context.Context, req *sms.TaskDevicesReq) (res *sms.TaskDevicesRes, err error) {
	return service.MainControllerSmsManagement().GetTaskDevices(ctx, req)
}
