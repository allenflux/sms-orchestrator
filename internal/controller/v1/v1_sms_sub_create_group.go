package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubCreateGroup(ctx context.Context, req *sms.SubCreateGroupReq) (res *sms.SubCreateGroupRes, err error) {
	return service.SubControllerDeviceManagement().GroupCreate(ctx, req)
}
