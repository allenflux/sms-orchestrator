package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/subUser"
)

func (c *ControllerSubUser) SubChangeStatus(ctx context.Context, req *subUser.SubChangeStatusReq) (res *subUser.SubChangeStatusRes, err error) {
	return service.SubUser().ChangeStatus(ctx, req)
}
