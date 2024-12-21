package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/subUser"
)

func (c *ControllerSubUser) SubDelete(ctx context.Context, req *subUser.SubDeleteReq) (res *subUser.SubDeleteRes, err error) {
	return service.SubUser().DeleteSubUser(ctx, req)
}
