package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubUpdateGroup(ctx context.Context, req *sms.SubUpdateGroupReq) (res *sms.SubUpdateGroupRes, err error) {
	return service.SubControllerDeviceManagement().UpdateGroup(ctx, req)
}
