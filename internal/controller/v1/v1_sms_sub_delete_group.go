package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubDeleteGroup(ctx context.Context, req *sms.SubDeleteGroupReq) (res *sms.SubDeleteGroupRes, err error) {
	return service.SubControllerDeviceManagement().GroupDelete(ctx, req)
}
