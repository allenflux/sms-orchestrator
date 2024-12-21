package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/subUser"
)

func (c *ControllerSubUser) SubUpdate(ctx context.Context, req *subUser.SubUpdateReq) (res *subUser.SubUpdateRes, err error) {
	return service.SubUser().UpdateSubUser(ctx, req)
}
