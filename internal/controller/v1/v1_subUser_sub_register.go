package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/subUser"
)

func (c *ControllerSubUser) SubRegister(ctx context.Context, req *subUser.SubRegisterReq) (res *subUser.SubRegisterRes, err error) {
	return service.SubUser().CreatedSubUser(ctx, req)
}
