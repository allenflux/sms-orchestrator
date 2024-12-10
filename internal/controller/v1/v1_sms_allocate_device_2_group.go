package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) AllocateDevice2Group(ctx context.Context, req *sms.AllocateDevice2GroupReq) (res *sms.AllocateDevice2GroupRes, err error) {
	return service.SubControllerDeviceManagement().AllocateDevice2Group(ctx, req)
}
