package v1

import (
	"context"
	"sms_backend/internal/service"

	"sms_backend/api/v1/allUser"
)

func (c *ControllerAllUser) Logout(ctx context.Context, req *allUser.LogoutReq) (res *allUser.LogoutRes, err error) {
	return service.Login().Logout(ctx, req)
}
