package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/allUser"
)

func (c *ControllerAllUser) ChangePassword(ctx context.Context, req *allUser.ChangePasswordReq) (res *allUser.ChangePasswordRes, err error) {
	return service.Login().ChangePassword(ctx, req)
}
