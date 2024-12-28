package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/sms"
)

func (c *ControllerSms) SubGroupList(ctx context.Context, req *sms.SubGroupListReq) (res *sms.SubGroupListRes, err error) {
	return service.SubControllerDeviceManagement().ListUserGroups(ctx, req)
}
